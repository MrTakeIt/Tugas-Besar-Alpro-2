package main

import "fmt"

const NMAX int = 100

type arrPositif [NMAX]string
type arrNegatif [NMAX]string
type arrKata [NMAX]string
type arrKomentar [NMAX]comment
type comment struct {
	isi           arrKata
	sentimen      string
	nilaiSentimen int
	pjgteks       int
	isialfabetis  arrKata
	jumlahKata    int
	jumlahPositif int
	jumlahNegatif int
}

func main() {
	var positif arrPositif
	var negatif arrNegatif
	var komentar arrKomentar
	var jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral int
	isiKomen(&komentar, &jumlahData)
	sentimenNegatif(&negatif)
	sentimenPositif(&positif)
	sentimen(&komentar, &jumlahData, &jumDataPositif, &jumDataNegatif, &jumDataNetral, &positif, &negatif)
	//Tampilan utama untuk menentukan pilihan menu lainnya
	var pilihanMenuUtama int
	for pilihanMenuUtama != 7 {
		tampilanUtama()
		fmt.Println("|                  MENU UTAMA                |")
		fmt.Println("==============================================")
		fmt.Println("|  1. Tambah Komentar                        |")
		fmt.Println("|  2. Ubah Komentar                          |")
		fmt.Println("|  3. Hapus Komentar                         |")
		fmt.Println("|  4. Cari Komentar                          |")
		fmt.Println("|  5. Statistik Komentar                     |")
		fmt.Println("|  6. Daftar Komentar                        |")
		fmt.Println("|  7. Keluar                                 |")
		fmt.Println("==============================================")
		fmt.Print("   Pilih menu: ")
		fmt.Scan(&pilihanMenuUtama)
		fmt.Println("==============================================")
		switch pilihanMenuUtama {
		case 1:
			//Menu menambah komentar
			menuTambahKomentar(&komentar, &jumlahData, &jumDataPositif, &jumDataNegatif, &jumDataNetral, &positif, &negatif)
		case 2:
			//Menu mengubah komentar
			menuUbahKomentar(&komentar, &jumlahData, &jumDataPositif, &jumDataNegatif, &jumDataNetral, &positif, &negatif)
		case 3:
			//Menu menghapus komentar
			menuHapusKomentar(&komentar, &jumlahData)
		case 4:
			//Menu mencari komentar
			menuCariKomentar(&komentar, &jumlahData)
		case 5:
			//Menu statistik komentar
			statistik(&komentar, &jumlahData)
		case 6:
			//Menu cetak komentar
			menuDaftarKomentar(&komentar, &jumlahData, &jumDataPositif, &jumDataNegatif, &jumDataNetral, &positif, &negatif)
		case 7:
			//Menu Keluar
			menuKeluar()
		default:
			fmt.Println("Mohon masukkan angka yang benar")
		}
	}
}

func menuTambahKomentar(komentar *arrKomentar, jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral *int, positif *arrPositif, negatif *arrNegatif) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	   Proses: Menambahkan komentar (string) ke dalam array komentar[].isi
	   F.S. Array yang terisi dengan komentar (string) dan jumlahData yang berubah  */
	var j, jumlahKomentar, jumKata, idxKomentar, idxKata int
	tampilanUtama()
	fmt.Println("|             MENU TAMBAH KOMENTAR           |")
	fmt.Println("==============================================")
	fmt.Print("   Jumlah komentar yang ingin ditambah: ")
	fmt.Scan(&jumlahKomentar)
	idxKomentar = *jumlahData
	//Proses penginputan komentar
	if idxKomentar < NMAX {
		for j < jumlahKomentar {
			fmt.Print("   Jumlah kata yang ingin ditambah: ")
			fmt.Scan(&jumKata)
			fmt.Print("   Tambahkan komentar: ")
			idxKata = 0
			komentar[idxKomentar].jumlahKata = 0
			for idxKata < jumKata {
				fmt.Scan(&komentar[idxKomentar].isi[idxKata])
				komentar[idxKomentar].jumlahKata++
				idxKata++
			}
			idxKomentar++
			j++
			*jumlahData++
			fmt.Println("==============================================")
			fmt.Println("|        Komentar berhasil ditambahkan       |")
			fmt.Println("==============================================")
		}
		sentimen(komentar, jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral, positif, negatif)
		daftarKomentar(komentar, jumlahData)
	} else {
		fmt.Println("==============================================")
		fmt.Println("|           Daftar komentar penuh!           |")
		fmt.Println("==============================================")
	}
	return
}

