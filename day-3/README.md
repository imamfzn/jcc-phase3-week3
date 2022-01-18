# Day #3: Kubernetes

## Setup Kubernetes

### Enable Kubernetes on Docker Desktop

Enable kubernetes pada docker desktop kalian dengan mengikuti step yang ada di https://docs.docker.com/desktop/kubernetes.

### Validasi Kubectl

Untuk melakukan interaksi dengan kubernetes engine / cluster client tools yang digunakan adalah `kubectl`. Jika menggunakan docker desktop seharusnya `kubectl` sudah ter-install. Kalau bisa, selalu gunakan `kubectl` pada terminal dalam mode `WSL`.

```bash
kubectl get nodes
```

Contoh output:

```bash
NAME             STATUS   ROLES                  AGE    VERSION
docker-desktop   Ready    control-plane,master   4d5h   v1.22.4
```

## Create your deployment

1. Build salah satu service yang kalian miliki dengan suatu `tag` tertentu.
2. 

## Setup your configmap

## Expose your deployment through kubernetes service


## Create your kubernetes ingress

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.1/deploy/static/provider/cloud/deploy.yaml
```

