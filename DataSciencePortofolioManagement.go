package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================================
// KONSTANTA, TIPE DATA, DAN VARIABEL GLOBAL
// ============================================================

const NMAX int = 100

type ArProyek [NMAX]Proyek

type Proyek struct {
	nama             string
	kategori         int
	teknologi        string
	deskripsi        string
	tingkatKesulitan int
	tanggal          int
	bulan            int
	tahun            int
}

var daftarProyek ArProyek
var jumlahProyek int = 0
var sc = bufio.NewScanner(os.Stdin)

// ============================================================
// FUNGSI INPUT DAN HELPER STRING
// ============================================================

func bacaBaris() string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func bacaInt() int {
	teks := bacaBaris()
	angka, err := strconv.Atoi(teks)
	if err != nil {
		return 0
	}
	return angka
}

func strSama(a, b string) bool {
	return strings.EqualFold(a, b)
}

func namaProyekSudahAda(nama string) bool {
	for i := 0; i < jumlahProyek; i++ {
		if strSama(daftarProyek[i].nama, nama) {
			return true
		}
	}
	return false
}

func namaKategori(k int) string {
	if k == 1 {
		return "Machine Learning"
	} else if k == 2 {
		return "Data Visualization"
	} else if k == 3 {
		return "Data Analysis"
	} else if k == 4 {
		return "NLP"
	} else if k == 5 {
		return "Computer Vision"
	}
	return "Tidak Diketahui"
}

// ============================================================
// MENU APLIKASI
// ============================================================

func menuUtama(pilihan *int) {
	fmt.Println("==============================================")
	fmt.Println("  Aplikasi Manajemen Portofolio Data Science")
	fmt.Println("==============================================")
	fmt.Println("1. Tampilkan Data Proyek")
	fmt.Println("2. Kelola Proyek (Tambah/Edit/Hapus)")
	fmt.Println("3. Cari Proyek")
	fmt.Println("4. Dashboard Portofolio")
	fmt.Println("5. Keluar")
	fmt.Println("==============================================")
	fmt.Print("Pilihan: ")
	*pilihan = bacaInt()
	for *pilihan < 1 || *pilihan > 5 {
		fmt.Print("Pilihan tidak valid. Pilihan: ")
		*pilihan = bacaInt()
	}
	fmt.Println()
}

func menuTampilkan(pilihan *int) {
	fmt.Println("============================================")
	fmt.Println("           TAMPILKAN DATA PROYEK")
	fmt.Println("============================================")
	fmt.Println("1. Tampilkan Semua Proyek")
	fmt.Println("2. Urutkan berdasarkan Tingkat Kesulitan (Selection Sort)")
	fmt.Println("3. Urutkan berdasarkan Tanggal (Insertion Sort)")
	fmt.Println("4. Statistik Proyek per Kategori")
	fmt.Println("5. Tampilkan Daftar Keahlian")
	fmt.Println("6. Kembali")
	fmt.Println("============================================")
	fmt.Print("Pilihan: ")
	*pilihan = bacaInt()
	for *pilihan < 1 || *pilihan > 6 {
		fmt.Print("Pilihan tidak valid. Pilihan: ")
		*pilihan = bacaInt()
	}
	fmt.Println()
}

func menuKelola(pilihan *int) {
	fmt.Println("============================================")
	fmt.Println("             KELOLA PROYEK")
	fmt.Println("============================================")
	fmt.Println("1. Tambah Proyek")
	fmt.Println("2. Edit Proyek")
	fmt.Println("3. Hapus Proyek")
	fmt.Println("4. Kembali")
	fmt.Println("============================================")
	fmt.Print("Pilihan: ")
	*pilihan = bacaInt()
	for *pilihan < 1 || *pilihan > 4 {
		fmt.Print("Pilihan tidak valid. Pilihan: ")
		*pilihan = bacaInt()
	}
	fmt.Println()
}