func tampilanUtama() {
	fmt.Println("==============================================")
	fmt.Println("|                   ADMIN                    |")
	fmt.Println("|          ANALISIS SENTIMEN KOMENTAR        |")
	fmt.Println("==============================================")
}

func menuCariKomentar(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Memilih melakukan Sequential Search atau Binary Search untuk mencari komentar berdasarkan kata kunci
	F.S. Menampilkan komentar yang telah dicari */

	//i = *idxKomentar, j = *idxKata
	tampilanUtama()
	fmt.Println("|             MENU CARI KOMENTAR             |")
	fmt.Println("==============================================")
	var kataKunci string
	var pilihanMenuCari, ketemu int
	fmt.Println("|  1. Sequential Search                      |")
	fmt.Println("|  2. Binary Search                          |")
	fmt.Println("==============================================")
	for {
		fmt.Print("   Pilih menu (1/2): ")
		fmt.Scan(&pilihanMenuCari)
		switch pilihanMenuCari {
		case 1:
			fmt.Print("   Masukkan kata kunci: ")
			fmt.Scan(&kataKunci)
			ketemu = cariKomentarSequential(komentar, *jumlahData, kataKunci)
			if ketemu != -1 {
				fmt.Println("   Komentar ditemukan pada nomor: ", ketemu+1)
				fmt.Printf("   %d. ", ketemu+1)
				for j := 0; j < komentar[ketemu].jumlahKata; j++ {
					fmt.Printf("%s ", komentar[ketemu].isi[j])
				}
				fmt.Println()
				fmt.Println("==============================================")
			} else {
				fmt.Println("|          Komentar tidak ditemukan          |")
				fmt.Println("==============================================")
			}
			return
		case 2:
			fmt.Print("   Masukkan kata kunci: ")
			fmt.Scan(&kataKunci)
			urutAlfabetAsc(komentar, jumlahData)
			ketemu = cariKomentarBinary(komentar, *jumlahData, kataKunci)
			if ketemu != -1 {
				fmt.Println("   Komentar ditemukan pada nomor: ", ketemu+1)
				fmt.Printf("   %d. ", ketemu+1)
				for j := 0; j < komentar[ketemu].jumlahKata; j++ {
					fmt.Printf("%s ", komentar[ketemu].isi[j])
				}
				fmt.Println()
				fmt.Println("==============================================")
			} else {
				fmt.Println("|          Komentar tidak ditemukan          |")
				fmt.Println("==============================================")
			}
			return
		default:
			fmt.Println("   Masukkan angka yang benar!")
		}
	}
}

func cariKomentarSequential(komentar *arrKomentar, jumlahData int, kataKunci string) int {
	/* Mengembalikan nilaia found = i jika komentar berhasil ditemukan berdasarkan kata kunci dan
	   mengembalikan nilai -1 jika komentar tidak ditemukan */
	var found int
	var ketemu = false
	found = -1
	i := 0
	for i < jumlahData && found == -1 {
		j := 0
		for j < komentar[i].jumlahKata && !ketemu {
			if komentar[i].isi[j] == kataKunci {
				ketemu = true
			}
			j++
		}
		if ketemu {
			found = i
			return found
		}
		i++
	}
	return found
}
func cariKomentarBinary(komentar *arrKomentar, jumlahData int, kataKunci string) int {
	/* Mengembalikan nilaia found = mid jika komentar berhasil ditemukan berdasarkan kata kunci dan
	   mengembalikan nilai -1 jika komentar tidak ditemukan */
	var left, mid, right int
	var found int
	found = -1
	for i := 0; i < jumlahData; i++ {
		left = 0
		right = komentar[i].jumlahKata - 1
		for left <= right && found == -1 {
			mid = (left + right) / 2
			if kataKunci > komentar[i].isialfabetis[mid] {
				left = mid + 1
			} else if kataKunci < komentar[i].isialfabetis[mid] {
				right = mid - 1
			} else {
				found = mid
			}
		}
		if found == mid {
			found = i
		}
	}
	return found
}

