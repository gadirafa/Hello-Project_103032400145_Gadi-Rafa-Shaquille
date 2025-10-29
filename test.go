#INI HASIL KOMENTAR

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const NMAX = 10

type Pinjaman struct {
	NamaPeminjam      string
	JumlahPinjaman    float64
	LamaPinjamanBulan int
	SukuBungaTahunan  float64
}

type TabPinjaman [NMAX]Pinjaman

func main() {
	var daftarPinjaman TabPinjaman
	var jumlahData int
	var pilihan int

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Input Data Pinjaman")
		fmt.Println("2. Tampilkan Seluruh Data")
		fmt.Println("3. Sorting Data")
		fmt.Println("4. Searching Nama Peminjam")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			jumlahData = inputData(&daftarPinjaman)
		case 2:
			tampilkanData(daftarPinjaman, jumlahData)
		case 3:
			menuSorting(&daftarPinjaman, jumlahData)
		case 4:
			menuSearching(daftarPinjaman, jumlahData)
		case 5:
			fmt.Println("Terima kasih! Program selesai.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inputData(data *TabPinjaman) int {
	var n int
	fmt.Print("Masukkan jumlah data pinjaman (maks 10): ")
	fmt.Scanln(&n)
	if n > NMAX {
		n = NMAX
		fmt.Println("Jumlah melebihi maksimum, hanya 10 data yang akan diterima.")
	}

	for i := 0; i < n; i++ {
		fmt.Printf("\nData pinjaman ke-%d:\n", i+1)
		fmt.Print("Nama Peminjam        : ")
		fmt.Scanln(&data[i].NamaPeminjam)
		fmt.Print("Jumlah Pinjaman (Rp) : ")
		fmt.Scanln(&data[i].JumlahPinjaman)
		fmt.Print("Lama Pinjaman (bulan): ")
		fmt.Scanln(&data[i].LamaPinjamanBulan)
		fmt.Print("Suku Bunga Tahunan % : ")
		fmt.Scanln(&data[i].SukuBungaTahunan)
	}
	return n
}

func tampilkanData(data TabPinjaman, n int) {
	if n == 0 {
		fmt.Println("Belum ada data pinjaman.")
		return
	}
	fmt.Println("\n=== DATA PINJAMAN ===")
	for i := 0; i < n; i++ {
		fmt.Printf("\nData ke-%d\n", i+1)
		fmt.Printf("Nama Peminjam        : %s\n", data[i].NamaPeminjam)
		fmt.Printf("Jumlah Pinjaman      : Rp %.2f\n", data[i].JumlahPinjaman)
		fmt.Printf("Lama Pinjaman (bulan): %d\n", data[i].LamaPinjamanBulan)
		fmt.Printf("Suku Bunga Tahunan   : %.2f%%\n", data[i].SukuBungaTahunan)
	}
}

func menuSorting(data *TabPinjaman, n int) {
	if n == 0 {
		fmt.Println("Data kosong, tidak bisa disorting.")
		return
	}
	fmt.Println("\n=== MENU SORTING ===")
	fmt.Println("1. Selection Sort (Jumlah Pinjaman)")
	fmt.Println("2. Insertion Sort (Nama Peminjam)")
	fmt.Print("Pilih metode (1/2): ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		selectionSort(data, n)
		fmt.Println("Data berhasil diurutkan berdasarkan jumlah pinjaman (terkecil ke terbesar).")
	case 2:
		insertionSort(data, n)
		fmt.Println("Data berhasil diurutkan berdasarkan nama peminjam (A-Z).")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func selectionSort(data *TabPinjaman, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].JumlahPinjaman < data[minIdx].JumlahPinjaman {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func insertionSort(data *TabPinjaman, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && strings.ToLower(data[j].NamaPeminjam) > strings.ToLower(key.NamaPeminjam) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func menuSearching(data TabPinjaman, n int) {
	if n == 0 {
		fmt.Println("Data kosong, tidak bisa dilakukan pencarian.")
		return
	}
	fmt.Println("\n=== MENU SEARCHING ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (otomatis urut A-Z)")
	fmt.Print("Pilih metode (1/2): ")
	var pilihan int
	fmt.Scanln(&pilihan)

	fmt.Print("Masukkan nama peminjam yang dicari: ")
	var target string
	fmt.Scanln(&target)

	switch pilihan {
	case 1:
		idx := sequentialSearch(data, n, target)
		if idx != -1 {
			fmt.Println("Data ditemukan:")
			tampilkanPinjaman(data[idx])
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	case 2:
		sorted := data
		sort.SliceStable(sorted[:n], func(i, j int) bool {
			return strings.ToLower(sorted[i].NamaPeminjam) < strings.ToLower(sorted[j].NamaPeminjam)
		})
		idx := binarySearch(sorted, n, target)
		if idx != -1 {
			fmt.Println("Data ditemukan:")
			tampilkanPinjaman(sorted[idx])
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func sequentialSearch(data TabPinjaman, n int, target string) int {
	for i := 0; i < n; i++ {
		if strings.EqualFold(data[i].NamaPeminjam, target) {
			return i
		}
	}
	return -1
}

func binarySearch(data TabPinjaman, n int, target string) int {
	low, high := 0, n-1
	target = strings.ToLower(target)
	for low <= high {
		mid := (low + high) / 2
		guess := strings.ToLower(data[mid].NamaPeminjam)
		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func tampilkanPinjaman(p Pinjaman) {
	fmt.Printf("Nama Peminjam        : %s\n", p.NamaPeminjam)
	fmt.Printf("Jumlah Pinjaman      : Rp %.2f\n", p.JumlahPinjaman)
	fmt.Printf("Lama Pinjaman (bulan): %d\n", p.LamaPinjamanBulan)
	fmt.Printf("Suku Bunga Tahunan   : %.2f%%\n", p.SukuBungaTahunan)
}
