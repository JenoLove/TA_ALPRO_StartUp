package auth

import (
	"fmt"
	"strings"

	"start-up/storage"
)

var CurrentUser *storage.User = nil

func Login(email string, password string) *storage.User {
	for i := 0; i < storage.UserCount; i++ {
		storedEmail := strings.TrimSpace(storage.Users[i].Email)
		storedPass := strings.TrimSpace(storage.Users[i].Password)

		if storedEmail == email && storedPass == password {
			if storage.Users[i].IsActive != nil && !(*storage.Users[i].IsActive) {
				fmt.Println("User dinonaktifkan.")
				return nil
			}

			CurrentUser = &storage.Users[i]
			return CurrentUser
		}
	}

	fmt.Println("Email atau Password salah.")
	return nil
}