func menuCari(pilihan *int) {
	fmt.Println("============================================")
	fmt.Println("               CARI PROYEK")
	fmt.Println("============================================")
	fmt.Println("1. Cari berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Cari berdasarkan Kategori (Binary Search)")
	fmt.Println("3. Kembali")
	fmt.Println("============================================")
	fmt.Print("Pilihan: ")
	*pilihan = bacaInt()
	for *pilihan < 1 || *pilihan > 3 {
		fmt.Print("Pilihan tidak valid. Pilihan: ")
		*pilihan = bacaInt()
	}
	fmt.Println()
}

// ============================================================
// VALIDASI INPUT
// ============================================================

func validasiTanggal(tanggal, bulan, tahun int) bool {
	if tahun < 2000 || tahun > 2100 {
		return false
	}
	if bulan < 1 || bulan > 12 {
		return false
	}
	if tanggal < 1 {
		return false
	}
	hariMaks := 31
	if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		hariMaks = 30
	} else if bulan == 2 {
		if tahun%400 == 0 || (tahun%100 != 0 && tahun%4 == 0) {
			hariMaks = 29
		} else {
			hariMaks = 28
		}
	}
	return tanggal <= hariMaks
}

func validasiKesulitan(k int) bool {
	return k >= 1 && k <= 5
}

func inputKategori() int {
	fmt.Println("Pilih Kategori:")
	fmt.Println("1. Machine Learning")
	fmt.Println("2. Data Visualization")
	fmt.Println("3. Data Analysis")
	fmt.Println("4. NLP")
	fmt.Println("5. Computer Vision")
	fmt.Print("Kategori (1-5): ")
	k := bacaInt()
	for k < 1 || k > 5 {
		fmt.Println("Kategori tidak valid!")
		fmt.Print("Kategori (1-5): ")
		k = bacaInt()
	}
	return k
}

func inputKesulitan() int {
	fmt.Print("Tingkat Kesulitan (1-5): ")
	k := bacaInt()
	for !validasiKesulitan(k) {
		fmt.Println("Tingkat kesulitan harus antara 1 sampai 5!")
		fmt.Print("Tingkat Kesulitan (1-5): ")
		k = bacaInt()
	}
	return k
}

func inputTanggal(tgl, bln, thn *int) {
	fmt.Print("Tanggal Dibuat (contoh: 31 12 2026): ")
	baris := bacaBaris()
	n, _ := fmt.Sscan(baris, tgl, bln, thn)
	for n != 3 || !validasiTanggal(*tgl, *bln, *thn) {
		fmt.Println("Tanggal tidak valid! Masukkan kembali.")
		fmt.Print("Tanggal Dibuat (contoh: 31 12 2026): ")
		baris = bacaBaris()
		n, _ = fmt.Sscan(baris, tgl, bln, thn)
	}
}

// ============================================================
// MENAMPILKAN DATA PROYEK
// ============================================================

func tampilSatuProyek(no int, p Proyek) {
	fmt.Println("--------------------------------------------")
	fmt.Printf("%d. Nama      : %s\n", no, p.nama)
	fmt.Printf("   Kategori  : %s\n", namaKategori(p.kategori))
	fmt.Printf("   Teknologi : %s\n", p.teknologi)
	fmt.Printf("   Deskripsi : %s\n", p.deskripsi)
	fmt.Printf("   Kesulitan : %d/5\n", p.tingkatKesulitan)
	fmt.Printf("   Tanggal   : %02d/%02d/%d\n", p.tanggal, p.bulan, p.tahun)
}

// Menampilkan seluruh proyek menggunakan rekursi
func tampilRekursif(idx int) {

	if idx >= jumlahProyek {
		return
	}

	tampilSatuProyek(
		idx+1,
		daftarProyek[idx],
	)

	tampilRekursif(idx + 1)
}

func tampilSemuaProyek() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek. Silakan tambah proyek terlebih dahulu.")
		fmt.Println()
		return
	}
	fmt.Println("============================================")
	fmt.Println("           DAFTAR SEMUA PROYEK")
	fmt.Println("============================================")
	tampilRekursif(0)
	fmt.Println("--------------------------------------------")
	fmt.Printf("Total: %d proyek\n", jumlahProyek)
	fmt.Println()
}

