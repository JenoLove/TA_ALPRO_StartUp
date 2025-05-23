package resource

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"start-up/storage"
)

func TambahStartup() {
	if storage.JumlahStartup >= storage.MAX_STARTUP {
		fmt.Println("Kapasitas penuh.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var s storage.Startup
	s.ID = storage.JumlahStartup + 1

	// Tambahkan ini untuk buang newline sisa jika dipanggil setelah fmt.Scan()
	fmt.Print("")
	reader.ReadString('\n') // <-- ini penting banget!

	fmt.Print("Nama Startup: ")
	nama, _ := reader.ReadString('\n')
	s.Nama = strings.TrimSpace(nama)

	fmt.Print("Bidang Usaha: ")
	bidang, _ := reader.ReadString('\n')
	s.BidangUsaha = strings.TrimSpace(bidang)

	fmt.Print("Tahun Berdiri: ")
	tahunStr, _ := reader.ReadString('\n')
	tahunStr = strings.TrimSpace(tahunStr)
	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		fmt.Println("Tahun tidak valid.")
		return
	}
	s.TahunBerdiri = tahun

	fmt.Print("Total Dana: ")
	danaStr, _ := reader.ReadString('\n')
	danaStr = strings.TrimSpace(danaStr)
	dana, err := strconv.Atoi(danaStr)
	if err != nil {
		fmt.Println("Total dana tidak valid.")
		return
	}
	s.TotalDana = dana

	storage.Startups[storage.JumlahStartup] = s
	storage.JumlahStartup++
	fmt.Println("Startup berhasil ditambahkan.")
}

func UbahStartup() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// Tampilkan daftar startup terlebih dahulu
	fmt.Println("\nDaftar Startup:")
	for i := 0; i < storage.JumlahStartup; i++ {
		fmt.Printf("%d. %s (Bidang: %s)\n", storage.Startups[i].ID, storage.Startups[i].Nama, storage.Startups[i].BidangUsaha)
	}

	fmt.Print("\nMasukkan ID Startup yang akan diubah: ")
	var id int
	fmt.Scan(&id)
	reader.ReadString('\n') // Membersihkan buffer

	// Cari startup berdasarkan ID
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

	fmt.Println("\nData saat ini:")
	fmt.Printf("Nama: %s\n", found.Nama)
	fmt.Printf("Bidang Usaha: %s\n", found.BidangUsaha)
	fmt.Printf("Tahun Berdiri: %d\n", found.TahunBerdiri)
	fmt.Printf("Total Dana: %d\n", found.TotalDana)

	fmt.Println("\nMasukkan data baru (kosongkan jika tidak ingin mengubah):")

	fmt.Print("Nama Startup: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama != "" {
		found.Nama = nama
	}

	fmt.Print("Bidang Usaha: ")
	bidang, _ := reader.ReadString('\n')
	bidang = strings.TrimSpace(bidang)
	if bidang != "" {
		found.BidangUsaha = bidang
	}

	fmt.Print("Tahun Berdiri: ")
	tahunStr, _ := reader.ReadString('\n')
	tahunStr = strings.TrimSpace(tahunStr)
	if tahunStr != "" {
		tahun, err := strconv.Atoi(tahunStr)
		if err == nil {
			found.TahunBerdiri = tahun
		} else {
			fmt.Println("Tahun berdiri tidak valid, data tidak diubah.")
		}
	}

	fmt.Print("Total Dana: ")
	danaStr, _ := reader.ReadString('\n')
	danaStr = strings.TrimSpace(danaStr)
	if danaStr != "" {
		dana, err := strconv.Atoi(danaStr)
		if err == nil {
			found.TotalDana = dana
		} else {
			fmt.Println("Total dana tidak valid, data tidak diubah.")
		}
	}

	fmt.Printf("Startup dengan ID %d berhasil diupdate.\n", id)
}

func HapusStartup() {
	if storage.JumlahStartup == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}

	// Tampilkan daftar startup terlebih dahulu
	fmt.Println("\nDaftar Startup:")
	for i := 0; i < storage.JumlahStartup; i++ {
		fmt.Printf("%d. %s\n", storage.Startups[i].ID, storage.Startups[i].Nama)
	}

	fmt.Print("\nMasukkan ID Startup yang akan dihapus: ")
	var id int
	fmt.Scan(&id)

	// Cari index startup berdasarkan ID
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

	// Konfirmasi penghapusan
	fmt.Printf("Anda yakin ingin menghapus startup %s (ID: %d)? (y/n): ", storage.Startups[index].Nama, id)
	var confirm string
	fmt.Scan(&confirm)

	if strings.ToLower(confirm) == "y" {
		// Geser elemen array ke kiri mulai dari index yang dihapus
		for i := index; i < storage.JumlahStartup-1; i++ {
			storage.Startups[i] = storage.Startups[i+1]
			storage.Startups[i].ID = i + 1 // Update ID
		}

		// Kurangi jumlah startup
		storage.JumlahStartup--
		fmt.Println("Startup berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}