func urutAlfabetAsc(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Mengurutkan array berdasarkan alfabet
	F.S. Isi array terurut secara alfabetis*/
	//i = *idxKomentar, j = *idxKata
	var temp comment
	var tempKata string
	for i := 0; i < *jumlahData; i++ {
		for j := 0; j < komentar[i].jumlahKata; j++ {
			komentar[i].isialfabetis[j] = komentar[i].isi[j]
		}
	}
	for i := 1; i < *jumlahData; i++ {
		j := i
		temp = komentar[j]
		for j > 0 && temp.isialfabetis[0] < komentar[j-1].isialfabetis[0] {
			komentar[j] = komentar[j-1]
			j--
		}
		komentar[j] = temp
	}
	for i := 0; i < *jumlahData; i++ {
		for k := 1; k < komentar[i].jumlahKata; k++ {
			j := k
			tempKata = komentar[i].isialfabetis[j]
			for j > 0 && tempKata < komentar[i].isialfabetis[j-1] {
				komentar[i].isialfabetis[j] = komentar[i].isialfabetis[j-1]
				j--
			}
			komentar[i].isialfabetis[j] = tempKata
		}
	}
}

func menuUbahKomentar(komentar *arrKomentar, jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral *int, positif *arrPositif, negatif *arrNegatif) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Memilih melakukan mengubah komentar yang dipilih berdasarkan nomor urut komentar
	F.S. Menampilkan komentar yang telah diubah jika komentar ada */
	var ubahIdx, jumKata int
	tampilanUtama()
	fmt.Println("|             MENU UBAH KOMENTAR             |")
	fmt.Println("==============================================")
	daftarKomentar(komentar, jumlahData)
	if *jumlahData != 0 {
		fmt.Print("   Ubah komentar nomor: ")
		fmt.Scan(&ubahIdx)
		fmt.Print("   Jumlah kata yang ingin ditambah: ")
		fmt.Scan(&jumKata)
		fmt.Print("   Ubah komentar menjadi: ")
		komentar[ubahIdx-1].jumlahKata = jumKata
		for j := 0; j < jumKata; j++ {
			fmt.Scan(&komentar[ubahIdx-1].isi[j])
		}
		sentimen(komentar, jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral, positif, negatif)
		fmt.Println("==============================================")
		fmt.Println("|        Komentar berhasil diperbarui        |")
		fmt.Println("==============================================")
		daftarKomentar(komentar, jumlahData)
	}
	if ubahIdx > *jumlahData || *jumlahData == 0 {
		fmt.Println("==============================================")
		fmt.Println("|           Komentar tidak tersedia          |")
		fmt.Println("==============================================")
	}
	return

}
func menuKeluar() {
	fmt.Println("==============================================")
	fmt.Println("|                Terima Kasih!               |")
	fmt.Println("==============================================")
	return
}

func menuHapusKomentar(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Mengahpus komentar dengan cara mengganti array i dengan array i+1
	F.S. Menampilkan komentar setelah ada komentar yang dihapus */
	var idxHapus int
	tampilanUtama()
	fmt.Println("|             MENU HAPUS KOMENTAR            |")
	fmt.Println("==============================================")
	daftarKomentar(komentar, jumlahData)
	fmt.Print("   Hapus komen no: ")
	fmt.Scan(&idxHapus)
	if idxHapus > 0 && idxHapus <= *jumlahData {
		for i := idxHapus; i < *jumlahData; i++ {
			komentar[i-1] = komentar[i]
		}
		fmt.Println("==============================================")
		fmt.Println("|         Komentar berhasil dihapus          |")
		fmt.Println("==============================================")
	} else {
		fmt.Println("==============================================")
		fmt.Println("|          Komentar tidak ditemukan          |")
		fmt.Println("==============================================")
		return
	}
	*jumlahData = *jumlahData - 1
	daftarKomentar(komentar, jumlahData)
	return
}