// ============================================================
// CRUD PROYEK
// Tambah, Edit, dan Hapus Data Proyek
// ============================================================

func tambahProyek() {
	if jumlahProyek >= NMAX {
		fmt.Println("Data proyek sudah penuh!")
		fmt.Println()
		return
	}

	var lanjut int = 1
	for lanjut == 1 {
		fmt.Println("============================================")
		fmt.Println("          <> Tambah Proyek Baru <>")
		fmt.Println("============================================")

		fmt.Print("Nama Proyek: ")
		nama := bacaBaris()

		for strings.TrimSpace(nama) == "" {
			fmt.Println("Nama proyek tidak boleh kosong!")
			fmt.Print("Nama Proyek: ")
			nama = bacaBaris()
		}

		for namaProyekSudahAda(nama) {
			fmt.Println("Nama proyek sudah ada!")
			fmt.Print("Nama Proyek: ")
			nama = bacaBaris()
		}

		daftarProyek[jumlahProyek].nama = nama

		daftarProyek[jumlahProyek].kategori = inputKategori()

		fmt.Print("Teknologi (pisah koma, contoh: Python,TensorFlow): ")
		teknologi := bacaBaris()

		for strings.TrimSpace(teknologi) == "" {
			fmt.Println("Teknologi tidak boleh kosong!")
			fmt.Print("Teknologi (pisah koma, contoh: Python,TensorFlow): ")
			teknologi = bacaBaris()
		}

		daftarProyek[jumlahProyek].teknologi = teknologi

		fmt.Print("Deskripsi Singkat: ")
		deskripsi := bacaBaris()

		for strings.TrimSpace(deskripsi) == "" {
			fmt.Println("Deskripsi tidak boleh kosong!")
			fmt.Print("Deskripsi Singkat: ")
			deskripsi = bacaBaris()
		}

		daftarProyek[jumlahProyek].deskripsi = deskripsi

		daftarProyek[jumlahProyek].tingkatKesulitan = inputKesulitan()

		inputTanggal(
			&daftarProyek[jumlahProyek].tanggal,
			&daftarProyek[jumlahProyek].bulan,
			&daftarProyek[jumlahProyek].tahun,
		)

		jumlahProyek++
		fmt.Println("============================================")
		fmt.Println("Proyek berhasil ditambahkan!")
		fmt.Println("============================================")

		if jumlahProyek >= NMAX {
			fmt.Println("Data proyek sudah penuh.")
			lanjut = 2
		} else {
			fmt.Print("Lanjut tambah proyek?\n1. Ya   2. Tidak\nPilihan: ")
			lanjut = bacaInt()
			for lanjut < 1 || lanjut > 2 {
				fmt.Print("Pilihan: ")
				lanjut = bacaInt()
			}
		}
		fmt.Println()
	}
}

