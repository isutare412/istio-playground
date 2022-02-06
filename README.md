# Istio Playground

Istio playground with out-of-box MSA servers included.

## Deployment

### Istio

```bash
# Deploy istio operator
make istio-operator

# Deploy istiod, ingress
make istio

# Deploy addons for istio
make istio-addon
```

### Servers

Enable envoy sidecar injection for your namespace.

```bash
export MYNS=dev
kubectl label namespace $MYNS istio-injection=enabled
```

Deploy applications into cluster.

```bash
# Deploy user server helm chart
make helm-user

# Deploy api server helm chart
make helm-api

# Deploy consumer helm chart
make helm-consumer
```
