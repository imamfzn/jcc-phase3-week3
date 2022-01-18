# Deploy Simple Crud App to Kubernetes

Kita akan mencoba melakukan deployment contoh service termudah kita kali ini menggunakan [simple-crud](./simple-crud) app. Silahkan di download dulu repo / directory service tersebut.

## Build Image
Masuk ke directory `simple-crud` pada local wsl kalian, lalu lakukan build image dan sesuaikan `tag` nya seperti dibawah ini:

```bash
docker build -t simple-crud:0.99.0 .
```

## Setup Mongodb on localhost
Karena `simple-crud` service ini membutuhkan dependensi `mongodb` sebagai penyimpanan data, kali ini cukup jalankan `mongodb` di local kalian lewat `docker` dengan port yang di forward adalah `27017` dan tanpa autentikasi apapun, atau bisa melakukan dengan command docker berikut:

```bash
docker volume create mongodb-kubernetes-volume
docker run -d --name mongodb-kubernetes -v mongodb-kubernetes-volume:/data/db -p 27017:27017 mongo:4.4
```

Pastikan mongodb tersebut berjalan dengan lancar dan bisa diakses dari local kalian.

## Setup Configmap

Selanjutnya, kita akan men-setup `configmap` kubernetes yang akan digunakan oleh service `simple-crud`. `configmap` merupakan kumpulan dari suatu `environment` value yang akan digunakan oleh suatu `container` di kubernetes. Jalankan:

```bash
kubectl apply -f deploy/configmap.yml
```

## Setup Deployment

Selanjutnya, kita akan mendeploy service `simple-crud` dengan menggunakan `image` yang sudah kita build. Jalankan perintah:

```bash
kubectl apply -f deploy/deployment.yml
```

File `deployment.yml` tersebut berisi dari spesifikasi deployment yang akan dilakukan di kubernetes, bisa terlihat dari `image` apa yang kita gunakan, deployment `name` yang akan tersimpan, `env` value yang digunakan , dll. Apabila sudah, coba jalankan perintah berikut:

```bash
kubectl get pods
```

Kurang lebih hasilnya akan keluar seperti (mungkin tidak sama persis)

```bash
NAME                              READY   STATUS    RESTARTS   AGE
simple-crud-6fb4684476-2tj45      1/1     Running   0          30m
simple-crud-6fb4684476-4mph8      1/1     Running   0          30m
```

`Pod` merupakan satuan deployment terkecil di kubernetes. 1 Pod bisa berisi dari beberapa `container`. Dalam hal ini kita hanya menggunakan 1 `container` dalam 1 `pod`. Bisa saja apabila deployment nya gagal, maka status pada `pod` bukanlah `Running`, bisa jadi: `CrashLoopBackOff`, dll.

Kita juga bisa masuk ke dalam pod yang sedang running dengan perintah:

```bash
kubectl exec -it <nama_pod> /bin/sh
```

Kita juga bisa melihat log dari setiap pod yang sedang running, dengan perintah:

```bash
kubectl logs -f <nama_pod>
```

Beberapa perintah tersebut sama saja dengan perintah suatu container di docker baik untuk masuk ke dalam container ataupun melihat log.

Bisa jadi kita memiliki banyak deployment, bisa dilihat dengan cara:

```bash
kubectl get deployment
```

## Setup Service

Saat ini aplikasi kita sudah running di kubernetes ada sebanyak 2 pod (silahkan lihat file `deployment.yml`). Namun, deployment yang kita miliki ini belum bisa diakses dari manapun dari luar cluster kubernetes ataupun dari dalam. Agar deployment tersebut dapat diakses, kita bisa menggunakan perintah:

```bash
kubectl expose deployment simple-crud --type NodePort
```

Setelah itu lakukan perintah sperti berikut:

```bash
kubectl get service
```

Kurang lebih hasilnya akan seperti berikut:

```bash
NAME             TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE 
kubernetes       ClusterIP   10.96.0.1        <none>        443/TCP          4d7h
simple-crud      NodePort    10.96.32.33      <none>        5000:30965/TCP   117m
```

`service` di kubernetes bertugas untuk membagikan traffic / request dari luar ke setiap pod yang ada pada suatu deployment. Kalau diliat pada contoh diatas, `port` yang di buat oleh suatu service tersebut adalah `30965`. Apabila kita sudah mendapatkan port tersebut, seharusnya port tersebut dapat di akses lewat local kita.

```bash
curl localhost:30965
```


## Setup Ingress

Sampai tahap ini, sebenarnya service kita sudah bisa diakses dari luar cluster kubernetes. Namun, ada beberap kendala ketika kita hanya menggunakan service sebagai penjembatan traffic dari luar dengan pod deployment kita. Oleh karena itu teradapat suatu controller bernama `ingress` di kubernetes yang tugasnya adalah sebagai `loadbalancer` suatu service.

Jalankan perintah berikut untuk memasang `ingress` pada `simple-crud` app:

```bash
kubectl apply -f deploy/ingress.yml
```

Kalau kalian lihat pada file `ingress.yml`, `ingress` akan menerima traffic dari suatu domain `api.simple.local`, nah agar di local kita bisa membaca domain tersebut kita perlu melakukan sedikit `tricky` pada `/etc/hosts` kita.

Lakukan perintah (pastikan di wsl)

```bash
echo "127.0.0.1 api.simple.local" | sudo tee -a /etc/hosts
```

Apabila sudah, sekarang coba buka browser / lakukan perintah berikut untuk melakukan hit api pada service `simple-crud` lewat ingress yang sudah dibuat:

```bash
curl api.simple.local
```

Seharusnya request tersebut menghasilkan response `200`, apabila blm berhasil tidak husah khawatir, kita akan bahas pada sesi live session.