func editProyek() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek untuk diedit.")
		fmt.Println()
		return
	}

	tampilSemuaProyek()

	var lanjut int = 1
	for lanjut == 1 {
		fmt.Print("Masukkan nomor proyek yang ingin diedit: ")
		idx := bacaInt() - 1

		if idx < 0 || idx >= jumlahProyek {
			fmt.Println("Nomor tidak valid!")
		} else {
			fmt.Println("============================================")
			fmt.Println("Pilih field yang ingin diubah:")
			fmt.Println("1. Nama Proyek")
			fmt.Println("2. Kategori")
			fmt.Println("3. Teknologi")
			fmt.Println("4. Deskripsi")
			fmt.Println("5. Tingkat Kesulitan")
			fmt.Println("6. Tanggal")
			fmt.Println("7. Ganti Semua")
			fmt.Println("============================================")
			fmt.Print("Pilihan: ")
			pil := bacaInt()
			for pil < 1 || pil > 7 {
				fmt.Print("Pilihan: ")
				pil = bacaInt()
			}

			if pil == 1 {

				fmt.Print("Nama Proyek baru: ")
				namaBaru := bacaBaris()

				for strings.TrimSpace(namaBaru) == "" {
					fmt.Println("Nama proyek tidak boleh kosong!")
					fmt.Print("Nama Proyek baru: ")
					namaBaru = bacaBaris()
				}

				for namaProyekSudahAda(namaBaru) &&
					!strSama(namaBaru, daftarProyek[idx].nama) {

					fmt.Println("Nama proyek sudah ada!")
					fmt.Print("Nama Proyek baru: ")
					namaBaru = bacaBaris()
				}

				daftarProyek[idx].nama = namaBaru
			} else if pil == 2 {
				daftarProyek[idx].kategori = inputKategori()
			} else if pil == 3 {
				fmt.Print("Teknologi baru: ")
				daftarProyek[idx].teknologi = bacaBaris()
			} else if pil == 4 {
				fmt.Print("Deskripsi baru: ")
				daftarProyek[idx].deskripsi = bacaBaris()
			} else if pil == 5 {
				daftarProyek[idx].tingkatKesulitan = inputKesulitan()
			} else if pil == 6 {
				inputTanggal(&daftarProyek[idx].tanggal, &daftarProyek[idx].bulan, &daftarProyek[idx].tahun)
			} else {
				fmt.Print("Nama Proyek baru: ")
				namaBaru := bacaBaris()

				for namaProyekSudahAda(namaBaru) &&
					!strSama(namaBaru, daftarProyek[idx].nama) {

					fmt.Println("Nama proyek sudah ada!")
					fmt.Print("Nama Proyek baru: ")
					namaBaru = bacaBaris()
				}

				daftarProyek[idx].nama = namaBaru

				daftarProyek[idx].kategori = inputKategori()

				fmt.Print("Teknologi baru: ")
				daftarProyek[idx].teknologi = bacaBaris()

				fmt.Print("Deskripsi baru: ")
				daftarProyek[idx].deskripsi = bacaBaris()

				daftarProyek[idx].tingkatKesulitan = inputKesulitan()

				inputTanggal(
					&daftarProyek[idx].tanggal,
					&daftarProyek[idx].bulan,
					&daftarProyek[idx].tahun,
				)
			}
			fmt.Println("Data berhasil diubah!")
		}

		fmt.Print("Lanjut edit proyek lain?\n1. Ya   2. Tidak\nPilihan: ")
		lanjut = bacaInt()
		for lanjut < 1 || lanjut > 2 {
			fmt.Print("Pilihan: ")
			lanjut = bacaInt()
		}
		fmt.Println()
	}
}

func hapusProyek() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek untuk dihapus.")
		fmt.Println()
		return
	}

	tampilSemuaProyek()

	var lanjut int = 1
	for lanjut == 1 {
		fmt.Print("Masukkan nomor proyek yang ingin dihapus: ")
		idx := bacaInt() - 1

		if idx < 0 || idx >= jumlahProyek {
			fmt.Println("Nomor tidak valid!")
		} else {
			namaHapus := daftarProyek[idx].nama
			for i := idx; i < jumlahProyek-1; i++ {
				daftarProyek[i] = daftarProyek[i+1]
			}
			daftarProyek[jumlahProyek-1] = Proyek{}
			jumlahProyek--
			fmt.Printf("Proyek \"%s\" berhasil dihapus!\n", namaHapus)
		}

		fmt.Print("Lanjut hapus proyek lain?\n1. Ya   2. Tidak\nPilihan: ")
		lanjut = bacaInt()
		for lanjut < 1 || lanjut > 2 {
			fmt.Print("Pilihan: ")
			lanjut = bacaInt()
		}
		fmt.Println()
	}
}

// ============================================================
// PENCARIAN DATA
// Sequential Search dan Binary Search
// ============================================================

// Sequential Search parsial berdasarkan nama proyek
func sequentialSearchNama(keyword string) int {
	for i := 0; i < jumlahProyek; i++ {

		nama := strings.ToLower(daftarProyek[i].nama)
		cari := strings.ToLower(keyword)

		if strings.Contains(nama, cari) {
			return i
		}
	}
	return -1
}

