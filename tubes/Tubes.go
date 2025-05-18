package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Sentiment string

const (
	Positif Sentiment = "Positif"
	Negatif Sentiment = "Negatif"
	Netral  Sentiment = "Netral"
)

type Komentar struct {
	Teks     string
	Sentimen Sentiment
}

var positifKeywords = []string{"bagus", "keren", "suka", "unik", "hebat", "mantap"}
var negatifKeywords = []string{"jelek", "buruk", "benci", "aneh", "sampah", "gagal"}

var komentarList []Komentar

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Aplikasi Analisis Sentimen ===")
		fmt.Println("1. Tambah Komentar")
		fmt.Println("2. Ubah Komentar")
		fmt.Println("3. Hapus Komentar")
		fmt.Println("4. Tampilkan Semua Komentar")
		fmt.Println("5. Cari Komentar (Sequential)")
		fmt.Println("6. Cari Komentar (Binary)")
		fmt.Println("7. Urutkan Berdasarkan Panjang Komentar")
		fmt.Println("8. Urutkan Berdasarkan Sentimen")
		fmt.Println("9. Tampilkan Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			tambahKomentar(scanner)
		case "2":
			ubahKomentar(scanner)
		case "3":
			hapusKomentar(scanner)
		case "4":
			tampilkanKomentar()
		case "5":
			cariKomentarSequential(scanner)
		case "6":
			cariKomentarBinary(scanner)
		case "7":
			fmt.Print("Ascending (y/n)? ")
			scanner.Scan()
			asc := strings.ToLower(scanner.Text()) == "y"
			urutkanPanjangKomentar(asc)
		case "8":
			urutkanSentimen()
		case "9":
			tampilkanStatistik()
		case "0":
			fmt.Println("Sampai jumpa, Dipi!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahKomentar(scanner *bufio.Scanner) {
	fmt.Print("Masukkan komentar: ")
	scanner.Scan()
	teks := scanner.Text()
	sentimen := analisisSentimen(teks)
	komentar := Komentar{Teks: teks, Sentimen: sentimen}
	komentarList = append(komentarList, komentar)
	fmt.Println("Komentar ditambahkan dengan sentimen:", sentimen)
}

func ubahKomentar(scanner *bufio.Scanner) {
	tampilkanKomentar()
	fmt.Print("Masukkan indeks komentar yang ingin diubah: ")
	var index int
	fmt.Scan(&index)
	if index < 0 || index >= len(komentarList) {
		fmt.Println("Indeks tidak valid.")
		return
	}
	fmt.Print("Masukkan komentar baru: ")
	scanner.Scan()
	teks := scanner.Text()
	sentimen := analisisSentimen(teks)
	komentarList[index] = Komentar{Teks: teks, Sentimen: sentimen}
	fmt.Println("Komentar diperbarui.")
}

func hapusKomentar(scanner *bufio.Scanner) {
	tampilkanKomentar()
	fmt.Print("Masukkan indeks komentar yang ingin dihapus: ")
	var index int
	fmt.Scan(&index)
	if index < 0 || index >= len(komentarList) {
		fmt.Println("Indeks tidak valid.")
		return
	}
	komentarList = append(komentarList[:index], komentarList[index+1:]...)
	fmt.Println("Komentar dihapus.")
}

func tampilkanKomentar() {
	if len(komentarList) == 0 {
		fmt.Println("Belum ada komentar.")
		return
	}
	fmt.Println("\nDaftar Komentar:")
	for i, k := range komentarList {
		fmt.Printf("[%d] \"%s\" → Sentimen: %s\n", i, k.Teks, k.Sentimen)
	}
}

func analisisSentimen(teks string) Sentiment {
	lower := strings.ToLower(teks)
	for _, kata := range positifKeywords {
		if strings.Contains(lower, kata) {
			return Positif
		}
	}
	for _, kata := range negatifKeywords {
		if strings.Contains(lower, kata) {
			return Negatif
		}
	}
	return Netral
}

func cariKomentarSequential(scanner *bufio.Scanner) {
	fmt.Print("Masukkan keyword yang ingin dicari: ")
	scanner.Scan()
	keyword := strings.ToLower(scanner.Text())

	ditemukan := false
	for i, k := range komentarList {
		if strings.Contains(strings.ToLower(k.Teks), keyword) {
			fmt.Printf("[%d] \"%s\" → Sentimen: %s\n", i, k.Teks, k.Sentimen)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Komentar tidak ditemukan.")
	}
}

func cariKomentarBinary(scanner *bufio.Scanner) {
	sort.Slice(komentarList, func(i, j int) bool {
		return komentarList[i].Teks < komentarList[j].Teks
	})

	fmt.Print("Masukkan keyword yang ingin dicari (tepat): ")
	scanner.Scan()
	keyword := strings.ToLower(scanner.Text())

	low, high := 0, len(komentarList)-1
	for low <= high {
		mid := (low + high) / 2
		current := strings.ToLower(komentarList[mid].Teks)
		if current == keyword {
			fmt.Printf("[%d] \"%s\" → Sentimen: %s\n", mid, komentarList[mid].Teks, komentarList[mid].Sentimen)
			return
		} else if keyword < current {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Komentar tidak ditemukan.")
}

func urutkanPanjangKomentar(ascending bool) {
	n := len(komentarList)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if len(komentarList[j].Teks) < len(komentarList[idx].Teks) {
					idx = j
				}
			} else {
				if len(komentarList[j].Teks) > len(komentarList[idx].Teks) {
					idx = j
				}
			}
		}
		komentarList[i], komentarList[idx] = komentarList[idx], komentarList[i]
	}
	fmt.Println("Komentar berhasil diurutkan berdasarkan panjang.")
}

func urutkanSentimen() {
	nilaiSentimen := func(s Sentiment) int {
		switch s {
		case Positif:
			return 1
		case Netral:
			return 2
		case Negatif:
			return 3
		}
		return 4
	}

	for i := 1; i < len(komentarList); i++ {
		key := komentarList[i]
		j := i - 1
		for j >= 0 && nilaiSentimen(komentarList[j].Sentimen) > nilaiSentimen(key.Sentimen) {
			komentarList[j+1] = komentarList[j]
			j--
		}
		komentarList[j+1] = key
	}
	fmt.Println("Komentar berhasil diurutkan berdasarkan sentimen.")
}

func tampilkanStatistik() {
	jmlPos, jmlNet, jmlNeg := 0, 0, 0
	for _, k := range komentarList {
		switch k.Sentimen {
		case Positif:
			jmlPos++
		case Netral:
			jmlNet++
		case Negatif:
			jmlNeg++
		}
	}
	fmt.Println("\nStatistik Sentimen Komentar:")
	fmt.Printf("Positif: %d\n", jmlPos)
	fmt.Printf("Netral : %d\n", jmlNet)
	fmt.Printf("Negatif: %d\n", jmlNeg)
}