# Microservices Introduction

## Monolithic Application
Suatu aplikasi dikatakan sebagai aplikasi monolithic ketika:
* Terdapat banyak fitur / domain function pada satu aplikasi
* Aplikasi memiliki size yang cukup besar
* Startup untuk menjalankan aplikasi membutuhkan waktu yang lama
* Kerusakan yang terjadi pada suatu fitur di aplikasi bisa jadi membuat kerusakan pada seluruh fitur
* Satu aplikasi besar yang dikerjakan oleh banyak anggota / team.

Namun dengan menggunakan aplikasi monolithic terdapat beberapa keuntungan:
* kita hanya perlu fokus pada 1 aplikasi saja ketika development
* deployment lebih mudah
* E2E test pun dapat dijalankan lebih mudah.

Beberapa concern utama yang terjadi ketika aplikasi monolithic sudah cukup gemuk, bisnis semakin berkembang, dan jumlah anggota ataupun tim bertambah diantaranya:
* Perubahan pada suatu aplikasi memerlukan deployment keseluruhan fitur
* Cost untuk melakukan scaling lebih mahal
* Ownership yang tidak jelas ketika jumlah anggota dan tim banyak

## Microservice Application

Microservice merupakan salah satu arsitektur aplikasi yang terdiri dari beberapa aplikasi-aplikasi kecil yang mengerjakan domain function tertentu secara spesifik. Dengan adanya arsitektur microservices ada beberapa keuntungan, diantaranya:

* Scaling menjadi lebih mudah dan murah
* Bisa melakukan suatu deployment independent pada suatu fitur tanpa mengganggu fitur lain
* Ownership menjadi lebih jelas ketika suatu service dipegang oleh suatu tim tertentu
* Kerusakan pada suatu fitur / service tidak menyebabkan kerusakan pada semua fitur

Namun dengan adanya arsitektur microservices aplikasi kita akan menjadi kompleks, diantaranya:

* akan lebih sulit untuk melakukan e2e test
* jika arsitektur tidak di-design dengan baik maka hanya akan menambah masalah baru dibanding menggunakan monolithic
* perlu memikirkan failure handling dari komunikasi antar service
* proses yang tadinya dilakukan oleh 1 service dapat menjadi berubah ketika dilakukan oleh beberapa service yang bisa menyebabkan masalah baru yang perlu dihandle

Microservice cocok digunakan ketika aplikasi monolithic kita sudah sangat bertambah besar yang melakukan banyak hal, juga ketika skala tim / jumlah anggota mulai bertambah banyak. Saat pertama kali membuat aplikasi dari awal, diutamakan tidak langsung menggunakan arsitektur microservice. Namun bertahap dengan memecah suatu aplikasi menjadi beberapa service jiga sudah dibutuhkan. Hal ini untuk mencegah over-engineering pada suatu aplikasi yang sebenarnya domain nya masih sedikit / kecil.

## Microservices Pattern

### Database Pattern
Berdasarkan penggunaan database, dibedakan menjadi:

* Shared database, beberapa service berbagi database yang sama. Bisa berbagi table yang sama, ataupun table yang berbeda.
* Independent database, setiap service memiliki database nya masing-masing, dan setiap database hanya bisa diakses oleh service tersebut saja.
* Hybrid, gabungan dari shared database dan independent database.

Yang harus diperhatikan adalah:
* Penggunaan shared database bisa jadi:
  * downtime yang terjadi pada database melibatkan beberapa service
  * menimbulkan issue security
  * namun query bisa menjadi lebih sederhana, lebih mudah, dan dimungkinkan melakukan transactional query
* Penggunaan independent database bisa jadi:
  * query menjadi lebih kompleks
  * data menjadi redundant
  * size pada storage menjadi lebih besar

### Inter Service Communication
Berdasarkan cara komunikasi antar service dibedakan menjadi:

* synchronous, contohnya: REST, RPC, gRPC
* asynchronous, contohnya: event driven, event bus, ...


## Memecah Monolithic Menjadi Microservices

Ada beberapa **mazhab** yang digunakan orang-orang untuk memecah aplikasi monolithic menjadi microservices. Namun yang paling umum adalah dipecah berdasarkan domain dari suatu bisnis. Saat memecah menjadi microservice yang perlu diperhatikan adalah, sebisa mungkin perubahan yang terjadi pada suatu service tidak akan berdampak pada service yang lain (Indenpendent Changes).


## Handling Failure

Karena komunikasi antar service bisa saja gagal, maka kita harus mengatasi hal-hal yang mungkin terjadi.

### Timeout

Selalu gunakan nilai `timeout` yang kita defined sendiri ketika melakukan pemanggilan ke service lain. Misal toleransi timeout yang kita defined adalah `10s` maka apabila service yang dipanggil tidak bisa memberikan response kurang dari nilai timeout yang kita tentukan, kita bisa mengasumsikan hal ini merupakan failure. Namun bisa jadi nilai timeout diatur dengan melihat performa dari suatu service yang dipanggil. Misal service yang dipanggil memang memiliki performa untuk memberikan response dengan waktu hinngga maksimal `30s` maka mungkin timeout yang kita defined mengikuti waktu tersebut.

### Retry

Misal suatu service-a memanggil service-b. Dengan kondisi tertentu service-b tiba-tiba tidak dapat dipanggil. Untuk mengatasi ini kita bisa menggunakan konsep `retry` saat pemanggilan service yaitu dengan melakukan pemanggilan berulang-ulang dengan harapan pemanggilan ke-sekian akan berhasil.

Misal kita melakukan retry sebanyak 3x ketika terjadi kegagalan pemanggilan.

```
1. service-a call service-b - failed
2. service-a call service-b - failed
3. service-a call service-b - success
```

Pada contoh tersebut pemanggilan ke service-b pada percobaan ketika berhasil. Apabila kita tidak melakukan pemanggilan menggunakan retry, maka otomatis serice-a akan menganggap service-b sedang tidak bisa diakses / error. Bisa jadi juga pada pemanggilan ke-2 service b sudah bisa dipanggil. Worst-case nya pemanggilan ke-3 pun bisa jadi gagal.

Retry biasanya di kombinasikan dengan pemasangan `timeout` untuk menentukan definisi timerange `failure` yang terjadi.

### Retry with Backoff

Sama dengan retry, namun ketika terjadi failure saat pemanggilan kita tidak langsung melakukan pemanggilan berikutnya, tapi memberikan jeda beberapa saat untuk memberikan ruang kepada service yang kita panggil untuk siap kembali. Misal:

```
1. service-a call service-b - failed - service-a wait for 100ms
2. service-a call service b - failed - service-a wait for 200ms
3. service-a call service-b - success
```

Nilai / waktu tunggu tersebut bisa konstan, bisa juga bertambah setiap pemanggilan.

### Circuit Breaker

Untuk teknik ini, silahkan coba cari referensi dan eksplore sendiri.

## Bacaan Menarik:
- https://microservices.io/index.html 
- https://medium.engineering/microservice-architecture-at-medium-9c33805eb74f