func cariNama() {
	fmt.Print("Masukkan nama proyek yang dicari: ")
	keyword := bacaBaris()
	fmt.Println()

	idx := sequentialSearchNama(keyword)
	if idx == -1 {
		fmt.Printf("Proyek dengan nama \"%s\" tidak ditemukan.\n", keyword)
	} else {
		fmt.Println("Proyek ditemukan:")
		tampilSatuProyek(idx+1, daftarProyek[idx])
		fmt.Println("--------------------------------------------")
	}
	fmt.Println()
}

func sortKategoriUntukBinarySearch(arr *ArProyek, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].kategori < arr[minIdx].kategori {
				minIdx = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = temp
	}
}

// Binary Search rekursif berdasarkan kategori proyek
func binarySearchKategori(arr ArProyek, kiri, kanan, keyword int) int {

	if kiri > kanan {
		return -1
	}

	tengah := (kiri + kanan) / 2

	if arr[tengah].kategori == keyword {
		return tengah
	}

	if keyword < arr[tengah].kategori {
		return binarySearchKategori(
			arr,
			kiri,
			tengah-1,
			keyword,
		)
	}

	return binarySearchKategori(
		arr,
		tengah+1,
		kanan,
		keyword,
	)
}

func cariKategori() {
	keyword := inputKategori()
	fmt.Println()

	var salinan ArProyek
	for i := 0; i < jumlahProyek; i++ {
		salinan[i] = daftarProyek[i]
	}
	sortKategoriUntukBinarySearch(&salinan, jumlahProyek)

	idx := binarySearchKategori(salinan, 0, jumlahProyek-1, keyword)

	if idx == -1 {
		fmt.Printf("Tidak ada proyek dengan kategori \"%s\".\n", namaKategori(keyword))
	} else {
		fmt.Printf("Proyek dengan kategori \"%s\":\n", namaKategori(keyword))
		fmt.Println("============================================")

		kiri := idx
		for kiri > 0 && salinan[kiri-1].kategori == keyword {
			kiri--
		}

		kanan := idx
		for kanan < jumlahProyek-1 &&
			salinan[kanan+1].kategori == keyword {
			kanan++
		}

		no := 1
		for i := kiri; i <= kanan; i++ {
			tampilSatuProyek(no, salinan[i])
			no++
		}

		fmt.Println("--------------------------------------------")
		fmt.Printf("Ditemukan %d proyek.\n", no-1)
	}
}

// ============================================================
// PENGURUTAN DATA
// Selection Sort dan Insertion Sort
// ============================================================

// Selection Sort berdasarkan tingkat kesulitan proyek
func selectionSortKesulitan(arr *ArProyek, n int, asc bool) {
	for i := 0; i < n-1; i++ {
		pilihanIdx := i
		for j := i + 1; j < n; j++ {
			if asc {
				if arr[j].tingkatKesulitan < arr[pilihanIdx].tingkatKesulitan {
					pilihanIdx = j
				}
			} else {
				if arr[j].tingkatKesulitan > arr[pilihanIdx].tingkatKesulitan {
					pilihanIdx = j
				}
			}
		}
		temp := arr[i]
		arr[i] = arr[pilihanIdx]
		arr[pilihanIdx] = temp
	}
}

func lebihBesar(a, b Proyek, asc bool) bool {
	if asc {
		if a.tahun != b.tahun {
			return a.tahun > b.tahun
		}
		if a.bulan != b.bulan {
			return a.bulan > b.bulan
		}
		return a.tanggal > b.tanggal
	}
	if a.tahun != b.tahun {
		return a.tahun < b.tahun
	}
	if a.bulan != b.bulan {
		return a.bulan < b.bulan
	}
	return a.tanggal < b.tanggal
}

// Insertion Sort berdasarkan tanggal pembuatan proyek
func insertionSortTanggal(arr *ArProyek, n int, asc bool) {
	for i := 1; i < n; i++ {
		kunci := arr[i]
		j := i - 1
		for j >= 0 && lebihBesar(arr[j], kunci, asc) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = kunci
	}
}

