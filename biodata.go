package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	nilai := os.Args[1]
	nilaitoint, err := strconv.ParseInt(nilai, 6, 6)
	if err != nil {
		fmt.Println("Absensi Tidak Sesuai")
	}
	cek := panggilan(int(nilaitoint))

	fmt.Println(cek)
	// fmt.Println("absen", nilai)
}

func panggilan(urutan int) Biodata {
	individu := []Biodata{
		{Id: 0, Nama: "Nama Peserta", Alamat: "Alamat Peserta", Pekerjaan: "Pekerjaan Peserta", Alasan: "Alasan Peserta"},
		{Id: 1, Nama: "Corry", Alamat: "Jakarta", Pekerjaan: "IT Engineer", Alasan: "Demi Masa Depan"},
		{Id: 2, Nama: "Hamdan", Alamat: "Jakarta", Pekerjaan: "Back End Programmer", Alasan: "Gaji Lebih Tinggi"},
		{Id: 3, Nama: "Ken", Alamat: "Jakarta", Pekerjaan: "Program Analyst", Alasan: "Bebas Kemana Aja"},
		{Id: 4, Nama: "Prasetya", Alamat: "Jakarta", Pekerjaan: "Algorithm Supervisor", Alasan: "Ikut Yang Terbaik"},
		{Id: 5, Nama: "Yoga", Alamat: "Jakarta", Pekerjaan: "Software Engineer", Alasan: "Trending"},
		{Id: 6, Nama: "Ajat", Alamat: "Jakarta", Pekerjaan: "Blockchain", Alasan: "Kesempatan Terbaik"},
	}

	return individu[urutan]
}
