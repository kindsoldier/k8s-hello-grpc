## gRPC hello world


### haproxy ingress

````
helm repo add haproxytech https://haproxytech.github.io/helm-charts
helm repo update
helm search repo haproxy
helm install haproxy haproxytech/kubernetes-ingress
```

### nginx-ingress

```
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
helm install nginx-ingress nginx-stable/nginx-ingress
```

### longhorn pvc

```
helm repo add longhorn https://charts.longhorn.io
helm repo update
helm install longhorn longhorn/longhorn --namespace longhorn --create-namespace --version 1.4.0
```
