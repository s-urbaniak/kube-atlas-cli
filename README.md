# kubectl atlas

A [`kubectl` plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) to interact with [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) leveraging the [MongoDB Atlas Operator](https://www.mongodb.com/docs/atlas/atlas-operator/).

## Quick Start

### Installation

1. Install [krew](https://krew.sigs.k8s.io/)
2. Execute:
```
kubectl krew index add atlas https://github.com/s-urbaniak/kube-atlas-cli.git
kubectl krew install atlas/atlas
```

### Upgrade

To upgrade, consult the [krew documentation](https://krew.sigs.k8s.io/docs/user-guide/upgrading-plugins/) and execute:
```
$ kubectl krew upgrade

Updated the local copy of plugin index "atlas".
  Upgrades available for installed plugins:
    * atlas/atlas v0.0.1 -> v0.0.2
Upgrading plugin: atlas/atlas
Upgraded plugin: atlas/atlas

```

3. Invoke the `atlas` kubectl plugin:
```
kubectl atlas --help
```

## Developer Notes

### Architecture

This plugin builds on top of the official Kubernetes Plugin framework [krew](https://github.com/kubernetes-sigs/krew/tree/master).

An overview about the plugin architecture is available at https://github.com/kubernetes-sigs/krew/blob/master/docs/KREW_ARCHITECTURE.md.

### Release

1. Push a tag in the form of `vx.y.z`:
```
git tag vx.y.z
git push origin --tags
```
2. Once the [release GitHub action](https://github.com/s-urbaniak/kube-atlas-cli/actions/workflows/release.yml) is done
and the release is available at https://github.com/s-urbaniak/kube-atlas-cli/releases, execute:
```
TAG_NAME=vx.y.z GITHUB_TOKEN=<REDACTED> UPSTREAM_KREW_INDEX_REPO_NAME=kube-atlas-cli UPSTREAM_KREW_INDEX_REPO_OWNER=s-urbaniak go run cmd/krew/main.go
```
Replace `GITHUB_TOKEN` and `TAG_NAME` accordingly.
