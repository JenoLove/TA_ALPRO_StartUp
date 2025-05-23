package resource

import (
	"fmt"

	"start-up/storage"
)

func TampilkanStartup() {
	for i := 0; i < storage.JumlahStartup; i++ {
		fmt.Printf("ID: %d | Nama: %s | Bidang: %s | Tahun: %d | Dana: %d | Anggota: %d\n",
			storage.Startups[i].ID,
			storage.Startups[i].Nama,
			storage.Startups[i].BidangUsaha,
			storage.Startups[i].TahunBerdiri,
			storage.Startups[i].TotalDana,
			storage.Startups[i].JumlahAnggota)
	}
}
