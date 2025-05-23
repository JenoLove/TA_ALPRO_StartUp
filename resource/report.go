package resource

import (
	"fmt"

	"start-up/storage"
)

func LaporanBidang() {
	var bidang [100]string
	var jumlah [100]int
	var total int = 0

	for i := 0; i < storage.JumlahStartup; i++ {
		found := false
		for j := 0; j < total; j++ {
			if bidang[j] == storage.Startups[i].BidangUsaha {
				jumlah[j]++
				found = true
			}
		}
		if !found {
			bidang[total] = storage.Startups[i].BidangUsaha
			jumlah[total] = 1
			total++
		}
	}

	fmt.Println("\n--- Laporan Startup per Bidang ---")
	for i := 0; i < total; i++ {
		fmt.Printf("%s: %d startup\n", bidang[i], jumlah[i])
	}
}
