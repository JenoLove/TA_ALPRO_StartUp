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
	// Cek apakah admin sudah ada
	for i := 0; i < UserCount; i++ {
		if Users[i].Email == "admin@gmail.com" {
			return // Sudah ada, tidak usah ditambah
		}
	}

	// Tambah akun admin
	Users[UserCount] = User{
		Nama:     "Admin",
		Email:    "admin@gmail.com",
		Password: "admin123",
		Role:     "admin",
	}
	UserCount++
}
