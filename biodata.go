package main

import "fmt"

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var individu = []Biodata{
		{nama: "Corry", alamat: "Jakarta", pekerjaan: "IT Engineer", alasan: "masa depan"},
		{nama: "Hamdan", alamat: "Jakarta", pekerjaan: "Back End Programmer", alasan: "masa depan"},
		{nama: "Ken", alamat: "Jakarta", pekerjaan: "Program Analyst", alasan: "masa depan"},
		{nama: "Prasetya", alamat: "Jakarta", pekerjaan: "Algorithm Supervisor", alasan: "masa depan"},
		{nama: "Yoga", alamat: "Jakarta", pekerjaan: "Software Engineer", alasan: "masa depan"},
		{nama: "Ajat", alamat: "Jakarta", pekerjaan: "Blockchain", alasan: "masa depan"},
	}

	fmt.Println(individu)
}
