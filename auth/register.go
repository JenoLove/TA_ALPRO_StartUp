package auth

import (
	"bufio"
	"fmt"
	"os"
	"start-up/storage"
)

func Register() {
	if storage.UserCount >= 100 {
		fmt.Println("Maksimum pengguna tercapai.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Lengkap: ")
	nama, _ := reader.ReadString('\n')
	fmt.Print("Tempat, tanggal lahir: ")
	ttl, _ := reader.ReadString('\n')
	fmt.Print("No HP: ")
	nohp, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	fmt.Print("Password: ")
	pass, _ := reader.ReadString('\n')
	storage.Users[storage.UserCount] = storage.User{
		Nama:     nama,
		TTL:      ttl,
		NoHP:     nohp,
		Email:    email,
		Password: pass,
		Role:     "karyawan",
	}

	storage.UserCount++
	fmt.Println("Registrasi berhasil!")
}
