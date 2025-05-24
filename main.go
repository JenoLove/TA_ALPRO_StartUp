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
		fmt.Println("\n === APLIKASI MANAJEMEN STARTUP ===\n")
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
			fmt.Print("email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			auth.CurrentUser = auth.Login(email, password)
			if auth.CurrentUser != nil {
				resource.Menu()
			}
		case "3":
			auth.LupaPassword()
		case "4":
			fmt.Println("Anda telah K=keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
