# Istio Playground

Istio playground with out-of-box MSA servers included.

## Deployment

### Servers

```bash
# Deploy user server helm chart
make helm-user

# Deploy api server helm chart
make helm-api

# Deploy consumer helm chart
make helm-consumer
```
