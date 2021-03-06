<a name="unreleased"></a>
## [Unreleased]

### Bug Fixes
- avoid removing old metadata on update

### Code Testing
- use external library for GoDog feature context
- remove old Ginko tests

### Features
- add some 'list' Step in KubernetesFeatureContext
- implements manual GC on delete
- implements FeatureContext for Kubernetes


<a name="v0.2.1"></a>
## [v0.2.1] - 2020-05-15
### Code Testing
- prepare tests for origin metadata
- prepare tests for the protection feature

### Features
- implements origin metadata
- implements fields protection feature
- add protected labels and annotation flags

### Pull Requests
- Merge pull request [#7](https://github.com/xunleii/sync-secrets-controller/issues/7) from xunleii/feat-add-origin-annotations
- Merge pull request [#6](https://github.com/xunleii/sync-secrets-controller/issues/6) from xunleii/feat-ignore-metadata


<a name="v0.2.0"></a>
## [v0.2.0] - 2020-04-16
### Code Refactoring
- move methods in appropriate files
- rewrote controller in order to valid specs
- rewrote controller & update tests
- rewrote ownedSecretReconciler with SynchronizeSecret
- rewrote secretReconciler with SynchronizeSecret
- rename context in Context

### Code Testing
- implement namespace reconcilier tests
- rewrote all test with more specs
- add SynchronizeSecret tests

### Features
- implement namespace reconcilier
- link namespace reconciler with manager
- add missing registry methods & update tests
- add target on SynchronizeSecret
- add synchronizeSecret methods
- add internal secret registry

### Pull Requests
- Merge pull request [#5](https://github.com/xunleii/sync-secrets-controller/issues/5) from xunleii/feat-sync-namespaces
- Merge pull request [#4](https://github.com/xunleii/sync-secrets-controller/issues/4) from xunleii/refact-controller


<a name="v0.1.1"></a>
## [v0.1.1] - 2020-04-04
### Pull Requests
- Merge pull request [#3](https://github.com/xunleii/sync-secrets-controller/issues/3) from xunleii/repo-remove-gorelease-dependency


<a name="v0.1.0"></a>
## [v0.1.0] - 2020-03-31
### Code Refactoring
- clean CI files & removed old ones
- rename controller to operator
- clean controller manager

### Code Testing
- add specs for reconcileOwnedSecret
- add specs for reconcileSecret

### Features
- add owned secret controller
- rewrite controller


<a name="v0.0.0"></a>
## v0.0.0 - 2020-03-08
### Features
- POC sync-controller


[Unreleased]: https://github.com/xunleii/sync-secrets-controller/compare/v0.2.1...HEAD
[v0.2.1]: https://github.com/xunleii/sync-secrets-controller/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/xunleii/sync-secrets-controller/compare/v0.1.1...v0.2.0
[v0.1.1]: https://github.com/xunleii/sync-secrets-controller/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/xunleii/sync-secrets-controller/compare/v0.0.0...v0.1.0
