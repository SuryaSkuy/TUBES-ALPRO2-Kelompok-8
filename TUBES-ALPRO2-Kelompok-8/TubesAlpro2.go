package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --- Variabel Global ---
// Menggunakan array statis (bukan slice) untuk menyimpan data utama seperti pengguna, teman, status, dan komentar.
// Variabel global hanya digunakan untuk array utama seperti `pengguna`, yang memuat semua data pengguna.
const MAKS_PENGGUNA = 20
const MAKS_TEMAN = 20
const MAKS_STATUS = 20
const MAKS_KOMENTAR = 20

type Pengguna struct {
	NamaPengguna string
	KataSandi    string
	Profil       string
	Teman        [MAKS_TEMAN]string
	Status       [MAKS_STATUS]string
	Komentar     [MAKS_STATUS][MAKS_KOMENTAR]string
	JumlahTeman  int
	JumlahStatus int
}

var pengguna [MAKS_PENGGUNA]Pengguna
var jumlahPengguna int
var indeksMasuk = -1

// Fungsi Utilitas
// `cariIndeksPengguna` menggunakan algoritma pencarian sequential untuk mencari pengguna berdasarkan nama pengguna.
// Fungsi ini modular dan digunakan di beberapa bagian aplikasi.
func cariIndeksPengguna(namaPengguna string) int {
	// Loop iterasi linear untuk mencari indeks pengguna berdasarkan nama pengguna.
	for i := 0; i < jumlahPengguna; i++ {
		if pengguna[i].NamaPengguna == namaPengguna {
			return i
		}
	}
	return -1
}

// Fungsi `masuk` digunakan untuk memverifikasi kredensial pengguna saat masuk.
// Parameter:
// - namaPengguna: string yang berisi nama pengguna yang ingin masuk.
// - kataSandi: string yang berisi kata sandi pengguna yang ingin masuk.
// Mengembalikan indeks pengguna jika kredensial valid, atau -1 jika tidak valid.
func masuk(namaPengguna, kataSandi string) int {
	for i := 0; i < jumlahPengguna; i++ {
		// Memeriksa apakah nama pengguna dan kata sandi cocok dengan data yang ada.
		if pengguna[i].NamaPengguna == namaPengguna && pengguna[i].KataSandi == kataSandi {
			return i // Mengembalikan indeks pengguna jika cocok.
		}
	}
	return -1 // Mengembalikan -1 jika tidak ada kecocokan.
}

// Fungsi Pengurutan
// `selectionSortPenggunaByNamaPengguna` menggunakan algoritma selection sort untuk mengurutkan data pengguna.
// Mendukung urutan ascending dan descending sesuai parameter.
func selectionSortPenggunaByNamaPengguna(urutan string) {
	for i := 0; i < jumlahPengguna-1; i++ {
		idx := i
		for j := i + 1; j < jumlahPengguna; j++ {
			// Menggunakan kondisi untuk menentukan urutan berdasarkan parameter `urutan`.
			if (urutan == "asc" && pengguna[j].NamaPengguna < pengguna[idx].NamaPengguna) || (urutan == "desc" && pengguna[j].NamaPengguna > pengguna[idx].NamaPengguna) {
				idx = j
			}
		}
		if idx != i {
			// Tukar elemen jika diperlukan.
			temp := pengguna[i]
			pengguna[i] = pengguna[idx]
			pengguna[idx] = temp
		}
	}
}

// Fungsi `insertionSortTeman` mengurutkan array teman menggunakan algoritma Insertion Sort.
// Parameter:
// - teman: array string yang berisi nama-nama teman yang akan diurutkan.
// - jumlah: jumlah elemen dalam array teman yang akan diurutkan.
// - urutan: string yang menentukan urutan pengurutan, bisa "asc" untuk ascending atau "desc" untuk descending.
func insertionSortTeman(teman []string, jumlah int, urutan string) {
	for i := 1; i < jumlah; i++ {
		kunci := teman[i] // Menyimpan nilai elemen saat ini sebagai kunci.
		j := i - 1
		if urutan == "asc" {
			// Memindahkan elemen yang lebih besar ke posisi berikutnya untuk ascending.
			for j >= 0 && teman[j] > kunci {
				teman[j+1] = teman[j]
				j--
			}
		} else if urutan == "desc" {
			// Memindahkan elemen yang lebih kecil ke posisi berikutnya untuk descending.
			for j >= 0 && teman[j] < kunci {
				teman[j+1] = teman[j]
				j--
			}
		}
		teman[j+1] = kunci // Menempatkan kunci pada posisi yang tepat.
	}
}

