package resource

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')

			var t storage.Tim
			fmt.Print("Nama Anggota: ")
			t.NamaAnggota, _ = reader.ReadString('\n')
			t.NamaAnggota = strings.TrimSpace(t.NamaAnggota)

			fmt.Print("Peran: ")
			t.Peran, _ = reader.ReadString('\n')
			t.Peran = strings.TrimSpace(t.Peran)

			storage.Startups[i].Tim[storage.Startups[i].JumlahAnggota] = t
			storage.Startups[i].JumlahAnggota++
			fmt.Println("Anggota ditambahkan.")
			return
		}
	}

	fmt.Println("Startup tidak ditemukan.")
}

func DaftarAnggotaTim() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada anggota yang terdaftar.")
		return
	}

	for i := 0; i < storage.JumlahStartup; i++ {
		s := storage.Startups[i]
		fmt.Printf("\n--- Startup: %s ---\n", s.Nama)

		if s.JumlahAnggota == 0 {
			fmt.Println("Belum ada anggota tim.")
			continue
		}

		noWidth := 4
		namaAnggotaWidth := 25
		peranWidth := 20

		totalWidth := noWidth + namaAnggotaWidth + peranWidth + 3*3 + 1

		fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")
		fmt.Printf("| %-*s | %-*s | %-*s |\n",
			noWidth, "No",
			namaAnggotaWidth, "Nama Anggota",
			peranWidth, "Peran")
		fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")

		for j := 0; j < s.JumlahAnggota; j++ {
			t := s.Tim[j]
			fmt.Printf("| %-*d | %-*s | %-*s |\n",
				noWidth, j+1,
				namaAnggotaWidth, t.NamaAnggota,
				peranWidth, t.Peran)
		}

		fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")
	}
}
