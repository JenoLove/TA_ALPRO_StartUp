package resource

import (
	"fmt"

	"start-up/storage"
)

func CariStartup() {
	var pilihan int
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	var nama string
	fmt.Print("Nama Startup: ")
	fmt.Scan(&nama)

	var indeks = -1
	if pilihan == 1 {
		for i := 0; i < storage.JumlahStartup; i++ {
			if storage.Startups[i].Nama == nama {
				indeks = i
				i = storage.JumlahStartup // untuk keluar
			}
		}
	} else if pilihan == 2 {
		SelectionSortByNamaAsc()
		l, r := 0, storage.JumlahStartup-1
		for l <= r {
			mid := (l + r) / 2
			if storage.Startups[mid].Nama == nama {
				indeks = mid
				l = r + 1
			} else if storage.Startups[mid].Nama < nama {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	if indeks != -1 {
		fmt.Println("Ditemukan:", storage.Startups[indeks])
	} else {
		fmt.Println("Tidak ditemukan.")
	}
}
