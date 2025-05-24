package resource

import (
	"fmt"
	"strings"

	"start-up/storage"
)

func TampilStartup(startups [100]storage.Startup, jumlah int) {
	const (
		idWidth    = 4
		namaWidth  = 22
		danaWidth  = 22
		tahunWidth = 15
	)

	totalTableWidth := idWidth + namaWidth + danaWidth + tahunWidth + 6*3 + 1

	judul := "=== DAFTAR START UP ===\n"
	padding := (totalTableWidth - len(judul)) / 2
	fmt.Println("\n" + strings.Repeat(" ", padding) + judul)

	fmt.Println("+" + strings.Repeat("-", totalTableWidth-2) + "+")

	printBorder := func() {
		fmt.Println("+" +
			strings.Repeat("-", idWidth+2) + "+" +
			strings.Repeat("-", namaWidth+2) + "+" +
			strings.Repeat("-", danaWidth+2) + "+" +
			strings.Repeat("-", tahunWidth+2) + "+")
	}

	printBorder()
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s |\n",
		idWidth, "ID",
		namaWidth, "Nama Startup",
		danaWidth, "Total Dana",
		tahunWidth, "Tahun Berdiri")
	printBorder()

	for i := 0; i < jumlah; i++ {
		startup := startups[i]
		id := fmt.Sprintf("%d", startup.ID)
		nama := truncateString(startup.Nama, namaWidth)
		dana := format(startup.TotalDana)
		tahun := fmt.Sprintf("%d", startup.TahunBerdiri)

		fmt.Printf("| %-*s | %-*s | %-*s | %-*s |\n",
			idWidth, id,
			namaWidth, nama,
			danaWidth, dana,
			tahunWidth, tahun)
	}

	printBorder()
}
