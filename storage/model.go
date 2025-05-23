package storage

type Tim struct {
	NamaAnggota string
	Peran       string
}

type Startup struct {
	ID            int
	Nama          string
	BidangUsaha   string
	TahunBerdiri  int
	TotalDana     int
	Tim           [5]Tim
	JumlahAnggota int
}

type User struct {
	Nama       string
	TTL        string
	NoHP       string
	Email      string
	Password   string
	Jabatan    string
	Departemen string
	Role       string
}

func InitAdmin() {
	for i := 0; i < UserCount; i++ {
		if Users[i].Email == "admin@gmail.com" {
			return
		}
	}

	Users[UserCount] = User{
		Nama:     "Admin",
		Email:    "admin@gmail.com",
		Password: "admin123",
		Role:     "admin",
	}
	UserCount++
}