func sentimen(komentar *arrKomentar, jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral *int, positif *arrPositif, negatif *arrNegatif) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Menentukan sentimen dari setiap komentar
	F.S. Komentar memiliki sentimen positif, netral, atau negatif */
	var j, k, l int

	for i := 0; i < *jumlahData; i++ {
		for j = 0; j < komentar[i].jumlahKata; j++ {
			for k = 0; k < 17; k++ {
				if komentar[i].isi[j] == positif[k] {
					komentar[i].sentimen = "(positif)"
					komentar[i].jumlahPositif += 1
					*jumDataPositif++
				} else if komentar[i].isi[j] == negatif[k] {
					komentar[i].sentimen = "(negatif)"
					komentar[i].jumlahNegatif += 1
					*jumDataNegatif++
				}
			}
			if komentar[i].isi[j] == "tidak" {
				for l = 0; l < 17; l++ {
					if komentar[i].isi[j+1] == positif[k] {
						komentar[i].sentimen = "(negatif)"
						komentar[i].jumlahNegatif += 1
						j++
						*jumDataNegatif++
					}
				}
			}
		}

		if (komentar[i].sentimen != "(positif)" && komentar[i].sentimen != "(negatif)") || komentar[i].jumlahPositif == komentar[i].jumlahNegatif {
			komentar[i].sentimen = "(netral)"
			*jumDataNetral++
		}
		j = 0
	}

}

func panjangKomen(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Menghitung panjang komentar menggunakan len()
	F.S. Mendefinisikan panjang komentar dari setiap index */
	for i := 0; i < *jumlahData; i++ {
		for j := 0; j < komentar[i].jumlahKata; j++ {
			if j != komentar[i].jumlahKata-1 {
				komentar[i].pjgteks += len(komentar[i].isi[j]) + 1
			} else {
				komentar[i].pjgteks += len(komentar[i].isi[j])
			}
		}
	}
}

func statistik(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Memilih cara mengurutkan komentar berdasarkan kriteria yang ada
	F.S. Komentar telah diurutkan berdasarkan menu pilihan */
	var pilihStatistik int
	tampilanUtama()
	fmt.Println("|               MENU STATISTIK               |")
	fmt.Println("==============================================")
	fmt.Println("   1.(Ascending) panjang teks")
	fmt.Println("   2.(Descending) panjang teks")
	fmt.Println("   3.(Ascending) sentimen")
	fmt.Println("   4.(Descending) sentimen")
	fmt.Println("==============================================")
	fmt.Print("   Pilih menu: ")
	fmt.Scan(&pilihStatistik)
	switch pilihStatistik {
	case 1:
		urutPanjangAsc(komentar, jumlahData)
		return
	case 2:
		urutPanjangDesc(komentar, jumlahData)
		return
	case 3:
		urutSentimenAsc(komentar, jumlahData)
		return
	case 4:
		urutSentimenDesc(komentar, jumlahData)
		return
	default:
		fmt.Println("   Masukkan angka yang benar!")
	}
}

func urutPanjangAsc(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Mengurutkan komentar berdasarkan panjang teks secara menaik(Ascending)
	F.S. Menampilkan komentar yang telah diurutkan berdasrkan panjang teks secara Ascending */
	//Selection Sort
	panjangKomen(komentar, jumlahData)
	var j, idx int
	var temp comment
	for i := 1; i <= *jumlahData; i++ {
		idx = i - 1
		j = i
		for j < *jumlahData {
			if komentar[idx].pjgteks > komentar[j].pjgteks {
				idx = j
			}
			j++
		}
		temp = komentar[idx]
		komentar[idx] = komentar[i-1]
		komentar[i-1] = temp
	}
	daftarKomentar(komentar, jumlahData)
}

func urutPanjangDesc(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Mengurutkan komentar berdasarkan panjang teks secara menurun(Descending)
	F.S. Menampilkan komentar yang telah diurutkan berdasrkan panjang teks secara Descending */
	//Selection Sort
	panjangKomen(komentar, jumlahData)
	var j, idx int
	var temp comment
	for i := 1; i <= *jumlahData; i++ {
		idx = i - 1
		j = i
		for j < *jumlahData {
			if komentar[idx].pjgteks < komentar[j].pjgteks {
				idx = j
			}
			j++
		}
		temp = komentar[idx]
		komentar[idx] = komentar[i-1]
		komentar[i-1] = temp
	}
	daftarKomentar(komentar, jumlahData)
}

