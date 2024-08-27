package main

import "fmt"

const NMAXCALEG = 100
const MAXUSERS = 1000

type Caleg struct {
	No_Urut int
	Nama    string
	Partai  string
	Suara   int
}

type ListCaleg struct {
	Daftar [NMAXCALEG]Caleg
	NCaleg int
}
type waktu struct {
	jam, menit int
}

var listCaleg ListCaleg
var pemilihan [MAXUSERS]int

func main() {
	var waktu_mulai, waktu_selesai waktu
	var flag_waktu bool

	waktu_mulai.jam = 6
	waktu_selesai.jam = 14

	flag_waktu = waktu_mulai.jam < waktu_selesai.jam

	listCaleg.NCaleg = 4

	listCaleg.Daftar[0].No_Urut = 1
	listCaleg.Daftar[0].Nama = "Supriadi"
	listCaleg.Daftar[0].Partai = "Kejora"

	listCaleg.Daftar[1].No_Urut = 2
	listCaleg.Daftar[1].Nama = "Agus"
	listCaleg.Daftar[1].Partai = "Bintang Bercahaya"

	listCaleg.Daftar[2].No_Urut = 3
	listCaleg.Daftar[2].Nama = "Irwansyah"
	listCaleg.Daftar[2].Partai = "Jeruji Besi"

	listCaleg.Daftar[3].No_Urut = 4
	listCaleg.Daftar[3].Nama = "Siti"
	listCaleg.Daftar[3].Partai = "Sejahtera"

	listCaleg.Daftar[4].No_Urut = 5
	listCaleg.Daftar[4].Nama = "Kaesang"
	listCaleg.Daftar[4].Partai = "Pertanian"

	var pilihan int
	var threshold int = 0
	var JumUser int

	for {
		fmt.Println("\nPilih Peran:")
		fmt.Println("1. Admin")
		fmt.Println("2. User")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			adminMenu(&listCaleg, &threshold, flag_waktu)
		case 2:
			if JumUser < MAXUSERS {
				userMenu(&listCaleg, &JumUser, flag_waktu)
			} else {
				fmt.Println("Batas maksimum user tercapai.")
			}
		case 3:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func adminMenu(C *ListCaleg, threshold *int, waktu_Ontime bool) {
	var pilihan, pilih_sort, mencari, idx int
	var X string
	idx = -1
	for {
		fmt.Println("\nMenu Admin:")
		if waktu_Ontime {
			menu_onTime()
		} else {
			menu_NotOnTime()
		}
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		if !waktu_Ontime {
			fmt.Println("Waktu pemilihan sudah habis. Anda hanya dapat menampilkan caleg atau keluar.")
			for pilihan > 1 {
				fmt.Print("Masukan Pilihan: ")
				fmt.Scan(&pilihan)
			}
		}

		switch pilihan {
		case 1:
			tampilkanCaleg(C)
		case 2:
			editCaleg(C)
		case 3:
			hapusCaleg(C)
		case 4:
			tambahCaleg(C)
		case 5:
			fmt.Println("1. Mencari Berdasarkan Nama")
			fmt.Println("2. Mencari Berdasarkan Partai")
			fmt.Println("3. Mencasi Berdasarkan Nomor Urut")
			fmt.Println("Masukan pilihan :")
			fmt.Scan(&mencari)
			if mencari == 1 {
				fmt.Print("Masukan Nama Caleg: ")
				fmt.Scan(&X)
				idx = mencari_namaCaleg(*C, X)
				if idx != -1 {
					fmt.Printf("NO: %d, Nama: %s, Partai: %s, Suara: %d\n", C.Daftar[idx].No_Urut, C.Daftar[idx].Nama, C.Daftar[idx].Partai, C.Daftar[idx].Suara)
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			} else if mencari == 2 {
				fmt.Print("Masukan Nama Partai: ")
				fmt.Scan(&X)
				idx = mencari_namaPartai(*C, X)
				if idx != -1 {
					fmt.Printf("NO: %d, Nama: %s, Partai: %s, Suara: %d\n", C.Daftar[idx].No_Urut, C.Daftar[idx].Nama, C.Daftar[idx].Partai, C.Daftar[idx].Suara)
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			} else if mencari == 3 {
				var y int
				fmt.Println("Masukan jumlah Nomor Urutan: ")
				fmt.Scan(&y)
				idx = mencari_nomorUrut(*C, y)
				if idx != -1 {
					fmt.Printf("NO: %d, Nama: %s, Partai: %s, Suara: %d\n", C.Daftar[idx].No_Urut, C.Daftar[idx].Nama, C.Daftar[idx].Partai, C.Daftar[idx].Suara)
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			} else {
				fmt.Print("Data tidak ditemukan")
			}
		case 6:
			fmt.Println("1. Asending")
			fmt.Println("2. Desending")
			fmt.Print("Masukan Pilihan: ")
			fmt.Scan(&pilih_sort)
			if pilih_sort == 1 {
				SortingSuara_ASC(C)
				Print_Sort(*C, C.NCaleg)
			} else if pilih_sort == 2 {
				SortingSuara_DESC(C)
				Print_Sort(*C, C.NCaleg)
			}

		case 7:
			tampilkanThreshold(C)
		case 0:
			return
		default:
			fmt.Println("Waktu pemilihan sudah habis")
		}
	}
}

func menu_onTime() {
	fmt.Println("1. Tampilkan Caleg")
	fmt.Println("2. Edit Caleg")
	fmt.Println("3. Hapus Caleg")
	fmt.Println("4. Tambah Caleg")
	fmt.Println("5. Mencari berdasarkan Nama, Partai, dan Nomor Urut")
	fmt.Println("6. Sorting")
	fmt.Println("7. Tampilkan Threshold")
	fmt.Println("0. Kembali")
}

func menu_NotOnTime() {
	fmt.Println("1. Tampilkan Caleg")
	fmt.Println("0. Kembali")
}
func userMenu(C *ListCaleg, JumUser *int, waktu_Ontime bool) {
	var pilihan int

	fmt.Println("\nDaftar Caleg:")
	for i := 0; i < C.NCaleg; i++ {
		fmt.Printf("NO: %d, %d. Nama: %s, Partai: %s\n", i+1, C.Daftar[i].No_Urut, C.Daftar[i].Nama, C.Daftar[i].Partai)
	}

	if waktu_Ontime {
		fmt.Print("Pilih Caleg  - ", C.NCaleg, ": ")
		fmt.Scan(&pilihan)

		if pilihan >= 1 && pilihan <= C.NCaleg {
			*JumUser++
			C.Daftar[pilihan-1].Suara++
			fmt.Println("Pilihan Anda telah disimpan.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanCaleg(C *ListCaleg) {
	fmt.Println("Daftar Caleg:")
	for i := 0; i < C.NCaleg; i++ {
		fmt.Printf("NO: %d, Nama: %s, Partai: %s, Suara: %d\n", C.Daftar[i].No_Urut, C.Daftar[i].Nama, C.Daftar[i].Partai, C.Daftar[i].Suara)
	}
}

func editCaleg(C *ListCaleg) {
	var namaCaleg string
	var pilihan int

	fmt.Print("Masukkan nama caleg yang akan diedit: ")
	fmt.Scan(&namaCaleg)

	index := mencari_namaCaleg(*C, namaCaleg)

	if index == -1 {
		fmt.Println("Caleg tidak ditemukan.")
		return
	}

	fmt.Printf("Caleg ditemukan: No: %d, Nama: %s, Partai: %s, Suara: %d\n",
		C.Daftar[index].No_Urut, C.Daftar[index].Nama, C.Daftar[index].Partai, C.Daftar[index].Suara)

	fmt.Println("Pilih data yang ingin diubah:")
	fmt.Println("1. Nama")
	fmt.Println("2. Partai")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var namaBaru string
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)

		if mencari_namaCaleg(*C, namaBaru) != -1 {
			fmt.Println("Nama caleg baru sudah ada dalam daftar.")
			return
		}

		C.Daftar[index].Nama = namaBaru
		fmt.Println("Nama caleg berhasil diubah.")

	case 2:
		var partaiBaru string
		fmt.Print("Masukkan partai baru: ")
		fmt.Scan(&partaiBaru)
		C.Daftar[index].Partai = partaiBaru
		fmt.Println("Partai caleg berhasil diubah.")

	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Printf("Data caleg setelah diedit: NO: %d, Nama: %s, Partai: %s, Suara: %d\n",
		C.Daftar[index].No_Urut, C.Daftar[index].Nama, C.Daftar[index].Partai, C.Daftar[index].Suara)
}

func tambahCaleg(C *ListCaleg) {
	if C.NCaleg >= NMAXCALEG {
		fmt.Println("Maaf, jumlah maksimum caleg sudah tercapai.")
		return
	}

	var newCaleg Caleg

	fmt.Print("Masukan No Urutan Caleg baru: ")
	fmt.Scan(&newCaleg.No_Urut)

	fmt.Print("Masukkan nama caleg baru: ")
	fmt.Scan(&newCaleg.Nama)

	fmt.Print("Masukkan partai caleg baru: ")
	fmt.Scan(&newCaleg.Partai)

	newCaleg.Suara = 0

	for i := 0; i < C.NCaleg; i++ {
		if C.Daftar[i].Nama == newCaleg.Nama {
			fmt.Println("Caleg dengan nama tersebut sudah ada.")
			return
		}
	}

	C.Daftar[C.NCaleg] = newCaleg
	C.NCaleg++

	fmt.Printf("NO %d Caleg %s dari partai %s berhasil ditambahkan.\n", newCaleg.No_Urut, newCaleg.Nama, newCaleg.Partai)
}

func hapusCaleg(C *ListCaleg) {
	var nama string

	fmt.Println("Masukkan nama caleg yang akan dihapus: ")
	fmt.Scan(&nama)
	for i := 0; i < C.NCaleg; i++ {
		if C.Daftar[i].Nama == nama {
			C.Daftar[i] = C.Daftar[C.NCaleg-1]
			C.NCaleg--
			fmt.Println("Caleg berhasil dihapus.")
			return
		}
	}
	fmt.Println("Caleg tidak ditemukan.")
}

func mencari_namaCaleg(C ListCaleg, X string) int {
	fmt.Println("Hasil Pencarian Berdasarkan Nama: ")
	var found int = -1
	var j int = 0
	for j < C.NCaleg && found == -1 {
		if C.Daftar[j].Nama == X {
			found = j
		}
		j = j + 1
	}
	return found

}

func mencari_namaPartai(C ListCaleg, X string) int {
	var found int = -1
	var j int = 0
	fmt.Println("Hasil Pencarian Berdasarkan partai: ")
	for j < C.NCaleg && found == -1 {
		if C.Daftar[j].Partai == X {
			found = j
		}
		j = j + 1
	}
	return found
}

func mencari_nomorUrut(C ListCaleg, x int) int {
	var bawah, atas, tengah int
	var ketemu int = -1

	bawah = 0
	atas = C.NCaleg - 1
	for bawah <= atas && ketemu == -1 {
		tengah = (bawah + atas) / 2
		if x > C.Daftar[tengah].No_Urut {
			bawah = tengah + 1
		} else if x < C.Daftar[tengah].No_Urut {
			atas = tengah - 1
		} else {
			ketemu = tengah
		}

	}

	return ketemu
}
func SortingSuara_DESC(C *ListCaleg) {
	var n int = C.NCaleg
	fmt.Print("Hasil Sorting Berdasarkan Descending: ")
	var pass, j int
	var temp Caleg
	pass = 1
	for pass <= n-1 {
		j = pass
		temp = C.Daftar[j]
		for j > 0 && temp.Suara > C.Daftar[j-1].Suara {
			C.Daftar[j] = C.Daftar[j-1]
			j = j - 1
		}
		C.Daftar[j].Suara = temp.Suara
		pass = pass + 1
	}
	fmt.Println()
}
func SortingSuara_ASC(C *ListCaleg) {
	fmt.Print("Hasil Sorting Berdasarkan Asending: ")
	var i, j, idx_min int
	var t Caleg
	i = 1
	for i <= C.NCaleg-1 {
		idx_min = i - 1
		j = i
		for j < C.NCaleg {
			if C.Daftar[idx_min].Suara > C.Daftar[j].Suara {
				idx_min = j
			}
			j = j + 1
		}
		t = C.Daftar[idx_min]
		C.Daftar[idx_min] = C.Daftar[i-1]
		C.Daftar[i-1] = t
		i = i + 1
	}
	fmt.Println()
}
func Print_Sort(C ListCaleg, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("No: %d, Nama: %s, Partai: %s, Suara: %d\n", C.Daftar[i].No_Urut, C.Daftar[i].Nama, C.Daftar[i].Partai, C.Daftar[i].Suara)
	}
}
func hitungThreshold(C ListCaleg) float64 {
	var totalSuara int
	for i := 0; i < C.NCaleg; i++ {
		totalSuara += C.Daftar[i].Suara
	}
	return float64(totalSuara) * 0.51
}

func tampilkanThreshold(C *ListCaleg) {
	threshold := hitungThreshold(*C)
	fmt.Printf("Threshold perolehan suara: %.2f\n", threshold)
	fmt.Println("Caleg yang memenuhi syarat kemenangan:")

	for i := 0; i < C.NCaleg; i++ {
		if float64(C.Daftar[i].Suara) > threshold {
			fmt.Printf("NO: %d, Nama: %s, Partai: %s, Suara: %d\n", C.Daftar[i].No_Urut, C.Daftar[i].Nama, C.Daftar[i].Partai, C.Daftar[i].Suara)
		}
	}
}
