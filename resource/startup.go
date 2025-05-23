package resource

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"start-up/storage"
)

func InputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func InputInt(prompt string) (int, bool) {
	input := InputString(prompt)
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input tidak valid.")
		return 0, false
	}
	return value, true
}

func TampilkanTabelStartup(s storage.Startup) {
	fmt.Println(strings.Repeat("-", 85))
	fmt.Printf("| %-4s | %-25s | %-20s | %-15s | %-12s |\n", "ID", "Nama Startup", "Bidang Usaha", "Tahun Berdiri", "Total Dana")
	fmt.Println(strings.Repeat("-", 85))
	fmt.Printf("| %-4d | %-25s | %-20s | %-15d | %-12d |\n", s.ID, s.Nama, s.BidangUsaha, s.TahunBerdiri, s.TotalDana)
	fmt.Println(strings.Repeat("-", 85))
}

func TambahStartup() {
	if storage.JumlahStartup >= storage.MAX_STARTUP {
		fmt.Println("Kapasitas penuh.")
		return
	}

	var s storage.Startup
	s.ID = storage.JumlahStartup + 1

	for {
		s.Nama = InputString("Nama Startup: ")
		if s.Nama != "" {
			break
		}
		fmt.Println("Nama tidak boleh kosong.")
	}

	for {
		s.BidangUsaha = InputString("Bidang Usaha: ")
		if s.BidangUsaha != "" {
			break
		}
		fmt.Println("Bidang Usaha tidak boleh kosong.")
	}

	for {
		tahun, ok := InputInt("Tahun Berdiri: ")
		if ok {
			s.TahunBerdiri = tahun
			break
		}
		fmt.Println("Tahun harus berupa angka.")
	}

	for {
		dana, ok := InputInt("Total Dana: ")
		if ok {
			s.TotalDana = dana
			break
		}
		fmt.Println("Total Dana harus berupa angka.")
	}

	storage.Startups[storage.JumlahStartup] = s
	storage.JumlahStartup++
	fmt.Println("Startup berhasil ditambahkan.")
}

func DaftarStartup() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}

	fmt.Println("\n--- Daftar Startup ---")

	idWidth := 4
	namaWidth := 20
	bidangWidth := 20
	tahunWidth := 15
	danaWidth := 12

	totalTableWidth := idWidth + namaWidth + bidangWidth + tahunWidth + danaWidth + 6*3 + 1

	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")

	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		idWidth, "ID",
		namaWidth, "Nama Startup",
		bidangWidth, "Bidang Usaha",
		tahunWidth, "Tahun Berdiri",
		danaWidth, "Total Dana")

	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")

	for i := 0; i < storage.JumlahStartup; i++ {
		s := storage.Startups[i]
		fmt.Printf("| %-*d | %-*s | %-*s | %-*d | %-*d |\n",
			idWidth, s.ID,
			namaWidth, s.Nama,
			bidangWidth, s.BidangUsaha,
			tahunWidth, s.TahunBerdiri,
			danaWidth, s.TotalDana)
	}

	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")
}

func UbahStartup() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}

	DaftarStartup()

	id, ok := InputInt("\nMasukkan ID Startup yang akan diubah: ")
	if !ok {
		return
	}

	var found *storage.Startup
	for i := 0; i < storage.JumlahStartup; i++ {
		if storage.Startups[i].ID == id {
			found = &storage.Startups[i]
			break
		}
	}

	if found == nil {
		fmt.Println("Startup dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n--- Data Saat Ini ---")
	TampilkanTabelStartup(*found)

	for {
		nama := InputString("Nama Startup: ")
		if nama != "" {
			found.Nama = nama
			break
		}
		fmt.Println("Nama tidak boleh kosong.")
	}

	for {
		bidang := InputString("Bidang Usaha: ")
		if bidang != "" {
			found.BidangUsaha = bidang
			break
		}
		fmt.Println("Bidang Usaha tidak boleh kosong.")
	}

	for {
		tahun, ok := InputInt("Tahun Berdiri: ")
		if ok {
			found.TahunBerdiri = tahun
			break
		}
		fmt.Println("Tahun harus berupa angka.")
	}

	for {
		dana, ok := InputInt("Total Dana: ")
		if ok {
			found.TotalDana = dana
			break
		}
		fmt.Println("Total Dana harus berupa angka.")
	}

	fmt.Printf("Startup dengan ID %d berhasil diupdate.\n", id)
}

func HapusStartup() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}

	DaftarStartup()

	id, ok := InputInt("\nMasukkan ID Startup yang akan dihapus: ")
	if !ok {
		return
	}

	index := -1
	for i := 0; i < storage.JumlahStartup; i++ {
		if storage.Startups[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Startup dengan ID tersebut tidak ditemukan.")
		return
	}

	confirm := InputString(fmt.Sprintf("Anda yakin ingin menghapus startup %s (ID: %d)? (y/n): ", storage.Startups[index].Nama, id))
	if strings.ToLower(confirm) == "y" {
		for j := index; j < storage.JumlahStartup-1; j++ {
			storage.Startups[j] = storage.Startups[j+1]
			storage.Startups[j].ID = j + 1
		}
		storage.JumlahStartup--
		fmt.Println("Startup berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}
