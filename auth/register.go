package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"start-up/storage"
)

func Register() {
	if storage.UserCount >= 100 {
		fmt.Println("Pengguna sudah maksimum.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var nama, ttl, nohp, email, pass string
	var valid bool

	for {
		fmt.Print("Nama Lengkap: ")
		namaInput, _ := reader.ReadString('\n')
		nama = strings.TrimSpace(namaInput)

		if nama == "" {
			fmt.Println("Tidak boleh kosong.")
			continue
		}
		valid = true
		for _, ch := range nama {
			if !unicode.IsLetter(ch) && ch != ' ' {
				valid = false
				break
			}
		}
		if valid {
			break
		}
		fmt.Println("Nama hanya boleh berisi huruf dan spasi.")
	}

	for {
		fmt.Print("Tempat, tanggal lahir: ")
		ttlInput, _ := reader.ReadString('\n')
		ttl = strings.TrimSpace(ttlInput)
		if ttl != "" {
			break
		}
		fmt.Println("Tidak boleh kosong.")
	}

	for {
		fmt.Print("No HP: ")
		nohpInput, _ := reader.ReadString('\n')
		nohp = strings.TrimSpace(nohpInput)

		if nohp == "" {
			fmt.Println("Tidak boleh kosong.")
			continue
		}
		valid = true
		for _, ch := range nohp {
			if !unicode.IsDigit(ch) {
				valid = false
				break
			}
		}
		if valid {
			break
		}
		fmt.Println("No HP hanya boleh berisi angka.")
	}

	for {
		fmt.Print("Email: ")
		emailInput, _ := reader.ReadString('\n')
		email = strings.TrimSpace(emailInput)
		if email != "" {
			break
		}
		fmt.Println("Tidak boleh kosong.")
	}

	for {
		fmt.Print("Password: ")
		passInput, _ := reader.ReadString('\n')
		pass = strings.TrimSpace(passInput)

		if ValidasiPassword(pass) {
			break
		} else {
			fmt.Println("Password harus minimal 8 karakter dan mengandung huruf serta angka")
		}
	}

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
