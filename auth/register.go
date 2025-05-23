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
		fmt.Println("Maksimum pengguna tercapai.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var nama, ttl, nohp, email, pass string
	var valid bool

	for {
		fmt.Print("Nama Lengkap: ")
		nama, _ = reader.ReadString('\n')
		nama = strings.TrimSpace(nama)

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
		fmt.Println("Nama hanya boleh berisi huruf.")
	}

	for {
		fmt.Print("Tempat, tanggal lahir: ")
		ttl, _ = reader.ReadString('\n')
		ttl = strings.TrimSpace(ttl)
		if ttl != "" {
			break
		}
		fmt.Println("Tidak boleh kosong.")
	}

	for {
		fmt.Print("No HP: ")
		nohp, _ = reader.ReadString('\n')
		nohp = strings.TrimSpace(nohp)

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
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if email != "" {
			break
		}
		fmt.Println("Tidak boleh kosong.")
	}

	for {
		fmt.Print("Password: ")
		pass, _ = reader.ReadString('\n')
		pass = strings.TrimSpace(pass)
		if len(pass) >= 8 {
			break
		}
		fmt.Println("Password minimal 8 karakter.")
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
