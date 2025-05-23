package resource

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"start-up/storage"
)

func CariStartup() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== Pencarian Startup ===\n")
	fmt.Println("1. Cari berdasarkan Nama")
	fmt.Println("2. Cari berdasarkan Bidang Usaha")
	fmt.Println("3. Kembali")
	fmt.Println("Pilih jenis pencarian: \n")

	jenisInput, _ := reader.ReadString('\n')
	jenisInput = strings.TrimSpace(jenisInput)

	var jenis int
	_, err := fmt.Sscanf(jenisInput, "%d", &jenis)
	if err != nil || jenis < 1 || jenis > 3 {
		fmt.Println("Pilihan tidak valid")
		return
	}
	if jenis == 3 {
		return
	}

	fmt.Println("\nPilih metode pencarian:\n")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih metode: ")

	metodeInput, _ := reader.ReadString('\n')
	metodeInput = strings.TrimSpace(metodeInput)

	var metode int
	_, err = fmt.Sscanf(metodeInput, "%d", &metode)
	if err != nil || (metode != 1 && metode != 2) {
		fmt.Println("Pilihan harus 1 atau 2")
		return
	}

	var keyword string
	if jenis == 1 {
		fmt.Print("\nMasukkan nama startup: ")
	} else {
		fmt.Print("\nMasukkan bidang usaha: ")
	}
	keywordInput, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keywordInput)

	if metode == 2 {
		if jenis == 1 {
			SelectionSortNama(true)
		} else {
			SelectionSortBidang(true)
		}
	}

	var hasils [100]storage.Startup
	var count int

	if metode == 1 {
		count = SequentialSearch(jenis, keyword, &hasils)
	} else {
		count = BinarySearch(jenis, keyword, &hasils)
	}

	if count > 0 {
		CetakStartupTable(hasils, count)
	} else {
		fmt.Println("\nTidak ditemukan startup yang sesuai dengan kriteria pencarian.")
	}
}

func CetakStartupTable(startups [100]storage.Startup, count int) {
	const (
		idWidth     = 4
		namaWidth   = 22
		bidangWidth = 22
		tahunWidth  = 15
		danaWidth   = 12
	)

	printBorder := func() {
		fmt.Println("+" + strings.Repeat("-", idWidth+2) +
			"+" + strings.Repeat("-", namaWidth+2) +
			"+" + strings.Repeat("-", bidangWidth+2) +
			"+" + strings.Repeat("-", tahunWidth+2) +
			"+" + strings.Repeat("-", danaWidth+2) + "+")
	}

	printBorder()
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		idWidth, "ID",
		namaWidth, "Nama Startup",
		bidangWidth, "Bidang Usaha",
		tahunWidth, "Tahun Berdiri",
		danaWidth, "Total Dana")
	printBorder()

	for i := 0; i < count; i++ {
		startup := startups[i]

		fmt.Printf("| %-*d | %-*s | %-*s | %-*d | %*s |\n",
			idWidth, startup.ID,
			namaWidth, truncateString(startup.Nama, namaWidth),
			bidangWidth, truncateString(startup.BidangUsaha, bidangWidth),
			tahunWidth, startup.TahunBerdiri,
			danaWidth, format(startup.TotalDana))
	}

	printBorder()
}

func truncateString(str string, maxLength int) string {
	if len(str) > maxLength {
		return str[:maxLength-3] + "..."
	}
	return str
}

func format(amount int) string {
	str := fmt.Sprintf("%d", amount)
	n := len(str)
	if n <= 3 {
		return str
	}

	var hasil strings.Builder
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			hasil.WriteString(".")
		}
		hasil.WriteByte(str[i])
	}
	return hasil.String()
}

func SequentialSearch(jenis int, keyword string, hasils *[100]storage.Startup) int {
	keywordLower := strings.ToLower(keyword)
	var count int = 0

	for i := 0; i < storage.JumlahStartup; i++ {
		var field string
		if jenis == 1 {
			field = strings.ToLower(storage.Startups[i].Nama)
		} else {
			field = strings.ToLower(storage.Startups[i].BidangUsaha)
		}

		if strings.Contains(field, keywordLower) {
			hasils[count] = storage.Startups[i]
			count++
		}
	}

	return count
}

func BinarySearch(jenis int, keyword string, hasils *[100]storage.Startup) int {
	keywordLower := strings.ToLower(keyword)
	n := storage.JumlahStartup

	low, high := 0, n-1
	foundIndex := -1
	for low <= high {
		mid := (low + high) / 2
		var field string
		if jenis == 1 {
			field = strings.ToLower(storage.Startups[mid].Nama)
		} else {
			field = strings.ToLower(storage.Startups[mid].BidangUsaha)
		}

		if field >= keywordLower {
			foundIndex = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if foundIndex == -1 {
		return 0
	}

	count := 0
	stop := false
	i := foundIndex
	for i < n && !stop {
		var field string
		if jenis == 1 {
			field = strings.ToLower(storage.Startups[i].Nama)
		} else {
			field = strings.ToLower(storage.Startups[i].BidangUsaha)
		}

		if strings.Contains(field, keywordLower) {
			hasils[count] = storage.Startups[i]
			count++
		} else {
			stop = true
		}
		i++
	}

	return count
}

func SelectionSortNama(asc bool) {
	n := storage.JumlahStartup
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			compare := strings.Compare(storage.Startups[j].Nama, storage.Startups[extremeIdx].Nama)
			if (asc && compare < 0) || (!asc && compare > 0) {
				extremeIdx = j
			}
		}
		if extremeIdx != i {
			storage.Startups[i], storage.Startups[extremeIdx] = storage.Startups[extremeIdx], storage.Startups[i]
		}
	}
}

func SelectionSortBidang(asc bool) {
	n := storage.JumlahStartup
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			compare := strings.Compare(storage.Startups[j].BidangUsaha, storage.Startups[extremeIdx].BidangUsaha)
			if (asc && compare < 0) || (!asc && compare > 0) {
				extremeIdx = j
			}
		}
		if extremeIdx != i {
			storage.Startups[i], storage.Startups[extremeIdx] = storage.Startups[extremeIdx], storage.Startups[i]
		}
	}
}
