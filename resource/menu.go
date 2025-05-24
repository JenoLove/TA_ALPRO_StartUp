package resource

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"start-up/auth"
)

func Menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== APLIKASI MANAJEMEN START-UP ===\n")

		menuNumber := 1
		fmt.Printf("%d. Daftar StartUp\n", menuNumber)
		menuNumber++

		if auth.CurrentUser != nil && auth.CurrentUser.Role == "admin" {
			fmt.Printf("%d. Tambah Startup\n", menuNumber)
			menuNumber++
			fmt.Printf("%d. Ubah Startup\n", menuNumber)
			menuNumber++
			fmt.Printf("%d. Hapus Startup\n", menuNumber)
			menuNumber++
		}

		fmt.Printf("%d. Cari Startup\n", menuNumber)
		menuNumber++
		fmt.Printf("%d. Urutkan Startup\n", menuNumber)
		menuNumber++
		fmt.Printf("%d. Laporan Startup per Bidang\n", menuNumber)
		menuNumber++
		fmt.Printf("%d. Daftar Anggota Tim\n", menuNumber)
		menuNumber++

		if auth.CurrentUser != nil && auth.CurrentUser.Role == "admin" {
			fmt.Printf("%d. Tambah Anggota Tim\n", menuNumber)
			menuNumber++
		}

		fmt.Printf("%d. Profil Saya\n", menuNumber)
		menuNumber++

		fmt.Printf("%d. Keluar\n", menuNumber)

		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		pilihan, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Pilihan harus berupa angka")
			continue
		}

		menuNumber = 1
		var handled bool = false

		if pilihan == menuNumber {
			DaftarStartup()
			handled = true
		}

		menuNumber++
		if !handled && auth.CurrentUser != nil && auth.CurrentUser.Role == "admin" {
			if pilihan == menuNumber {
				TambahStartup()
				handled = true
			}
			menuNumber++
			if !handled && pilihan == menuNumber {
				UbahStartup()
				handled = true
			}
			menuNumber++
			if !handled && pilihan == menuNumber {
				HapusStartup()
				handled = true
			}
			menuNumber++
		}

		if !handled && pilihan == menuNumber {
			CariStartup()
			handled = true
		}
		menuNumber++

		if !handled && pilihan == menuNumber {
			UrutkanStartup()
			handled = true
		}
		menuNumber++

		if !handled && pilihan == menuNumber {
			LaporanBidang()
			handled = true
		}
		menuNumber++

		if !handled && pilihan == menuNumber {
			DaftarAnggotaTim()
			handled = true
		}
		menuNumber++

		if !handled && auth.CurrentUser != nil && auth.CurrentUser.Role == "admin" {
			if pilihan == menuNumber {
				TambahAnggotaTim()
				handled = true
			}
			menuNumber++
		}

		if !handled && pilihan == menuNumber {
			switch auth.CurrentUser.Role {
			case "admin":
				ProfilAdmin()
			case "karyawan":
				TampilkanProfilKaryawan()
			default:
				fmt.Println("Peran tidak dikenali.")
			}
			handled = true
		}
		menuNumber++

		if !handled && pilihan == menuNumber {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		}

		if !handled {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
