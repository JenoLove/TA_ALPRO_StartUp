package resource

import (
	"fmt"

	"start-up/storage"
)

func TambahAnggotaTim() {
	var id int
	fmt.Print("Masukkan ID Startup: ")
	fmt.Scan(&id)

	for i := 0; i < storage.JumlahStartup; i++ {
		if storage.Startups[i].ID == id {
			if storage.Startups[i].JumlahAnggota >= 5 {
				fmt.Println("Tim penuh.")
				return
			}

			var t storage.Tim
			fmt.Print("Nama Anggota: ")
			fmt.Scan(&t.NamaAnggota)
			fmt.Print("Peran: ")
			fmt.Scan(&t.Peran)

			storage.Startups[i].Tim[storage.Startups[i].JumlahAnggota] = t
			storage.Startups[i].JumlahAnggota++
			fmt.Println("Anggota ditambahkan.")
			return
		}
	}

	fmt.Println("Startup tidak ditemukan.")
}
