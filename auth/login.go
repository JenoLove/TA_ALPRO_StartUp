package auth

import (
	"bufio"
	"fmt"
	"os"
	"start-up/storage"
	"strings"
)

var CurrentUser *storage.User = nil

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	for i := 0; i < storage.UserCount; i++ {
		storedEmail := strings.TrimSpace(storage.Users[i].Email)
		storedPass := strings.TrimSpace(storage.Users[i].Password)

		if storedEmail == email && storedPass == password {
			CurrentUser = &storage.Users[i]
			return
		}
	}

	fmt.Println("Email atau Password salah.")
	CurrentUser = nil
}
