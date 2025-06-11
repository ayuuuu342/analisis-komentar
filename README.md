# analisis-komentar
Aplikasi Analisis Sentimen Komentar Media Sosial
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// deklarasi tipe data
type Sentimen string

const (
	Positif Sentimen = "Positif"
	Negatif Sentimen = "Negatif"
	Netral  Sentimen = "Netral"
)

type Komentar struct {
	Teks     string
	Sentimen Sentimen
}

var positifKeywords = []string{"bagus", "keren", "suka", "aneh", "hebat", "mantap"}
var negatifKeywords = []string{"jelek", "buruk", "benci", "unik", "sampah", "gagal"}
var komentarList []Komentar

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Inisialisasi data komentar dari "database"
	komentarList = []Komentar{
		{"Aplikasi ini sangat bagus dan mantap!", analisisSentimen("Aplikasi ini sangat bagus dan mantap!")},
		{"Saya benci tampilannya.", analisisSentimen("Saya benci tampilannya.")},
		{"Tampilannya cukup unik.", analisisSentimen("Tampilannya cukup unik.")},
		{"Tidak terlalu buruk.", analisisSentimen("Tidak terlalu buruk.")},
		{"Layanan pelanggan sangat hebat!", analisisSentimen("Layanan pelanggan sangat hebat!")},
	}

	// meminta nama pengguna
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	// menampilkan menu
	fmt.Printf("Selamat datang, %s!\n", nama)
	for {
		fmt.Println("\n=== Aplikasi Analisis Sentimen ===")
		fmt.Println("1. Ubah Komentar")
		fmt.Println("2. Hapus Komentar")
		fmt.Println("3. Tampilkan Semua Komentar")
		fmt.Println("4. Cari Komentar (Sequential)")
		fmt.Println("5. Cari Komentar (Binary)")
		fmt.Println("6. Urutkan Berdasarkan Panjang Komentar")
		fmt.Println("7. Urutkan Berdasarkan Sentimen")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			ubahKomentar(scanner)
		case "2":
			hapusKomentar(scanner)
		case "3":
			tampilkanKomentar()
		case "4":
			cariKomentarSequential(scanner)
		case "5":
			cariKomentarBinary(scanner)
		case "6":
			fmt.Print("Ascending (y/n)? ")
			scanner.Scan()
			asc := strings.ToLower(scanner.Text()) == "y"
			urutkanPanjangKomentar(asc)
		case "7":
			urutkanSentimen()
		case "8":
			tampilkanStatistik()
		case "0":
			fmt.Printf("Sampai jumpa, %s!\n", nama)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// analisis sentimen berdasarkan keyword
func analisisSentimen(teks string) Sentimen {
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

// ubah komentar berdasarkan indeks
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

// hapus komentar berdasarkan indeks
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

// tampilkan semua komentar
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

// pencarian komentar secara sequential
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

// pencarian binary (tepat)
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

// urutkan komentar berdasarkan panjang (selection sort)
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

// urutkan komentar berdasarkan sentimen (insertion sort)
func urutkanSentimen() {
	nilaiSentimen := func(s Sentimen) int {
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

// tampilkan statistik sentimen
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