func urutSentimenAsc(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Mengurutkan komentar berdasarkan sentimen secara menaik(Ascending)
	F.S. Menampilkan komentar yang telah diurutkan berdasrkan sentimen secara Ascending */
	//Insertion Sort
	nilaiSentimen(komentar, jumlahData)
	var j int
	var temp comment
	for i := 0; i < *jumlahData; i++ {
		j = i
		temp = komentar[j]
		for j > 0 && temp.nilaiSentimen < komentar[j-1].nilaiSentimen {
			komentar[j] = komentar[j-1]
			j--
		}
		komentar[j] = temp
	}
	daftarKomentar(komentar, jumlahData)
}
func urutSentimenDesc(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Mengurutkan komentar berdasarkan sentimen secara menaik(Descending)
	F.S. Menampilkan komentar yang telah diurutkan berdasrkan sentimen secara Descending */
	//Insertion Sort
	nilaiSentimen(komentar, jumlahData)
	var j int
	var temp comment
	for i := 0; i < *jumlahData; i++ {
		j = i
		temp = komentar[j]
		for j > 0 && temp.nilaiSentimen > komentar[j-1].nilaiSentimen {
			komentar[j] = komentar[j-1]
			j--
		}
		komentar[j] = temp
	}
	daftarKomentar(komentar, jumlahData)
}

func nilaiSentimen(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	Proses: Menentukan nilai sentimen dari setiap komentar
	F.S. Komentar memiliki nilai sentimen positif (3), netral(2), atau negatif(1) */
	for i := 0; i < *jumlahData; i++ {
		if komentar[i].sentimen == "(positif)" {
			komentar[i].nilaiSentimen = 3
		} else if komentar[i].sentimen == "(negatif)" {
			komentar[i].nilaiSentimen = 1
		} else {
			komentar[i].nilaiSentimen = 2
		}
	}

}

func menuDaftarKomentar(komentar *arrKomentar, jumlahData, jumDataPositif, jumDataNegatif, jumDataNetral *int, positif *arrPositif, negatif *arrNegatif) {
	var pilihDaftarKomentar, k int
	tampilanUtama()
	fmt.Println("|            MENU DAFTAR KOMENTAR            |")
	fmt.Println("==============================================")
	fmt.Println("   1.Tampilkan Komentar")
	fmt.Println("   2.Tampilkan Komentar Positif")
	fmt.Println("   3.Tampilkan Komentar Netral")
	fmt.Println("   4.Tampilkan Komentar Negatif")
	for {
		fmt.Print("   Pilih menu: ")
		fmt.Scan(&pilihDaftarKomentar)
		fmt.Println("==============================================")
		fmt.Println("|               DAFTAR KOMENTAR              |")
		fmt.Println("==============================================")
		switch pilihDaftarKomentar {
		case 1:
			if *jumlahData == 0 {
				fmt.Println("|          Komentar tidak tersedia!          |")
			} else {
				for i := 0; i < *jumlahData; i++ {
					fmt.Printf("   %d. ", i+1)
					for j := 0; j < komentar[i].jumlahKata; j++ {
						fmt.Printf("%s ", komentar[i].isi[j])
					}
					fmt.Printf("%s\n", komentar[i].sentimen)
				}
			}
		case 2:
			k = 1
			if *jumDataPositif == 0 {
				fmt.Println("|          Komentar tidak tersedia!          |")
			} else {
				for i := 0; i < *jumlahData; i++ {
					if komentar[i].sentimen == "(positif)" {
						fmt.Printf("   %d. ", k)
						k++
						for j := 0; j < komentar[i].jumlahKata; j++ {
							fmt.Printf("%s ", komentar[i].isi[j])
						}
						fmt.Printf("%s\n", komentar[i].sentimen)
					}
				}
			}
		case 3:
			k = 1
			if *jumDataNetral == 0 {
				fmt.Println("|          Komentar tidak tersedia!          |")
			} else {
				for i := 0; i < *jumlahData; i++ {
					if komentar[i].sentimen == "(netral)" {
						fmt.Printf("   %d. ", k)
						k++
						for j := 0; j < komentar[i].jumlahKata; j++ {
							fmt.Printf("%s ", komentar[i].isi[j])
						}
						fmt.Printf("%s\n", komentar[i].sentimen)
					}
				}
			}

		case 4:
			k = 1
			if *jumDataNegatif == 0 {
				fmt.Println("|          Komentar tidak tersedia!          |")
			} else {
				for i := 0; i < *jumlahData; i++ {
					if komentar[i].sentimen == "(negatif)" {
						fmt.Printf("   %d. ", k)
						k++
						for j := 0; j < komentar[i].jumlahKata; j++ {
							fmt.Printf("%s ", komentar[i].isi[j])
						}
						fmt.Printf("%s\n", komentar[i].sentimen)
					}
				}
			}
		default:
			fmt.Println("   Masukkan angka yang benar!")

		}
		fmt.Println("==============================================")
		return
	}
}

