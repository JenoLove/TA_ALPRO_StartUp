package resource

import (
	"fmt"

	"start-up/storage"
)

func SelectionSortByNamaAsc() {
	for i := 0; i < storage.JumlahStartup-1; i++ {
		min := i
		for j := i + 1; j < storage.JumlahStartup; j++ {
			if storage.Startups[j].Nama < storage.Startups[min].Nama {
				min = j
			}
		}
		if min != i {
			storage.Startups[i], storage.Startups[min] = storage.Startups[min], storage.Startups[i]
		}
	}
}

func InsertionSortByTahunDesc() {
	for i := 1; i < storage.JumlahStartup; i++ {
		key := storage.Startups[i]
		j := i - 1
		for j >= 0 && storage.Startups[j].TahunBerdiri < key.TahunBerdiri {
			storage.Startups[j+1] = storage.Startups[j]
			j--
		}
		storage.Startups[j+1] = key
	}
}

func UrutkanStartup() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada data startup untuk diurutkan.")
		return
	}

	var kategori, metode, urutan int

	fmt.Println("\n--- Pengurutan Data Startup ---")

	fmt.Println("Kategori pengurutan:")
	fmt.Println("1. Total Pendanaan")
	fmt.Println("2. Tahun Berdiri")
	fmt.Print("Pilih (1/2): ")
	if _, err := fmt.Scanln(&kategori); err != nil || (kategori != 1 && kategori != 2) {
		fmt.Println("Input tidak valid. Silakan masukkan angka 1 atau 2.")
		return
	}

	fmt.Println("Metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih (1/2): ")
	if _, err := fmt.Scanln(&metode); err != nil || (metode != 1 && metode != 2) {
		fmt.Println("Input tidak valid. Silakan masukkan angka 1 atau 2.")
		return
	}

	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih (1/2): ")
	if _, err := fmt.Scanln(&urutan); err != nil || (urutan != 1 && urutan != 2) {
		fmt.Println("Input tidak valid. Silakan masukkan angka 1 atau 2.")
		return
	}

	switch metode {
	case 1:
		if kategori == 1 {
			SelectionSortPendanaan(urutan == 1)
		} else {
			SelectionSortTahun(urutan == 1)
		}
	case 2:
		if kategori == 1 {
			InsertionSortPendanaan(urutan == 1)
		} else {
			InsertionSortTahun(urutan == 1)
		}
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	TampilkanStartup(storage.Startups[:storage.JumlahStartup])
}

func SelectionSortPendanaan(asc bool) {
	for i := 0; i < storage.JumlahStartup-1; i++ {
		idx := i
		for j := i + 1; j < storage.JumlahStartup; j++ {
			if asc && storage.Startups[j].TotalDana < storage.Startups[idx].TotalDana {
				idx = j
			}
			if !asc && storage.Startups[j].TotalDana > storage.Startups[idx].TotalDana {
				idx = j
			}
		}
		storage.Startups[i], storage.Startups[idx] = storage.Startups[idx], storage.Startups[i]
	}
}

func SelectionSortTahun(asc bool) {
	for i := 0; i < storage.JumlahStartup-1; i++ {
		idx := i
		for j := i + 1; j < storage.JumlahStartup; j++ {
			if asc && storage.Startups[j].TahunBerdiri < storage.Startups[idx].TahunBerdiri {
				idx = j
			}
			if !asc && storage.Startups[j].TahunBerdiri > storage.Startups[idx].TahunBerdiri {
				idx = j
			}
		}
		storage.Startups[i], storage.Startups[idx] = storage.Startups[idx], storage.Startups[i]
	}
}

func InsertionSortPendanaan(asc bool) {
	for i := 1; i < storage.JumlahStartup; i++ {
		key := storage.Startups[i]
		j := i - 1
		for j >= 0 && ((asc && storage.Startups[j].TotalDana > key.TotalDana) || (!asc && storage.Startups[j].TotalDana < key.TotalDana)) {
			storage.Startups[j+1] = storage.Startups[j]
			j--
		}
		storage.Startups[j+1] = key
	}
}

func InsertionSortTahun(asc bool) {
	for i := 1; i < storage.JumlahStartup; i++ {
		key := storage.Startups[i]
		j := i - 1
		for j >= 0 && ((asc && storage.Startups[j].TahunBerdiri > key.TahunBerdiri) || (!asc && storage.Startups[j].TahunBerdiri < key.TahunBerdiri)) {
			storage.Startups[j+1] = storage.Startups[j]
			j--
		}
		storage.Startups[j+1] = key
	}
}