func urutKesulitan() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek.")
		fmt.Println()
		return
	}

	var pilArah int
	fmt.Println("Urutan:")
	fmt.Println("1. Ascending (termudah ke tersulit)")
	fmt.Println("2. Descending (tersulit ke termudah)")
	fmt.Print("Pilihan: ")
	pilArah = bacaInt()
	for pilArah < 1 || pilArah > 2 {
		fmt.Print("Pilihan: ")
		pilArah = bacaInt()
	}
	fmt.Println()

	var salinan ArProyek
	for i := 0; i < jumlahProyek; i++ {
		salinan[i] = daftarProyek[i]
	}

	asc := pilArah == 1
	selectionSortKesulitan(&salinan, jumlahProyek, asc)

	if asc {
		fmt.Println("============================================")
		fmt.Println("  Proyek (Termudah -> Tersulit) [Selection Sort]")
	} else {
		fmt.Println("============================================")
		fmt.Println("  Proyek (Tersulit -> Termudah) [Selection Sort]")
	}
	fmt.Println("============================================")
	for i := 0; i < jumlahProyek; i++ {
		tampilSatuProyek(i+1, salinan[i])
	}
	fmt.Println("--------------------------------------------")
	fmt.Println()
}

func urutTanggal() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek.")
		fmt.Println()
		return
	}

	var pilArah int
	fmt.Println("Urutan:")
	fmt.Println("1. Ascending (terlama ke terbaru)")
	fmt.Println("2. Descending (terbaru ke terlama)")
	fmt.Print("Pilihan: ")
	pilArah = bacaInt()
	for pilArah < 1 || pilArah > 2 {
		fmt.Print("Pilihan: ")
		pilArah = bacaInt()
	}
	fmt.Println()

	var salinan ArProyek
	for i := 0; i < jumlahProyek; i++ {
		salinan[i] = daftarProyek[i]
	}

	asc := pilArah == 1
	insertionSortTanggal(&salinan, jumlahProyek, asc)

	if asc {
		fmt.Println("============================================")
		fmt.Println("  Proyek (Terlama -> Terbaru) [Insertion Sort]")
	} else {
		fmt.Println("============================================")
		fmt.Println("  Proyek (Terbaru -> Terlama) [Insertion Sort]")
	}
	fmt.Println("============================================")
	for i := 0; i < jumlahProyek; i++ {
		tampilSatuProyek(i+1, salinan[i])
	}
	fmt.Println("--------------------------------------------")
	fmt.Println()
}

// Menentukan proyek termudah dan tersulit (Min-Max)
func minMaxKesulitan() {

	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	min := 0
	max := 0

	for i := 1; i < jumlahProyek; i++ {

		if daftarProyek[i].tingkatKesulitan <
			daftarProyek[min].tingkatKesulitan {

			min = i
		}

		if daftarProyek[i].tingkatKesulitan >
			daftarProyek[max].tingkatKesulitan {

			max = i
		}
	}

	fmt.Println("Proyek Termudah:")
	tampilSatuProyek(1, daftarProyek[min])

	fmt.Println()

	fmt.Println("Proyek Tersulit:")
	tampilSatuProyek(1, daftarProyek[max])
}

// ============================================================
// STATISTIK DAN ANALISIS PORTOFOLIO
// ============================================================

func statistikKategori() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek.")
		fmt.Println()
		return
	}

	var jumlah [6]int
	for i := 0; i < jumlahProyek; i++ {
		jumlah[daftarProyek[i].kategori]++
	}

	fmt.Println("============================================")
	fmt.Println("      STATISTIK PROYEK PER KATEGORI")
	fmt.Println("============================================")
	for k := 1; k <= 5; k++ {
		fmt.Printf("%-20s : %d proyek\n", namaKategori(k), jumlah[k])
	}
	fmt.Println("--------------------------------------------")
	fmt.Printf("%-20s : %d proyek\n", "Total", jumlahProyek)
	fmt.Println()

	maks := 1

	for i := 2; i <= 5; i++ {
		if jumlah[i] > jumlah[maks] {
			maks = i
		}
	}

	mins := -1

	for i := 1; i <= 5; i++ {
		if jumlah[i] > 0 {
			mins = i
			break
		}
	}

	for i := 1; i <= 5; i++ {
		if jumlah[i] > 0 && jumlah[i] < jumlah[mins] {
			mins = i
		}
	}

	fmt.Println()
	fmt.Println("Kategori Terbanyak :",
		namaKategori(maks))

	fmt.Println("Kategori Tersedikit :",
		namaKategori(mins))
}

