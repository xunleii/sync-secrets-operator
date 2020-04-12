package controller

import (
	"context"
	"net/http"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	kconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/xunleii/sync-secrets-operator/pkg/registry"
)

const (
	prefixAnnotation            = "secret.sync.klst.pw"
	AllNamespacesAnnotation     = prefixAnnotation + "/all-namespaces"
	NamespaceSelectorAnnotation = prefixAnnotation + "/namespace-selector"

	requeueAfter = 5 * time.Second
)

type (
	Controller struct {
		Context
		metricsBindAddress     string
		healthProbeBindAddress string
	}
)

func NewController(metricsBindAddress, healthProbeBindAddress string, ignoredNamespaces []string) *Controller {
	return &Controller{
		Context: Context{
			Context:           context.Background(),
			IgnoredNamespaces: ignoredNamespaces,
			registry:          registry.New(),
		},
		metricsBindAddress:     metricsBindAddress,
		healthProbeBindAddress: healthProbeBindAddress,
	}
}

func (c *Controller) Run(stop <-chan struct{}) {
	mgr, err := manager.New(kconfig.GetConfigOrDie(), manager.Options{
		MetricsBindAddress:     c.metricsBindAddress,
		HealthProbeBindAddress: c.healthProbeBindAddress,
		ReadinessEndpointName:  "/-/readyz",
		LivenessEndpointName:   "/-/healthz",
	})
	if err != nil {
		klog.Fatalf("Unable to set up overall controller manager: %s", err)
	}
	c.client = mgr.GetClient()

	_ = mgr.AddReadyzCheck("readyz", func(req *http.Request) error { return nil })
	_ = mgr.AddHealthzCheck("healthz", func(req *http.Request) error { return nil })

	secretCtrl, err := controller.New("sync-secrets", mgr, controller.Options{
		Reconciler: &SecretReconciler{Context: &c.Context},
	})
	if err != nil {
		klog.Fatalf("Unable to set up individual controller (sync-secrets): %s", err)
	}

	err = secretCtrl.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		klog.Fatalf("Unable to watch %T: %s", &corev1.Secret{}, err)
	}

	ownedSecretCtrl, err := controller.New("sync-owned-secrets", mgr, controller.Options{
		Reconciler: &OwnedSecretReconcilier{Context: &c.Context},
	})
	if err != nil {
		klog.Fatalf("Unable to set up individual controller (sync-owned-secrets): %s", err)
	}

	err = ownedSecretCtrl.Watch(
		&source.Kind{Type: &corev1.Secret{}},
		&handler.EnqueueRequestForOwner{OwnerType: &corev1.Secret{}},
		predicate.Funcs{
			// Ignore creation because it is managed by the secret owner
			CreateFunc:  func(event.CreateEvent) bool { return false },
			DeleteFunc:  func(event.DeleteEvent) bool { return true },
			UpdateFunc:  func(event.UpdateEvent) bool { return true },
			GenericFunc: func(event.GenericEvent) bool { return true },
		},
	)
	if err != nil {
		klog.Fatalf("Unable to watch owned %T: %s", &corev1.Secret{}, err)
	}

	err = mgr.Start(stop)
	if err != nil {
		klog.Fatalf("Failed to run overall controller manager: %s", err)
	}
}
