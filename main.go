package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"start-up/auth"
	"start-up/resource"
	"start-up/storage"
)

func main() {
	storage.InitAdmin()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== APLIKASI MANAJEMEN STARTUP ===")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Lupa Password")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			auth.Register()
		case "2":
			auth.Login()
			if auth.CurrentUser != nil {
				resource.TampilkanMenu()
			}
		case "3":
			auth.ForgotPassword()
		case "4":
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
