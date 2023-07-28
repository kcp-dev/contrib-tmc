# TMC - Transparent Multi-Cluster

Experimental multi-cluster workload management platform for Kubernetes, built on top of [KCP](https://github.com/kcp-dev/kcp)

## What is TMC?

TMC is a multi-cluster management platform that allows you to manage multiple Kubernetes clusters from a single control plane. TMC should be a plugin of KCP and extend its functionality to support multi-cluster workload management.

## Quick start

```
# Build CLI binaries

  make build

# Copy binaries to PATH

  # Either copy the binaries to a directory in your PATH
  cp bin/{kubectl-tmc,kubectl-workloads} /usr/local/bin/kubectl-tmc

  # Or add the bin directory to your PATH
  export PATH=$PATH:$(pwd)/bin

# Start TMC-KCP

  go run ./cmd/tmc start

# Create TMC workspace

  kubectl tmc workspace create tmc-ws --type tmc --enter

# Create SyncTarget for remote cluster

  kubectl tmc workload sync cluster-1 --syncer-image ghcr.io/kcp-dev/kcp-dev/contrib-tmc/syncer:latest --output-file cluster-1.yaml

# Login into child cluster

  KUBECONFIG=<pcluster-config> kubectl apply -f "cluster-1.yaml"

# Bind compute resources

  kubectl tmc bind compute root:tmc-ws

# Create a workload on TMC-KCP cluster

  kubectl create deployment kuard --image gcr.io/kuar-demo/kuard-amd64:blue
```

## Known issues

- [ ] TMC currently does not support sharding
- [ ] Placements do not work cross-workspaces https://github.com/kcp-dev/contrib-tmc/issues/4
- [ ] Tunneling does not work as of now

## Background

https://github.com/kcp-dev/kcp/issues/2954
