package main

import (
	"fmt"
	"pertemuan-delapan/go-postgresql/config"
	"pertemuan-delapan/go-postgresql/services"
)

func main() {
	config.InitDb()
	db := config.GetDb()

	defer db.Close()

	// services.CreateEmployee(db, "Erick", "erick@gmail.com", 23, "Tech Data")

	fmt.Printf("services.GetEmployee(db): %v\n", services.GetEmployee(db))

	// count := services.UpdateEmployee(db, 3, "Erick", "erick@gmail.com", 23, "Finance")
	// if count < 1 {
	// 	fmt.Println("Gagal mengubah data!")
	// 	os.Exit(1)
	// }

	// fmt.Println("Berhasil mengubah data!")

	// fmt.Printf("services.DeleteEmployee(db, 1): %v\n", services.DeleteEmployee(db, 1))
}
