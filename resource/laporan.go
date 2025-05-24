package resource

import (
	"fmt"
	"strings"

	"start-up/storage"
)

func HitungJumlahBidang(startups [100]storage.Startup, jumlahStartup int) (bidang [100]string, jumlah [100]int, total int) {
	total = 0
	for i := 0; i < jumlahStartup; i++ {
		found := false
		j := 0
		for j < total && !found {
			if bidang[j] == startups[i].BidangUsaha {
				jumlah[j]++
				found = true
			}
			j++
		}
		if !found {
			bidang[total] = startups[i].BidangUsaha
			jumlah[total] = 1
			total++
		}
	}
	return
}

func CetakLaporanBidang(bidang [100]string, jumlah [100]int, total int) {
	noWidth := 4
	bidangWidth := 25
	jumlahWidth := 10

	totalTableWidth := noWidth + bidangWidth + jumlahWidth + 3*3 + 1

	judul := "=== LAPORAN PER BIDANG ===\n"
	spasiKiri := (totalTableWidth - len(judul)) / 2

	fmt.Println()
	fmt.Printf("%s%s\n", strings.Repeat(" ", spasiKiri), judul)
	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")
	fmt.Printf("| %-*s | %-*s | %-*s |\n",
		noWidth, "No",
		bidangWidth, "Bidang Usaha",
		jumlahWidth, "Jumlah")
	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")

	for i := 0; i < total; i++ {
		fmt.Printf("| %-*d | %-*s | %-*d |\n",
			noWidth, i+1,
			bidangWidth, bidang[i],
			jumlahWidth, jumlah[i])
	}

	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")
}

func LaporanBidang() {
	bidang, jumlah, total := HitungJumlahBidang(storage.Startups, storage.JumlahStartup)
	CetakLaporanBidang(bidang, jumlah, total)
}
