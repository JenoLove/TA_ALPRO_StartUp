package auth

import (
	"bufio"
	"fmt"
	"os"
	"start-up/storage"
	"strings"
)

func ForgotPassword() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Email Anda: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	found := false
	for i := 0; i < storage.UserCount; i++ {
		storedEmail := strings.TrimSpace(storage.Users[i].Email)
		if storedEmail == email {
			found = true
			fmt.Print("Masukkan Password Baru: ")
			newPass, _ := reader.ReadString('\n')
			newPass = strings.TrimSpace(newPass)

			storage.Users[i].Password = newPass
			fmt.Println("Password berhasil diubah.")
			break
		}
	}

	if !found {
		fmt.Println("Email tidak ditemukan.")
	}
}