func tampilKeahlian() {
	if jumlahProyek == 0 {
		fmt.Println("Belum ada proyek.")
		fmt.Println()
		return
	}

	var keahlianList [NMAX]string
	var jumlahKeahlian int = 0

	for i := 0; i < jumlahProyek; i++ {
		tek := daftarProyek[i].teknologi
		start := 0
		for idx := 0; idx <= len(tek); idx++ {
			if idx == len(tek) || tek[idx] == ',' {
				if idx > start {
					t := strings.TrimSpace(tek[start:idx])
					sudahAda := false
					for k := 0; k < jumlahKeahlian && !sudahAda; k++ {
						if strSama(keahlianList[k], t) {
							sudahAda = true
						}
					}
					if !sudahAda && jumlahKeahlian < NMAX {
						keahlianList[jumlahKeahlian] = t
						jumlahKeahlian++
					}
				}
				start = idx + 1
			}
		}
	}

	fmt.Println("============================================")
	fmt.Println("    DAFTAR KEAHLIAN DARI SEMUA PROYEK")
	fmt.Println("============================================")
	for i := 0; i < jumlahKeahlian; i++ {
		fmt.Printf("%d. %s\n", i+1, keahlianList[i])
	}
	fmt.Printf("\nTotal keahlian: %d\n", jumlahKeahlian)
	fmt.Println()
}

// Dashboard ringkasan portofolio pengguna
func dashboardPortofolio() {

	fmt.Println("================================")
	fmt.Println("      DASHBOARD PORTOFOLIO")
	fmt.Println("================================")

	fmt.Println("Total Proyek:",
		jumlahProyek)

	statistikKategori()

	fmt.Println()

	minMaxKesulitan()
}

// ============================================================
// PROGRAM UTAMA
// ============================================================

func main() {
	fmt.Println("=================================================")
	fmt.Println("  Aplikasi Manajemen Portofolio Data Science")
	fmt.Println("       untuk Pemula - Tugas Besar AlPro")
	fmt.Println("=================================================")
	fmt.Println("    Guidomelvin			103052530008   ")
	fmt.Println("    Ryan Azrya Altriyananda	103052500098   ")

	var pilihanUtama int
	var pilihanSub int

	menuUtama(&pilihanUtama)

	for pilihanUtama != 5 {
		if pilihanUtama == 1 {
			menuTampilkan(&pilihanSub)
			if pilihanSub == 1 {
				tampilSemuaProyek()
			} else if pilihanSub == 2 {
				urutKesulitan()
			} else if pilihanSub == 3 {
				urutTanggal()
			} else if pilihanSub == 4 {
				statistikKategori()
			} else if pilihanSub == 5 {
				tampilKeahlian()
			}
		} else if pilihanUtama == 2 {
			menuKelola(&pilihanSub)
			if pilihanSub == 1 {
				tambahProyek()
			} else if pilihanSub == 2 {
				editProyek()
			} else if pilihanSub == 3 {
				hapusProyek()
			}
		} else if pilihanUtama == 3 {
			menuCari(&pilihanSub)
			if pilihanSub == 1 {
				cariNama()
			} else if pilihanSub == 2 {
				cariKategori()
			}
		} else if pilihanUtama == 4 {
			dashboardPortofolio()
		}

		menuUtama(&pilihanUtama)
	}

	fmt.Println("================================================")
	fmt.Println("  Terima kasih telah menggunakan aplikasi ini!")
	fmt.Println("================================================")
}
