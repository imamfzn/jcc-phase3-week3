# Setup Kubernetes

## Enable Kubernetes on Docker Desktop

Enable kubernetes pada docker desktop kalian dengan mengikuti step yang ada di https://docs.docker.com/desktop/kubernetes.

## Validasi Kubectl

Untuk melakukan interaksi dengan kubernetes engine / cluster client tools yang digunakan adalah `kubectl`. Jika menggunakan docker desktop seharusnya `kubectl` sudah ter-install. Kalau bisa, selalu gunakan `kubectl` pada terminal dalam mode `WSL`.

```bash
kubectl get nodes
```

Contoh output:

```bash
NAME             STATUS   ROLES                  AGE    VERSION
docker-desktop   Ready    control-plane,master   4d5h   v1.22.4
```

## Create Kubernetes Dashboard

Untuk mempermudah melihat resource yang kita buat di kubernetes, kita bisa menggunakan kubernetes dashboard. Dashboard ini juga bisa berguna sebagai basic monitoring service kita di kubernetes.
Silahkan ikuti langkah-langkah yang ada pada: https://andrewlock.net/running-kubernetes-and-the-dashboard-with-docker-desktop/. Apabila tidak berhasil / resource tidak muncul di dashboard, silahkan skip saja, karena tidak mandatory.

## Setup Kubernetes Ingress
`ingress` merupakan salah satu controller di kubernetes yang bertugas sebagai `loadbalancer` service suatu deployment.

Lakukan perintah berikut untuk men-setup `ingress`:

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.1/deploy/static/provider/cloud/deploy.yaml
```