func daftarKomentar(komentar *arrKomentar, jumlahData *int) {
	/* I.S. Terdefinisi array dinamis arrKomentar, i integer, dan jumlahData integer
	   F.S. Menampilkan komentar (string) serta nilai sentimen jika ada */
	fmt.Println("==============================================")
	fmt.Println("|               DAFTAR KOMENTAR              |")
	fmt.Println("==============================================")
	if *jumlahData == 0 {
		fmt.Println("|          Komentar tidak tersedia!          |")
	} else {
		for i := 0; i < *jumlahData; i++ {
			fmt.Printf("   %d. ", i+1)
			for j := 0; j < komentar[i].jumlahKata; j++ {
				fmt.Printf("%s ", komentar[i].isi[j])
			}
			fmt.Printf("%s\n", komentar[i].sentimen)
		}
	}
	fmt.Println("==============================================")
	return
}

func isiKomen(komentar *arrKomentar, jumlahData *int) {
	*jumlahData = 10
	komentar[0].isi[0] = "keren"
	komentar[0].isi[1] = "banget"
	komentar[0].jumlahKata = 2

	komentar[1].isi[0] = "tidak"
	komentar[1].isi[1] = "menarik"
	komentar[1].jumlahKata = 2

	komentar[2].isi[0] = "sangat"
	komentar[2].isi[1] = "membantu"
	komentar[2].jumlahKata = 2

	komentar[3].isi[0] = "sangat"
	komentar[3].isi[1] = "buruk"
	komentar[3].jumlahKata = 2

	komentar[4].isi[0] = "bagus"
	komentar[4].isi[1] = "dan"
	komentar[4].isi[2] = "rapi"
	komentar[4].jumlahKata = 3

	komentar[5].isi[0] = "terlalu"
	komentar[5].isi[1] = "jelek"
	komentar[5].jumlahKata = 2

	komentar[6].isi[0] = "penjelasan"
	komentar[6].isi[1] = "oke"
	komentar[6].jumlahKata = 2

	komentar[7].isi[0] = "kurang"
	komentar[7].isi[1] = "banyak"
	komentar[7].jumlahKata = 2

	komentar[8].isi[0] = "mantap"
	komentar[8].isi[1] = "hasilnya"
	komentar[8].jumlahKata = 2

	komentar[9].isi[0] = "hmmm"
	komentar[9].isi[1] = "okee"
	komentar[9].jumlahKata = 2
}

func sentimenPositif(positif *arrPositif) {
	positif[0] = "keren"
	positif[1] = "bagus"
	positif[2] = "mantap"
	positif[3] = "hebat"
	positif[4] = "baik"
	positif[5] = "top"
	positif[6] = "sempurna"
	positif[7] = "jempolan"
	positif[8] = "brilian"
	positif[9] = "cerdas"
	positif[10] = "terbaik"
	positif[11] = "menawan"
	positif[12] = "ramah"
	positif[13] = "asik"
	positif[14] = "membantu"
	positif[15] = "menarik"
	positif[16] = "berkelas"
}

func sentimenNegatif(negatif *arrNegatif) {
	negatif[0] = "jelek"
	negatif[1] = "buruk"
	negatif[2] = "lemah"
	negatif[3] = "bodoh"
	negatif[4] = "payah"
	negatif[5] = "rusak"
	negatif[6] = "menyebalkan"
	negatif[7] = "jahat"
	negatif[8] = "malas"
	negatif[9] = "busuk"
	negatif[10] = "gagal"
	negatif[11] = "bohong"
	negatif[12] = "palsu"
	negatif[13] = "sombong"
	negatif[14] = "menipu"
	negatif[15] = "seram"
	negatif[16] = "nyesal"
}
