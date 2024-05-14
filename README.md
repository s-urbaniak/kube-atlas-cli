# kubectl atlas

A `kubectl` plugin to interact with [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) leveraging the [MongoDB Atlas Operator](https://www.mongodb.com/docs/atlas/atlas-operator/).

## Quick Start

1. Install [krew](https://krew.sigs.k8s.io/)
2. Execute:
```
kubectl krew index add atlas https://github.com/s-urbaniak/kube-atlas-cli.git
kubectl krew install atlas/atlas
kubectl atlas
```
