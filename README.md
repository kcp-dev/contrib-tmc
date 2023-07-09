# TMC - Transparent Multi-Cluster

WIP - This is a work in progress

## What is TMC?

TMC is a multi-cluster management platform that allows you to manage multiple Kubernetes clusters from a single control plane. TMC should be plugin of KCP
and extend its functionality to support multi-cluster workload management.

## Quick start

```
# Build CLI binaries

  make build

# Start TMC-KCP

  go run ./cmd/tmc start

# Create TMC workspace

  kubectl tmc workspace create tmc-ws --type tmc --enter

# Create SyncTarget for remote cluster

  kubectl tmc workload sync cluster-1 --syncer-image quay.io/faroshq/tmc/syncer:latest --output-file cluster-1.yaml

# Bind compute resources

  kubectl kcp bind compute root:tmc-ws

# Login into child cluster

  KUBECONFIG=<pcluster-config> kubectl apply -f "cluster-1.yaml"

# Create a workload on TMC-KCP cluster

  kubectl create deployment kuard --image gcr.io/kuar-demo/kuard-amd64:blue


```

## Background

https://github.com/kcp-dev/kcp/issues/2954