// Fungsi Inti
// --- Subprogram Modular ---
// Setiap fitur diimplementasikan dalam fungsi modular dengan parameter yang jelas.
// Contoh: Fungsi `daftar` menangani proses registrasi pengguna baru.
func daftar() {
	if jumlahPengguna >= MAKS_PENGGUNA {
		fmt.Println("Batas pengguna tercapai.")
		return
	}

	// Input nama pengguna dan kata sandi.
	var namaPengguna, kataSandi string
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan kata sandi: ")
	fmt.Scanln(&kataSandi)

	// Validasi apakah nama pengguna sudah ada dengan memanfaatkan `cariIndeksPengguna`.
	if cariIndeksPengguna(namaPengguna) != -1 {
		fmt.Println("Nama pengguna sudah ada.")
		return
	}

	// Menambahkan pengguna baru ke array `pengguna`.
	pengguna[jumlahPengguna] = Pengguna{NamaPengguna: namaPengguna, KataSandi: kataSandi}
	jumlahPengguna++
	fmt.Println("Registrasi berhasil!")
}

// Fungsi `handlerMasuk` digunakan untuk menangani proses masuk pengguna.
// Fungsi ini memverifikasi kredensial pengguna dan mengatur indeks pengguna yang masuk.
func handlerMasuk() {
	// Memeriksa apakah sudah ada pengguna yang masuk.
	if indeksMasuk != -1 {
		fmt.Println("Sudah masuk.")
		return
	}

	// Meminta input nama pengguna dan kata sandi.
	var namaPengguna, kataSandi string
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan kata sandi: ")
	fmt.Scanln(&kataSandi)

	// Memverifikasi kredensial pengguna.
	indeks := masuk(namaPengguna, kataSandi)
	if indeks == -1 {
		fmt.Println("Nama pengguna atau kata sandi salah.")
		return
	}

	// Mengatur indeks pengguna yang masuk.
	indeksMasuk = indeks
	fmt.Println("Selamat datang,", pengguna[indeksMasuk].NamaPengguna, "!")
}

// Fungsi `keluar` digunakan untuk menangani proses keluar pengguna.
// Fungsi ini memeriksa apakah ada pengguna yang masuk, jika tidak ada maka akan menampilkan pesan kesalahan.
// Jika ada pengguna yang masuk, maka akan menampilkan pesan selamat tinggal dan mengatur indeksMasuk menjadi -1.
func keluar() {
	if indeksMasuk == -1 {
		fmt.Println("Tidak ada pengguna yang masuk.")
		return
	}

	fmt.Println("Selamat tinggal,", pengguna[indeksMasuk].NamaPengguna, "!")
	indeksMasuk = -1
}

// Fungsi `perbaruiProfil` digunakan untuk memperbarui informasi profil pengguna yang sedang masuk.
func perbaruiProfil() {
	// Memeriksa apakah ada pengguna yang sudah masuk.
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Meminta input informasi profil baru dari pengguna.
	fmt.Print("Masukkan informasi profil baru: ")
	pembaca := bufio.NewReader(os.Stdin)
	profil, err := pembaca.ReadString('\n')
	if err != nil {
		// Menangani kesalahan saat membaca input.
		fmt.Println("Terjadi kesalahan saat membaca input.")
		return
	}
	// Memperbarui profil pengguna yang sedang masuk dengan informasi baru.
	pengguna[indeksMasuk].Profil = strings.TrimSpace(profil)
	fmt.Println("Profil diperbarui!")
}

