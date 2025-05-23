package resource

import (
	"fmt"
	"start-up/storage"
	"strings"
)

func TampilkanStartup(startups []storage.Startup) {
	const (
		idWidth         = 4
		namaWidth       = 22
		danaWidth       = 22
		tahunWidth      = 15
		totalTableWidth = idWidth + namaWidth + tahunWidth + danaWidth + 12
	)

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

	for _, startup := range startups {
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
