# kubectl atlas

A `kubectl` plugin to interact with [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) leveraging the [MongoDB Atlas Operator](https://www.mongodb.com/docs/atlas/atlas-operator/).

## Quick Start

1. Install [krew](https://krew.sigs.k8s.io/)
2. Execute:
```
kubectl krew index add atlas https://github.com/s-urbaniak/kube-atlas-cli.git
kubectl krew install atlas/atlas
```

To upgrade, consult the [krew documentation](https://krew.sigs.k8s.io/docs/user-guide/upgrading-plugins/).

3. Invoke the `atlas` kubectl plugin:
```
kubectl atlas --help
```

## Developer Notes

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