// Fungsi `tambahTeman` digunakan untuk menambahkan teman baru ke daftar teman pengguna yang sedang masuk.
func tambahTeman() {
	// Memeriksa apakah ada pengguna yang sudah masuk.
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Memeriksa apakah daftar teman sudah penuh.
	if pengguna[indeksMasuk].JumlahTeman >= MAKS_TEMAN {
		fmt.Println("Daftar teman penuh.")
		return
	}

	// Meminta input nama pengguna teman yang ingin ditambahkan.
	var namaTeman string
	fmt.Print("Masukkan nama pengguna teman yang ingin ditambahkan: ")
	fmt.Scanln(&namaTeman)

	// Mencari indeks pengguna teman berdasarkan nama pengguna.
	indeksTeman := cariIndeksPengguna(namaTeman)
	if indeksTeman == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Memeriksa apakah pengguna mencoba menambahkan dirinya sendiri sebagai teman.
	if namaTeman == pengguna[indeksMasuk].NamaPengguna {
		fmt.Println("Anda tidak dapat menambahkan diri sendiri sebagai teman.")
		return
	}

	// Memeriksa apakah teman sudah ada dalam daftar teman.
	for i := 0; i < pengguna[indeksMasuk].JumlahTeman; i++ {
		if pengguna[indeksMasuk].Teman[i] == namaTeman {
			fmt.Println("Pengguna ini sudah menjadi teman Anda.")
			return
		}
	}

	// Menambahkan teman baru ke daftar teman pengguna yang sedang masuk.
	pengguna[indeksMasuk].Teman[pengguna[indeksMasuk].JumlahTeman] = namaTeman
	pengguna[indeksMasuk].JumlahTeman++
	fmt.Println("Teman berhasil ditambahkan!")
}

// Fungsi `hapusTeman` digunakan untuk menghapus teman dari daftar teman pengguna yang sedang masuk.
func hapusTeman() {
	// Memeriksa apakah ada pengguna yang sudah masuk.
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Meminta input nama pengguna teman yang ingin dihapus.
	var namaTeman string
	fmt.Print("Masukkan nama pengguna teman yang ingin dihapus: ")
	fmt.Scanln(&namaTeman)

	// Mencari indeks teman dalam daftar teman pengguna yang sedang masuk.
	indeksTeman := -1
	i := 0
	for i < pengguna[indeksMasuk].JumlahTeman && indeksTeman == -1 {
		if pengguna[indeksMasuk].Teman[i] == namaTeman {
			indeksTeman = i
		}
		i++
	}

	// Memeriksa apakah teman ditemukan dalam daftar teman.
	if indeksTeman == -1 {
		fmt.Println("Teman tidak ditemukan dalam daftar teman Anda.")
		return
	}

	// Menggeser teman ke kiri untuk menghapus teman pada indeks yang diberikan.
	for i := indeksTeman; i < pengguna[indeksMasuk].JumlahTeman-1; i++ {
		pengguna[indeksMasuk].Teman[i] = pengguna[indeksMasuk].Teman[i+1]
	}
	// Mengosongkan teman terakhir setelah penggeseran.
	pengguna[indeksMasuk].Teman[pengguna[indeksMasuk].JumlahTeman-1] = ""
	// Mengurangi jumlah teman pengguna.
	pengguna[indeksMasuk].JumlahTeman--

	fmt.Println("Teman berhasil dihapus!")
}

// Fungsi `komentarPadaStatus` digunakan untuk menambahkan komentar pada status pengguna lain.
func komentarPadaStatus() {
	// Memeriksa apakah ada pengguna yang sudah masuk.
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Meminta input nama pengguna yang statusnya ingin dikomentari.
	var namaPengguna string
	fmt.Print("Masukkan nama pengguna yang statusnya ingin Anda komentari: ")
	fmt.Scanln(&namaPengguna)

	// Mencari indeks pengguna berdasarkan nama pengguna.
	indeksPengguna := cariIndeksPengguna(namaPengguna)
	if indeksPengguna == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Memeriksa apakah pengguna memiliki status.
	if pengguna[indeksPengguna].JumlahStatus == 0 {
		fmt.Println("Pengguna ini tidak memiliki status.")
		return
	}

	// Menampilkan semua status pengguna yang ditemukan.
	fmt.Println("--- Status Pengguna ---")
	for i := 0; i < pengguna[indeksPengguna].JumlahStatus; i++ {
		println(i+1, ".", pengguna[indeksPengguna].Status[i])
	}

	// Meminta input nomor status yang ingin dikomentari.
	var indeksStatus int
	fmt.Print("Pilih nomor status untuk dikomentari: ")
	fmt.Scanln(&indeksStatus)

	// Memeriksa validitas nomor status yang dipilih.
	if indeksStatus < 1 || indeksStatus > pengguna[indeksPengguna].JumlahStatus {
		fmt.Println("Nomor status tidak valid.")
		return
	}

	// Meminta input komentar dari pengguna.
	fmt.Print("Masukkan komentar Anda: ")
	pembaca := bufio.NewScanner(os.Stdin)
	pembaca.Scan()
	komentar := pembaca.Text()

	// Mengurangi indeksStatus untuk menyesuaikan dengan indeks berbasis nol.
	indeksStatus--
	// Menambahkan komentar ke status yang dipilih jika slot komentar masih tersedia.
	for i := 0; i < MAKS_KOMENTAR; i++ {
		if pengguna[indeksPengguna].Komentar[indeksStatus][i] == "" {
			pengguna[indeksPengguna].Komentar[indeksStatus][i] = pengguna[indeksMasuk].NamaPengguna + ": " + strings.TrimSpace(komentar)
			fmt.Println("Komentar ditambahkan!")
			return
		}
	}

	// Menampilkan pesan jika slot komentar sudah penuh.
	fmt.Println("Komentar untuk status ini penuh.")
}

