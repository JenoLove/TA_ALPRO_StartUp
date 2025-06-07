package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"start-up/storage"
)

func LupaPassword() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email Anda: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	found := false
	for i := 0; i < storage.UserCount; i++ {
		storedEmail := strings.TrimSpace(storage.Users[i].Email)
		if storedEmail == email {
			found = true

			var passBaru string
			for {
				fmt.Print("Password Baru: ")
				passBaruInput, _ := reader.ReadString('\n')
				passBaru = strings.TrimSpace(passBaruInput)

				if ValidasiPassword(passBaru) {
					break
				} else {
					fmt.Println("Password harus minimal 8 karakter dan mengandung huruf serta angka")
				}
			}

			storage.Users[i].Password = passBaru
			fmt.Println("Password berhasil diubah.")
			break
		}
	}

	if !found {
		fmt.Println("Email tidak ditemukan.")
	}
}

func ValidasiPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasText := false
	hasDigit := false

	for _, ch := range password {
		if unicode.IsLetter(ch) {
			hasText = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}

	return hasText && hasDigit
}
