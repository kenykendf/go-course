package main

import (
	"fmt"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	individu := []Biodata{
		{Nama: "Corry", Alamat: "Jakarta", Pekerjaan: "IT Engineer", Alasan: "masa depan"},
		{Nama: "Hamdan", Alamat: "Jakarta", Pekerjaan: "Back End Programmer", Alasan: "masa depan"},
		{Nama: "Ken", Alamat: "Jakarta", Pekerjaan: "Program Analyst", Alasan: "masa depan"},
		{Nama: "Prasetya", Alamat: "Jakarta", Pekerjaan: "Algorithm Supervisor", Alasan: "masa depan"},
		{Nama: "Yoga", Alamat: "Jakarta", Pekerjaan: "Software Engineer", Alasan: "masa depan"},
		{Nama: "Ajat", Alamat: "Jakarta", Pekerjaan: "Blockchain", Alasan: "masa depan"},
	}

	fmt.Println(individu)
}