// Fungsi untuk melihat dan mengurutkan daftar teman pengguna
func lihatDanUrutkanTeman() {
	// Memeriksa apakah pengguna sudah masuk
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Mendapatkan jumlah teman dari pengguna yang masuk
	jumlahTeman := pengguna[indeksMasuk].JumlahTeman
	// Memeriksa apakah pengguna memiliki teman
	if jumlahTeman == 0 {
		fmt.Println("Anda tidak memiliki teman.")
		return
	}

	// Meminta input urutan pengurutan dari pengguna
	var urutan string
	fmt.Print("Masukkan urutan pengurutan untuk teman (asc/desc): ")
	fmt.Scanln(&urutan)

	// Memeriksa validitas urutan pengurutan dan mengurutkan teman jika valid
	if urutan == "asc" || urutan == "desc" {
		teman := pengguna[indeksMasuk].Teman[:jumlahTeman]
		insertionSortTeman(teman, jumlahTeman, urutan)
		fmt.Println("Teman diurutkan!")
	} else {
		// Menampilkan pesan jika urutan pengurutan tidak valid
		fmt.Println("Urutan pengurutan tidak valid. Menampilkan teman tanpa diurutkan.")
	}

	// Menampilkan daftar teman pengguna
	fmt.Println("--- Teman Anda ---")
	for i := 0; i < jumlahTeman; i++ {
		fmt.Println(i+1, ". ", pengguna[indeksMasuk].Teman[i])
	}
}

// Fungsi untuk mengurutkan pengguna menggunakan algoritma Selection Sort
func urutkanPenggunaDenganSelectionSort() {
	// Meminta input urutan pengurutan dari pengguna
	var urutan string
	fmt.Print("Masukkan urutan pengurutan (asc/desc): ")
	fmt.Scanln(&urutan)

	// Memeriksa validitas urutan pengurutan dan mengurutkan pengguna jika valid
	if urutan == "asc" || urutan == "desc" {
		selectionSortPenggunaByNamaPengguna(urutan)
		fmt.Println("Pengguna berhasil diurutkan menggunakan Selection Sort!")
	} else {
		// Menampilkan pesan jika urutan pengurutan tidak valid
		fmt.Println("Urutan pengurutan tidak valid. Silakan coba lagi.")
	}
}

// Fungsi untuk melihat beranda
func lihatBeranda() {
	// Memeriksa apakah pengguna sudah masuk
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Menampilkan header beranda
	fmt.Println("--- Beranda ---")
	// Loop melalui semua pengguna
	for i := 0; i < jumlahPengguna; i++ {
		// Menampilkan nama pengguna
		fmt.Println(pengguna[i].NamaPengguna + ":")
		// Loop melalui semua status pengguna
		for j := 0; j < pengguna[i].JumlahStatus; j++ {
			// Menampilkan status pengguna
			fmt.Println("  - " + pengguna[i].Status[j])
			// Loop melalui semua komentar pada status
			for k := 0; k < MAKS_KOMENTAR; k++ {
				// Memeriksa dan menampilkan komentar jika ada
				if pengguna[i].Komentar[j][k] != "" {
					fmt.Println("    * " + pengguna[i].Komentar[j][k])
				}
			}
		}
	}
}

