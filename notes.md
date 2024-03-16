# Notes Go Language

## Sejarah Golang
- Dibuat di Google menggunakan bahasa pemrograman C
- Dirilis ke publik sebagai *open-source* tahun 2009
- Golang populer sejak digunakan untuk membuat *Docker* pada tahun 2011
- Saat ini mulai banyak teknologi baru yang menggunakan Golang, seperti Kubernetes, Prometheus, CockroachDB, dll.
- Saat ini populer untuk pembuatan Backend API

## Kenapa Golang?
- Bahasa C sangat sederhana sehingga tidak butuh waktu lama untuk mempelajarinya
- Golang mendukung baik concurrency programming -> multicore processor
- Golang mendukung *garbage collector* sehingga tidak butuh melakukan management memory secara manual seperti di bahasa C
- Salah satu bahasa pemrograman yang sedang naik daun

## Process Development Program Golang

**Go Compiler** meng-compile file *.go* sehingga terbuatlah binary file yang bisa dijalankan

## Instalasi dan Running Golang
- Kunjungi [Website Golang](https://golang.org), download, dan instal compiler Golang
- Untuk mengecek apakah sudah terinstal, jalankan perintah `go version` untuk mengecek Golang sudah berhasil terinstal

## Membuat Project

Project di Golang biasanya disebut sebagai module. Berikut perintah untuk membuat module.

`go mod init <nama-module>`

Untuk menjalankan program tersebut, kita bisa menggunakan perintah di terminal.

`go run <nama file go>`