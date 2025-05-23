package resource

import (
	"fmt"

	"start-up/auth"
)

func TampilkanMenu() {
	for {
		fmt.Println("\n--- Aplikasi Manajemen Startup ---")

		// Menu untuk Admin
		menuNumber := 1
		if auth.CurrentUser != nil && auth.CurrentUser.Role == "admin" {
			fmt.Printf("%d. Tambah Startup\n", menuNumber)
			menuNumber++
			fmt.Printf("%d. Ubah Startup\n", menuNumber)
			menuNumber++
			fmt.Printf("%d. Hapus Startup\n", menuNumber)
			menuNumber++
			fmt.Printf("%d. Tambah Anggota Tim\n", menuNumber)
			menuNumber++
		}

		// Menu untuk semua user
		fmt.Printf("%d. Cari Startup\n", menuNumber)
		menuNumber++
		fmt.Printf("%d. Urutkan Startup\n", menuNumber)
		menuNumber++
		fmt.Printf("%d. Laporan Startup per Bidang\n", menuNumber)
		menuNumber++
		fmt.Printf("%d. Keluar\n", menuNumber)

		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		// Reset menuNumber untuk pengecekan
		menuNumber = 1
		if auth.CurrentUser != nil && auth.CurrentUser.Role == "admin" {
			switch pilihan {
			case menuNumber:
				TambahStartup()
				continue
			case menuNumber + 1:
				UbahStartup()
				continue
			case menuNumber + 2:
				HapusStartup()
				continue
			case menuNumber + 3:
				TambahAnggotaTim()
				continue
			}
			menuNumber += 4
		}

		switch pilihan {
		case menuNumber:
			CariStartup()
		case menuNumber + 1:
			UrutkanStartup()
		case menuNumber + 2:
			LaporanBidang()
		case menuNumber + 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