// Fungsi untuk melihat semua pengguna
func lihatPengguna() {
	// Meminta input urutan pengurutan dari pengguna
	var urutan string
	fmt.Print("Masukkan urutan pengurutan untuk pengguna (asc/desc): ")
	fmt.Scanln(&urutan)

	// Memeriksa validitas urutan pengurutan dan mengurutkan pengguna jika valid
	if urutan == "asc" || urutan == "desc" {
		selectionSortPenggunaByNamaPengguna(urutan)
		fmt.Println("Pengguna diurutkan!")
	} else {
		// Menampilkan pesan jika urutan pengurutan tidak valid
		fmt.Println("Urutan pengurutan tidak valid. Menampilkan pengguna tanpa diurutkan.")
	}

	// Menampilkan header daftar semua pengguna
	fmt.Println("--- Semua Pengguna ---")
	// Loop melalui semua pengguna
	for i := 0; i < jumlahPengguna; i++ {
		// Menampilkan nama pengguna
		fmt.Println(i+1, ".", pengguna[i].NamaPengguna)
	}
}

// --- Operasi Edit, Cari, dan Hapus ---
// Menggunakan algoritma pencarian sequential untuk mencari pengguna atau mengedit data mereka.
// Contoh: `cariPengguna` memungkinkan pengguna mencari profil berdasarkan nama pengguna.
func cariPengguna() {
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Input nama pengguna untuk pencarian.
	var namaPengguna string
	fmt.Print("Masukkan nama pengguna untuk mencari: ")
	fmt.Scanln(&namaPengguna)

	// Menggunakan `cariIndeksPengguna` untuk menemukan data.
	indeksPengguna := cariIndeksPengguna(namaPengguna)
	if indeksPengguna == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Menampilkan data pengguna yang ditemukan.
	fmt.Println("--- Profil Pengguna ---")
	fmt.Println("Nama Pengguna:", pengguna[indeksPengguna].NamaPengguna)
	fmt.Println("Profil:", pengguna[indeksPengguna].Profil)
	fmt.Println("Teman:")
	for i := 0; i < pengguna[indeksPengguna].JumlahTeman; i++ {
		fmt.Printf("  - %s\n", pengguna[indeksPengguna].Teman[i])
	}
}

func tambahDataDummy() {
	pengguna[0] = Pengguna{
		NamaPengguna: "isfa",
		KataSandi:    "kata123",
		Profil:       "Halo, saya isfa",
		Teman:        [MAKS_TEMAN]string{"nasrul", "adam", "rapi"},
		JumlahTeman:  3,
		JumlahStatus: 2,
		Status:       [MAKS_STATUS]string{"Hari ini sangat menyenangkan!", "Coding itu seru!"},
	}

	pengguna[1] = Pengguna{
		NamaPengguna: "nasrul",
		KataSandi:    "sandiaman",
		Profil:       "Hai, saya nasrul.",
		Teman:        [MAKS_TEMAN]string{"isfa", "adam"},
		JumlahTeman:  2,
		JumlahStatus: 1,
		Status:       [MAKS_STATUS]string{"Suka menjelajahi tempat baru."},
	}

	pengguna[2] = Pengguna{
		NamaPengguna: "adam",
		KataSandi:    "adam123",
		Profil:       "Fotografer dan petualang.",
		Teman:        [MAKS_TEMAN]string{"isfa", "nasrul", "soma"},
		JumlahTeman:  3,
		JumlahStatus: 3,
		Status:       [MAKS_STATUS]string{"Alam itu indah!", "Baru saja memotret pemandangan menakjubkan!", "Hidup adalah petualangan."},
	}

	pengguna[3] = Pengguna{
		NamaPengguna: "rapi",
		KataSandi:    "rapi456",
		Profil:       "Penggemar teknologi.",
		Teman:        [MAKS_TEMAN]string{"isfa"},
		JumlahTeman:  1,
		JumlahStatus: 1,
		Status:       [MAKS_STATUS]string{"Baru saja merakit PC baru!"},
	}

	pengguna[4] = Pengguna{
		NamaPengguna: "soma",
		KataSandi:    "soma123",
		Profil:       "Suka mendaki dan aktivitas luar ruangan.",
		Teman:        [MAKS_TEMAN]string{"adam"},
		JumlahTeman:  1,
		JumlahStatus: 2,
		Status:       [MAKS_STATUS]string{"Menjelajahi pegunungan.", "Terapi alam!"},
	}

	pengguna[5] = Pengguna{
		NamaPengguna: "nabila",
		KataSandi:    "nabila789",
		Profil:       "Pencinta makanan dan koki.",
		Teman:        [MAKS_TEMAN]string{"bones"},
		JumlahTeman:  1,
		JumlahStatus: 1,
		Status:       [MAKS_STATUS]string{"Memasak hidangan lezat hari ini."},
	}

	pengguna[6] = Pengguna{
		NamaPengguna: "bones",
		KataSandi:    "bones2023",
		Profil:       "Pencinta buku dan penulis.",
		Teman:        [MAKS_TEMAN]string{"nabila"},
		JumlahTeman:  1,
		JumlahStatus: 1,
		Status:       [MAKS_STATUS]string{"Sedang membaca novel yang menarik."},
	}

	pengguna[7] = Pengguna{
		NamaPengguna: "aldo",
		KataSandi:    "aldo456",
		Profil:       "Insinyur dan penggemar DIY.",
		Teman:        [MAKS_TEMAN]string{"ario"},
		JumlahTeman:  1,
		JumlahStatus: 1,
		Status:       [MAKS_STATUS]string{"Membangun sesuatu yang keren!"},
	}

	pengguna[8] = Pengguna{
		NamaPengguna: "ario",
		KataSandi:    "ario987",
		Profil:       "Pencinta kebugaran dan gym.",
		Teman:        [MAKS_TEMAN]string{"aldo"},
		JumlahTeman:  1,
		JumlahStatus: 2,
		Status:       [MAKS_STATUS]string{"Sesi gym selesai!", "Hari ini merasa kuat."},
	}

	pengguna[9] = Pengguna{
		NamaPengguna: "naufal",
		KataSandi:    "naufal123",
		Profil:       "Musisi dan seniman.",
		Teman:        [MAKS_TEMAN]string{},
		JumlahTeman:  0,
		JumlahStatus: 2,
		Status:       [MAKS_STATUS]string{"Berlatih keterampilan gitar.", "Musik adalah hidup."},
	}

	jumlahPengguna = 10
	fmt.Println("Data dummy berhasil ditambahkan!")
}

