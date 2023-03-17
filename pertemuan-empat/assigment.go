package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name, Address, Job, Reason string
}

func main() {

	persons := make(map[int]Person, 2)

	persons[1] = Person{
		Name:    "Dody ",
		Address: "Kota Bandung",
		Job:     "Software Engineer",
		Reason:  "Memiliki kemampuan konkurensi yang baik",
	}
	persons[2] = Person{
		Name:    "Eryck Kurniawan",
		Address: "Jakarta Selatan",
		Job:     "Web Designer",
		Reason:  "Performa yang tinggi dan mudah dipelajari",
	}

	getData := func (number int) {
		fmt.Printf("Nama: %s\n", persons[number].Name)
		fmt.Printf("Alamat: %s\n", persons[number].Address)
		fmt.Printf("Pekerjaan: %s\n", persons[number].Job)
		fmt.Printf("Alasan Memilih Golang: %s\n", persons[number].Reason)
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run namafile.go <nomor>")
		os.Exit(1)
	}

	nomor, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor harus berupa angka")
		os.Exit(2)
	}

	switch nomor {
	case 1:
		getData(nomor)
	case 2:
		getData(nomor)
	default:
		fmt.Println("Data absen tidak tersedia")
	}
}
