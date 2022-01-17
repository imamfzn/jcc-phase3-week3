# Day #2 

Silahkan eksplorasi dari task yang ada di bawah ini untuk dibahas pada live session.

## Tasks

Ada 3 Task utama yang peru kalian kerjakan:

1. Refactor Monolithic to Microservices
2. Eksplorasi API Gateway
3. Implementasi Graceful Shutdown

### Refactor Monolithic to Microservices
Lanjutkan pekerjaan kemarin yang belum selesai.

### API Gateway
Kita akan mencoba menggunakan [Kong](https://konghq.com/kong/) sebagai bahan eksplorasi api gateway dan [konga](https://github.com/pantsel/konga) sebagai admin dashboard dari kong.

Ada 3 hal yang perlu dilakukan:

1. Setup KONG
2. Setup Service & Route pada Kong
3. Setup JWT di Kong

Referensi yang mungkin bermanfaat untuk setup kong:

- https://blog.leonprasetya.my.id/p/instalasi-kong-api-gateway-konga-dengan-docker-compose/
- https://blog.leonprasetya.my.id/p/pembahasan-services-dan-routes-pada-kong-api-gateway/
- https://blog.leonprasetya.my.id/p/implementasi-jwt-pada-kong/

#### Setup KONG

1. Download [docker-compose.yml](./kong/docker-compose.yml) yang berisi dari beberapa container untuk melakukan setup dan menjalankan kong sebagai docker-compose.
2. Jalankan perintah:
    ```bash
    docker-compose -p jcc up --build
    ```
3. Tunggu beberapa saat hingga selesai, jika sudah coba buka `localhost:1337` pada browser kalian
4. Apabila alamat tersebut tidak bisa diakses, silahkan beberapa sekali lagi perintah sblmnya (pastikan sudah keluar dari proses sblmnya, tekan `ctrl+c`). Ini hal yang tricky karena kong & konga harus berjalan setelah migrasi database dikerjakan.
5. Akan muncul halaman register admin, silahkan isi sesuka kalian sebagai credential admin dashboard kong
6. Login menggunakan account yang sudah didaftarkan
7. Buat connection baru dengan mengisi:
    ```
    Name: <bebas>
    Kong Admin URL: http://kong:8001
    ```
8. Setelah ini seharusnya kita akan diarahkan ke dashboard utama konga

#### Setup Service & Route

Silahkan eksplorasi pada konga untuk:

1. Menambahkan / mendaftarkan services yang sudah dibuat pada kong services
2. Untuk masing-masing services buat route untuk mengakses endpoint public dari setiap service

Coba buat sebuah route yang mengarahkan:

```
/api/v1/login -> endpoint login di service auth
/api/v1/users -> endpoint users di service user
/api/v1/ .... -> dst
```

#### JWT
Coba eksplore bagaimana menambahkan plugins jwt pada kong, dan mengimplementasikan nya pada beberapa route.

Catatan:
JWT token tetap digenerate oleh service auth, kong hanya memeriksa jwt token yang dibawa oleh client.


### Gracefull Shutdown
Graceful shutdown merupakan suatu mekanisme untuk mematikan aplikasi / service agar lebih safely. Safely yang dimaksud diantaranya:
1. Memastikan bahwa semua pekerjaan / request yang dikerjakan oleh aplikasi kita sudah dikerjakan terlebih dahulu seblum aplikasi keluar / shutdown
2. Memastikan semua resource yang kita gunakan pada aplikasi sudah di `close`, misal seperti koneksi ke database, cache, dll.

Hal ini perlu dilakukan untuk menghindari terjadinya error di tengah jalan saat melakukan restart aplikasi ataupun deployment. Silahkan cari referensi terkait graceful shutdown.

Coba tambahkan mekanisme `graceful shutdown` pada setiap microservices, dengan catatan:
  1. service baru akan mati ketika menerika signal `SIGTERM`, ataupun `SIGINT` (ctrl+c)
  1. service baru akan mati ketika semua request sudah selesai dikerjakan
  2. close koneksi ke database sebelum service mati
  3. pasang timeout selama `1m` untuk jaga jaga terjadi hang ketika melakukan close connection / http server