// Fungsi untuk menghapus status pengguna
// Parameter:
// - namaPengguna: string yang berisi nama pengguna yang ingin menghapus status.
// - indeksStatus: integer yang menunjukkan indeks status yang ingin dihapus.
func hapusStatus() {
	// Meminta input untuk parameter yang diperlukan
	var namaPengguna string
	var indeksStatus int
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan indeks status: ")
	fmt.Scanln(&indeksStatus)

	// Mencari indeks pengguna berdasarkan nama pengguna
	indeks := cariIndeksPengguna(namaPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Memeriksa apakah indeks status valid
	if indeksStatus < 0 || indeksStatus >= pengguna[indeks].JumlahStatus {
		fmt.Println("Indeks status tidak valid.")
		return
	}

	// Menggeser status ke kiri untuk menghapus status pada indeks yang diberikan
	for i := indeksStatus; i < pengguna[indeks].JumlahStatus-1; i++ {
		pengguna[indeks].Status[i] = pengguna[indeks].Status[i+1]
	}
	// Mengosongkan status terakhir setelah penggeseran
	pengguna[indeks].Status[pengguna[indeks].JumlahStatus-1] = ""
	// Mengurangi jumlah status pengguna
	pengguna[indeks].JumlahStatus--

	fmt.Println("Status berhasil dihapus.")
}

// Fungsi `postStatus` digunakan untuk mengunggah status baru oleh pengguna yang sedang masuk.
func postStatus() {
	// Memeriksa apakah ada pengguna yang sudah masuk.
	if indeksMasuk == -1 {
		fmt.Println("Silakan masuk terlebih dahulu.")
		return
	}

	// Memeriksa apakah daftar status sudah penuh.
	if pengguna[indeksMasuk].JumlahStatus >= MAKS_STATUS {
		fmt.Println("Daftar status penuh.")
		return
	}

	// Meminta input status baru dari pengguna.
	fmt.Print("Masukkan status Anda: ")
	pembaca := bufio.NewReader(os.Stdin)
	status, err := pembaca.ReadString('\n')
	if err != nil {
		// Menangani kesalahan saat membaca input.
		fmt.Println("Terjadi kesalahan saat membaca input.")
		return
	}
	// Menyimpan status baru ke dalam array status pengguna yang sedang masuk.
	pengguna[indeksMasuk].Status[pengguna[indeksMasuk].JumlahStatus] = strings.TrimSpace(status)
	// Menambah jumlah status pengguna.
	pengguna[indeksMasuk].JumlahStatus++
	fmt.Println("Status berhasil diunggah!")
}

// Fungsi untuk mencari status pengguna menggunakan algoritma binary search
// Catatan: Algoritma binary search mengharuskan data diurutkan terlebih dahulu.
// Parameter:
// - namaPengguna: string yang berisi nama pengguna yang ingin dicari statusnya.
// - kataKunci: string yang berisi kata kunci yang ingin dicari dalam status.
func cariStatusBinarySearch() {
	// Meminta input untuk parameter yang diperlukan
	var namaPengguna string
	var kataKunci string
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scanln(&kataKunci)

	// Mencari indeks pengguna berdasarkan nama pengguna
	indeks := cariIndeksPengguna(namaPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Mengurutkan status pengguna sebelum melakukan binary search
	insertionSortTeman(pengguna[indeks].Status[:pengguna[indeks].JumlahStatus], pengguna[indeks].JumlahStatus, "asc")

	// Melakukan binary search untuk mencari status yang mengandung kata kunci
	low, high := 0, pengguna[indeks].JumlahStatus-1
	statusDitemukan := false
	for low <= high && !statusDitemukan {
		mid := (low + high) / 2
		if strings.Contains(pengguna[indeks].Status[mid], kataKunci) {
			fmt.Printf("Status ditemukan pada indeks %d: %s\n", mid, pengguna[indeks].Status[mid])
			statusDitemukan = true
			fmt.Println("Status:", pengguna[indeks].Status[mid]) // Menampilkan status setelah ditemukan
		} else if pengguna[indeks].Status[mid] < kataKunci {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !statusDitemukan {
		fmt.Println("Tidak ada status yang mengandung kata kunci tersebut.")
	}
}

// Fungsi untuk mengedit status pengguna
// Parameter:
// - namaPengguna: string yang berisi nama pengguna yang ingin mengedit status.
// - indeksStatus: integer yang menunjukkan indeks status yang ingin diedit.
// - statusBaru: string yang berisi status baru yang akan menggantikan status lama.
func editStatus() {
	// Meminta input untuk parameter yang diperlukan
	var namaPengguna string
	var indeksStatus int
	var statusBaru string
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan indeks status: ")
	fmt.Scanln(&indeksStatus)
	fmt.Print("Masukkan status baru: ")
	fmt.Scanln(&statusBaru)

	// Mencari indeks pengguna berdasarkan nama pengguna
	indeks := cariIndeksPengguna(namaPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Memeriksa apakah indeks status valid
	if indeksStatus < 0 || indeksStatus >= pengguna[indeks].JumlahStatus {
		fmt.Println("Indeks status tidak valid.")
		return
	}

	// Memperbarui status pada indeks yang diberikan dengan status baru
	pengguna[indeks].Status[indeksStatus] = statusBaru
	fmt.Println("Status berhasil diperbarui.")
}

// Fungsi untuk mengedit komentar pada status pengguna
// Parameter:
// - namaPengguna: string yang berisi nama pengguna yang ingin mengedit komentar.
// - indeksStatus: integer yang menunjukkan indeks status yang ingin dikomentari.
// - indeksKomentar: integer yang menunjukkan indeks komentar yang ingin diedit.
// - komentarBaru: string yang berisi komentar baru yang akan menggantikan komentar lama.
func editKomentar() {
	// Meminta input untuk parameter yang diperlukan
	var namaPengguna string
	var indeksStatus int
	var indeksKomentar int
	var komentarBaru string
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan indeks status: ")
	fmt.Scanln(&indeksStatus)
	fmt.Print("Masukkan indeks komentar: ")
	fmt.Scanln(&indeksKomentar)
	fmt.Print("Masukkan komentar baru: ")
	fmt.Scanln(&komentarBaru)

	// Mencari indeks pengguna berdasarkan nama pengguna
	indeks := cariIndeksPengguna(namaPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Memeriksa apakah indeks status dan indeks komentar valid
	if indeksStatus < 0 || indeksStatus >= pengguna[indeks].JumlahStatus {
		fmt.Println("Indeks status tidak valid.")
		return
	}
	if indeksKomentar < 0 || indeksKomentar >= MAKS_KOMENTAR || pengguna[indeks].Komentar[indeksStatus][indeksKomentar] == "" {
		fmt.Println("Indeks komentar tidak valid.")
		return
	}

	// Memperbarui komentar pada indeks yang diberikan dengan komentar baru
	pengguna[indeks].Komentar[indeksStatus][indeksKomentar] = komentarBaru
	fmt.Println("Komentar berhasil diperbarui.")
}

// Fungsi untuk menghapus komentar pada status pengguna
// Parameter:
// - namaPengguna: string yang berisi nama pengguna yang ingin menghapus komentar.
// - indeksStatus: integer yang menunjukkan indeks status yang ingin dikomentari.
// - indeksKomentar: integer yang menunjukkan indeks komentar yang ingin dihapus.
func hapusKomentar() {
	// Meminta input untuk parameter yang diperlukan
	var namaPengguna string
	var indeksStatus int
	var indeksKomentar int
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&namaPengguna)
	fmt.Print("Masukkan indeks status: ")
	fmt.Scanln(&indeksStatus)
	fmt.Print("Masukkan indeks komentar: ")
	fmt.Scanln(&indeksKomentar)

	// Mencari indeks pengguna berdasarkan nama pengguna
	indeks := cariIndeksPengguna(namaPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	// Memeriksa apakah indeks status dan indeks komentar valid
	if indeksStatus < 0 || indeksStatus >= pengguna[indeks].JumlahStatus {
		fmt.Println("Indeks status tidak valid.")
		return
	}
	if indeksKomentar < 0 || indeksKomentar >= MAKS_KOMENTAR || pengguna[indeks].Komentar[indeksStatus][indeksKomentar] == "" {
		fmt.Println("Indeks komentar tidak valid.")
		return
	}

	// Menggeser komentar ke kiri untuk menghapus komentar pada indeks yang diberikan
	for i := indeksKomentar; i < MAKS_KOMENTAR-1; i++ {
		pengguna[indeks].Komentar[indeksStatus][i] = pengguna[indeks].Komentar[indeksStatus][i+1]
	}
	// Mengosongkan komentar terakhir setelah penggeseran
	pengguna[indeks].Komentar[indeksStatus][MAKS_KOMENTAR-1] = ""

	fmt.Println("Komentar berhasil dihapus.")
}

// --- Menu ---
// Mengintegrasikan semua fitur melalui menu utama yang modular.
// Setiap pilihan menu memanggil subprogram terkait.
func menu() {
	for {
		fmt.Println("\n--- Aplikasi Media Sosial ---")
		fmt.Println("1. Daftar")
		fmt.Println("2. Masuk")
		fmt.Println("3. logout")
		fmt.Println("4. Perbarui Profil")
		fmt.Println("5. Tambah Teman")
		fmt.Println("6. Hapus Teman")
		fmt.Println("7. Posting Status")
		fmt.Println("8. Edit Status")
		fmt.Println("9. Hapus Status")
		fmt.Println("10. Cari Pengguna")
		fmt.Println("11. Lihat dan Urutkan Pengguna")
		fmt.Println("12. Lihat dan Urutkan Teman")
		fmt.Println("13. Komentar pada Status")
		fmt.Println("14. Edit Komentar")
		fmt.Println("15. Hapus Komentar")
		fmt.Println("16. Lihat Beranda")
		fmt.Println("17. Lihat Beranda")
		fmt.Println("18. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			daftar() // Create
		case 2:
			handlerMasuk()
		case 3:
			keluar()
		case 4:
			perbaruiProfil() // Update
		case 5:
			tambahTeman() // Update
		case 6:
			hapusTeman() // Delete
		case 7:
			postStatus() // Create
		case 8:
			editStatus() // Update
		case 9:
			hapusStatus() // Delete
		case 10:
			cariPengguna() // Read
		case 11:
			lihatPengguna() // Read
		case 12:
			lihatDanUrutkanTeman() // Read
		case 13:
			komentarPadaStatus() // Create
		case 14:
			editKomentar() // Update
		case 15:
			hapusKomentar() // Delete
		case 16:
			lihatBeranda() // Read
		case 17:
			cariStatusBinarySearch() //Read
		case 18:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func main() {
	tambahDataDummy()
	menu()
}
