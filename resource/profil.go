package resource

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"start-up/auth"
	"start-up/storage"
)

func ProfilAdmin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		noWidth := 4
		namaWidth := 20
		emailWidth := 25
		roleWidth := 10
		statusWidth := 10
		totalWidth := noWidth + namaWidth + emailWidth + roleWidth + statusWidth + 6*3 + 1

		judul := "=== Manajemen User ==="
		spasiKiri := (totalWidth - len(judul)) / 2
		fmt.Printf("\n%s%s\n\n", strings.Repeat(" ", spasiKiri), judul)

		if storage.UserCount == 0 {
			fmt.Println("Tidak ada user yang terdaftar.")
			return
		}

		fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")
		fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s |\n",
			noWidth, "No",
			namaWidth, "Nama",
			emailWidth, "Email",
			roleWidth, "Role",
			statusWidth, "Aktif")
		fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")

		for i := 0; i < storage.UserCount; i++ {
			user := storage.Users[i]
			aktif := false
			if user.IsActive != nil {
				aktif = *user.IsActive
			}
			fmt.Printf("| %-*d | %-*s | %-*s | %-*s | %-*v |\n",
				noWidth, i+1,
				namaWidth, user.Nama,
				emailWidth, user.Email,
				roleWidth, user.Role,
				statusWidth, aktif)
		}
		fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")

		fmt.Print("\nPilih user untuk dinonaktifkan/aktifkan (0 untuk kembali): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		pilihan, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input harus berupa angka.")
			continue
		}

		if pilihan == 0 {
			return
		}

		if pilihan < 1 || pilihan > storage.UserCount {
			fmt.Println("Data tidak valid.")
			continue
		}

		target := &storage.Users[pilihan-1]

		if target.IsActive == nil {
			statusBaru := false
			target.IsActive = &statusBaru
		} else {
			statusBaru := !(*target.IsActive)
			target.IsActive = &statusBaru
		}

		status := "dinonaktifkan"
		if *target.IsActive {
			status = "diaktifkan"
		}
		fmt.Printf("User %s berhasil %s.\n", target.Nama, status)
	}
}

func TampilkanProfilKaryawan() {
	reader := bufio.NewReader(os.Stdin)
	u := auth.CurrentUser

	noWidth := 4
	namaWidth := 20
	ttlWidth := 15
	noHPWidth := 15
	emailWidth := 30
	totalWidth := noWidth + namaWidth + ttlWidth + noHPWidth + emailWidth + 6*3 + 1

	judul := "=== Profil Saya ==="
	spasiKiri := (totalWidth - len(judul)) / 2
	fmt.Printf("\n%s%s\n\n", strings.Repeat(" ", spasiKiri), judul)

	fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		noWidth, "No",
		namaWidth, "Nama",
		ttlWidth, "TTL",
		noHPWidth, "No HP",
		emailWidth, "Email",
	)
	fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")

	fmt.Printf("| %-*d | %-*s | %-*s | %-*s | %-*s |\n",
		noWidth, 1,
		namaWidth, strings.TrimSpace(u.Nama),
		ttlWidth, strings.TrimSpace(u.TTL),
		noHPWidth, strings.TrimSpace(u.NoHP),
		emailWidth, strings.TrimSpace(u.Email),
	)

	fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")

	fmt.Print("\nApakah anda yakin ingin mengubah data? (y/n): ")
	jawaban, _ := reader.ReadString('\n')
	jawaban = strings.TrimSpace(strings.ToLower(jawaban))

	if jawaban == "y" || jawaban == "ya" {
		fmt.Print("Nama baru: ")
		u.Nama, _ = reader.ReadString('\n')

		fmt.Print("TTL baru: ")
		u.TTL, _ = reader.ReadString('\n')

		fmt.Print("No HP baru: ")
		u.NoHP, _ = reader.ReadString('\n')

		fmt.Print("Email baru: ")
		u.Email, _ = reader.ReadString('\n')

		u.Nama = strings.TrimSpace(u.Nama)
		u.TTL = strings.TrimSpace(u.TTL)
		u.NoHP = strings.TrimSpace(u.NoHP)
		u.Email = strings.TrimSpace(u.Email)

		fmt.Println("\nData berhasil diperbarui.")
	}
}
