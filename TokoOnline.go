package main

import "fmt"

// mengatur banyak array, pj = penjual, pb = pembeli, brg = barang.
const NMAXPJ int = 50
const NMAXPB int = 50
const NMAXBRG int = 500

// pengaturan struct masing-masing variable yang digunakan

type akunpj struct {
	Usernamepj string
	Passwordpj string
	Approved   bool
}

type akunpb struct {
	Usernamepb string
	Passwordpb string
	Approved   bool
}

type barang struct {
	Nama    string
	Harga   float64
	Stok    int
	Penjual string
}

type transaksi struct {
	Pembeli  string
	Barang   string
	Quantity int
}

type penjual struct {
	arrPenjual [NMAXPJ]akunpj
	totPj      int
}

type pembeli struct {
	arrPembeli [NMAXPB]akunpb
	totPb      int
}

type toko struct {
	arrBarang    [NMAXBRG]barang
	totBarang    int
	arrTransaksi []transaksi
}

// menentukan apa yang ingin dilakukan user
func main() {
	var mauapa, regisapa, loginapa, kodeotp int
	var pj penjual
	var pb pembeli
	var tk toko
	var logadm, passadm string

	for {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("Selamat datang di Tel-u Shop")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
		fmt.Println("--------------------------------------------------------------")
		fmt.Scan(&mauapa)

		if mauapa == 1 {
			for {
				fmt.Println("--------------------------------------------------------------")
				fmt.Println("ingin registrasi sebagai?")
				fmt.Println("1. Penjual")
				fmt.Println("2. Pembeli")
				fmt.Println("3. Kembali ke Menu Utama")
				fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
				fmt.Println("--------------------------------------------------------------")
				fmt.Scan(&regisapa)

				if regisapa == 1 {
					Registrasipj(&pj)
				} else if regisapa == 2 {
					Registrasipb(&pb)
				} else if regisapa == 3 {
					break
				}
			}
		} else if mauapa == 2 {
			for {
				fmt.Println("--------------------------------------------------------------")
				fmt.Println("ingin login sebagai?")
				fmt.Println("1. Penjual")
				fmt.Println("2. Pembeli")
				fmt.Println("3. Admin")
				fmt.Println("4. Kembali ke Menu Utama")
				fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
				fmt.Println("--------------------------------------------------------------")
				fmt.Scan(&loginapa)

				if loginapa == 1 {
					LoginPj(&pj, &tk)
				} else if loginapa == 2 {
					LoginPb(&pb, &tk)
				} else if loginapa == 3 {
					fmt.Println("Masukkan Username Admin:")
					fmt.Scan(&logadm)
					fmt.Println("Masukkan Password Admin:")
					fmt.Scan(&passadm)
					if logadm == "admin" && passadm == "admin" {
						fmt.Println("Username dan password admin benar!")
						fmt.Println("Masukkan Kode OTP:")
						fmt.Scan(&kodeotp)
						if kodeotp == 123 {
							fmt.Println("Login admin berhasil!")
							kerjaanAdm(&pj, &pb)
						} else {
							fmt.Println("Kode OTP salah!")
						}
					} else {
						fmt.Println("Username atau password admin salah!")
					}
				} else if loginapa == 4 {
					break
				}
			}
		} else if mauapa == 3 {
			fmt.Println("Terima kasih telah menggunakan Tel-u Shop!")
			return
		} else {
			fmt.Println("Pilihan tidak valid, silahkan coba lagi.")
		}
	}
}

// Tempat registrasi dan pengecekan username penjual dan pembeli
func Registrasipj(A *penjual) {
	var saha, sahapass string
	fmt.Println("Masukkan Username untuk nama toko yang diinginkan:")
	fmt.Scan(&saha)
	if searchkembar(*A, saha) != -1 {
		fmt.Println("Username sudah ada yang memiliki, mohon gunakan yang lain.")
	} else {
		fmt.Println("Username tersedia, masukkan password yang diinginkan:")
		fmt.Scan(&sahapass)
		A.arrPenjual[A.totPj].Usernamepj = saha
		A.arrPenjual[A.totPj].Passwordpj = sahapass
		A.arrPenjual[A.totPj].Approved = false
		A.totPj++
		fmt.Println("Registrasi penjual berhasil! Menunggu persetujuan admin.")
	}
}

// mendeteksi array yang kembar (jika ada)
func searchkembar(A penjual, B string) int {
	for i := 0; i < A.totPj; i++ {
		if B == A.arrPenjual[i].Usernamepj {
			return i
		}
	}
	return -1
}

func Registrasipb(A *pembeli) {
	var saha, sahapass string
	fmt.Println("Masukkan Username untuk nama yang diinginkan:")
	fmt.Scan(&saha)
	if searchkembarj(*A, saha) != -1 {
		fmt.Println("Username sudah ada yang memiliki, mohon gunakan yang lain.")
	} else {
		fmt.Println("Username tersedia, masukkan password yang diinginkan:")
		fmt.Scan(&sahapass)
		A.arrPembeli[A.totPb].Usernamepb = saha
		A.arrPembeli[A.totPb].Passwordpb = sahapass
		A.arrPembeli[A.totPb].Approved = false
		A.totPb++
		fmt.Println("Registrasi pembeli berhasil! Menunggu persetujuan admin.")
	}
}

func searchkembarj(A pembeli, B string) int {
	for i := 0; i < A.totPb; i++ {
		if B == A.arrPembeli[i].Usernamepb {
			return i
		}
	}
	return -1
}

// login pembeli dan penjual.
func LoginPj(A *penjual, T *toko) {
	var username, password string
	fmt.Println("Masukkan Username:")
	fmt.Scan(&username)
	fmt.Println("Masukkan Password:")
	fmt.Scan(&password)
	for i := 0; i < A.totPj; i++ {
		if A.arrPenjual[i].Usernamepj == username && A.arrPenjual[i].Passwordpj == password {
			if A.arrPenjual[i].Approved {
				fmt.Println("Login berhasil!")
				penjualMenu(T, username)
				return
			} else {
				fmt.Println("Akun Anda belum disetujui oleh admin.")
				return
			}
		}
	}
	fmt.Println("Username atau Password salah!")
}

func LoginPb(A *pembeli, T *toko) {
	var username, password string
	fmt.Println("Masukkan Username:")
	fmt.Scan(&username)
	fmt.Println("Masukkan Password:")
	fmt.Scan(&password)
	for i := 0; i < A.totPb; i++ {
		if A.arrPembeli[i].Usernamepb == username && A.arrPembeli[i].Passwordpb == password {
			if A.arrPembeli[i].Approved {
				fmt.Println("Login berhasil!")
				pembeliMenu(T, username)
				return
			} else {
				fmt.Println("Akun Anda belum disetujui oleh admin.")
				return
			}
		}
	}
	fmt.Println("Username atau Password salah!")
}

// penjual Menu, untuk menentukan apa yg ingin dilakukan penjual.
func penjualMenu(T *toko, username string) {
	var pilihan int
	for pilihan != 5 {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Lihat Barang")
		fmt.Println("5. Keluar")
		fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
		fmt.Println("--------------------------------------------------------------")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tambahBarang(T, username)
		} else if pilihan == 2 {
			ubahBarang(T, username)
		} else if pilihan == 3 {
			hapusBarang(T, username)
		} else if pilihan == 4 {
			lihatBarangPenjual(T, username)
		} else if pilihan == 5 {
			fmt.Println("Keluar dari menu penjual")
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// untuk menambah barang
func tambahBarang(T *toko, username string) {
	var nama string
	var harga float64
	var stok int
	fmt.Println("Masukkan Nama Barang:")
	fmt.Scan(&nama)
	fmt.Println("Masukkan Harga Barang:")
	fmt.Scan(&harga)
	fmt.Println("Masukkan Stok Barang:")
	fmt.Scan(&stok)
	T.arrBarang[T.totBarang] = barang{Nama: nama, Harga: harga, Stok: stok, Penjual: username}
	T.totBarang++
	fmt.Println("Barang berhasil ditambahkan!")
}

// untuk mengubah barang (yg sudah ada)
func ubahBarang(T *toko, username string) {
	var nama string
	var harga float64
	var stok int
	fmt.Println("Masukkan Nama Barang yang akan diubah:")
	fmt.Scan(&nama)
	for i := 0; i < T.totBarang; i++ {
		if T.arrBarang[i].Nama == nama && T.arrBarang[i].Penjual == username {
			fmt.Println("Masukkan Harga Baru:")
			fmt.Scan(&harga)
			fmt.Println("Masukkan Stok Baru:")
			fmt.Scan(&stok)
			T.arrBarang[i].Harga = harga
			T.arrBarang[i].Stok = stok
			fmt.Println("Barang berhasil diubah!")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan atau Anda tidak memiliki hak untuk mengubah barang ini!")
}

// untuk menghapus barang
func hapusBarang(T *toko, username string) {
	var nama string
	fmt.Println("Masukkan Nama Barang yang akan dihapus:")
	fmt.Scan(&nama)
	for i := 0; i < T.totBarang; i++ {
		if T.arrBarang[i].Nama == nama && T.arrBarang[i].Penjual == username {
			T.arrBarang[i] = T.arrBarang[T.totBarang-1]
			T.totBarang--
			fmt.Println("Barang berhasil dihapus!")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan atau Anda tidak memiliki hak untuk menghapus barang ini!")
}

// Untuk melihat barang milik penjual yang sedang login
func lihatBarangPenjual(T *toko, username string) {
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Daftar Barang:")
	for i := 0; i < T.totBarang; i++ {
		if T.arrBarang[i].Penjual == username {
			fmt.Printf("Nama: %s, Harga: %.2f, Stok: %d\n", T.arrBarang[i].Nama, T.arrBarang[i].Harga, T.arrBarang[i].Stok)
		}
	}
	fmt.Println("--------------------------------------------------------------")
}

// pembeli Menu, untuk menentukan apa yg ingin dilakukan pembeli.
func pembeliMenu(T *toko, username string) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("1. Beli Barang")
		fmt.Println("2. Lihat Barang")
		fmt.Println("3. Keluar")
		fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
		fmt.Println("--------------------------------------------------------------")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			beliBarang(T, username)
		} else if pilihan == 2 {
			lihatBarang(T)
		} else if pilihan == 3 {
			fmt.Println("Keluar dari menu pembeli")
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// pembeli ingin membeli barang
func beliBarang(T *toko, username string) {
	var namaToko, namaBarang string
	var quantity int
	fmt.Println("Masukkan Nama Toko:")
	fmt.Scan(&namaToko)
	fmt.Println("Masukkan Nama Barang yang ingin dibeli:")
	fmt.Scan(&namaBarang)
	fmt.Println("Masukkan Jumlah yang ingin dibeli:")
	fmt.Scan(&quantity)
	for i := 0; i < T.totBarang; i++ {
		if T.arrBarang[i].Penjual == namaToko && T.arrBarang[i].Nama == namaBarang {
			if T.arrBarang[i].Stok >= quantity {
				T.arrBarang[i].Stok -= quantity
				transaksiBaru := transaksi{Pembeli: username, Barang: namaBarang, Quantity: quantity}
				T.arrTransaksi = append(T.arrTransaksi, transaksiBaru)
				fmt.Println("Pembelian berhasil!")
				cetakStruk(transaksiBaru, T.arrBarang[i], username)
			} else {
				fmt.Println("Stok tidak mencukupi!")
			}
			return
		}
	}
	fmt.Println("Barang tidak ditemukan!")
}

// pencetak struk
func cetakStruk(t transaksi, b barang, pembeli string) {
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Struk Pembelian")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Nama Barang: %s\n", b.Nama)
	fmt.Printf("Jumlah Barang: %d\n", t.Quantity)
	fmt.Printf("Total Harga: %.2f\n", float64(t.Quantity)*b.Harga)
	fmt.Printf("Dari Toko: %s\n", b.Penjual)
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Terima kasih sudah melakukan pembelian %s, semoga melakukan pembelian kembali.\n", pembeli)
	fmt.Println("--------------------------------------------------------------")
}

// untuk melihat barang yg tersedia di toko
func lihatBarang(T *toko) {
	var pilihan int
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("1. Urutkan Harga Ascending (Rendah ke Tinggi)")
	fmt.Println("2. Urutkan Harga Descending (Tinggi ke Rendah)")
	fmt.Println("3. Lihat Tanpa Urutkan")
	fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
	fmt.Println("--------------------------------------------------------------")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		sortBarang(T, true)
	} else if pilihan == 2 {
		sortBarang(T, false)
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Daftar Barang:")
	for i := 0; i < T.totBarang; i++ {
		fmt.Printf("Nama: %s, Harga: %.2f, Stok: %d, Toko: %s\n", T.arrBarang[i].Nama, T.arrBarang[i].Harga, T.arrBarang[i].Stok, T.arrBarang[i].Penjual)
	}
	fmt.Println("--------------------------------------------------------------")
}

func sortBarang(T *toko, ascending bool) {
	for i := 0; i < T.totBarang-1; i++ {
		idx := i
		for j := i + 1; j < T.totBarang; j++ {
			if ascending {
				if T.arrBarang[j].Harga < T.arrBarang[idx].Harga {
					idx = j
				}
			} else {
				if T.arrBarang[j].Harga > T.arrBarang[idx].Harga {
					idx = j
				}
			}
		}
		if idx != i {
			T.arrBarang[i], T.arrBarang[idx] = T.arrBarang[idx], T.arrBarang[i]
		}
	}
}

// Kerjaan Admin
func kerjaanAdm(pj *penjual, pb *pembeli) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("1. Setujui/Tolak Registrasi Penjual")
		fmt.Println("2. Setujui/Tolak Registrasi Pembeli")
		fmt.Println("3. Keluar")
		fmt.Println("Silahkan dipilih angka dari perintah yang ingin dilakukan")
		fmt.Println("--------------------------------------------------------------")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("--------------------------------------------------------------")
			fmt.Println("Daftar Penjual yang belum disetujui:")
			for i := 0; i < pj.totPj; i++ {
				if !pj.arrPenjual[i].Approved {
					fmt.Printf("Username: %s\n", pj.arrPenjual[i].Usernamepj)
				}
			}
			setujuiTolakRegistrasiPenjual(pj)
		} else if pilihan == 2 {
			fmt.Println("--------------------------------------------------------------")
			fmt.Println("Daftar Pembeli yang belum disetujui:")
			for i := 0; i < pb.totPb; i++ {
				if !pb.arrPembeli[i].Approved {
					fmt.Printf("Username: %s\n", pb.arrPembeli[i].Usernamepb)
				}
			}
			setujuiTolakRegistrasiPembeli(pb)
		} else if pilihan == 3 {
			fmt.Println("Keluar dari menu admin")
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// untuk menolak / menerima akun pembeli dan penjual yang registrasi.
func setujuiTolakRegistrasiPenjual(pj *penjual) {
	var username string
	var approve int
	fmt.Println("Masukkan Username Penjual yang ingin disetujui/ditolak:")
	fmt.Scan(&username)
	for i := 0; i < pj.totPj; i++ {
		if pj.arrPenjual[i].Usernamepj == username {
			fmt.Println("1. Setujui")
			fmt.Println("2. Tolak")
			fmt.Scan(&approve)
			if approve == 1 {
				pj.arrPenjual[i].Approved = true
				fmt.Println("Penjual disetujui!")
			} else if approve == 2 {
				pj.arrPenjual[i] = pj.arrPenjual[pj.totPj-1]
				pj.totPj--
				fmt.Println("Penjual ditolak!")
			} else {
				fmt.Println("Pilihan tidak valid")
			}
			return
		}
	}
	fmt.Println("Penjual tidak ditemukan!")
}

func setujuiTolakRegistrasiPembeli(pb *pembeli) {
	var username string
	var approve int
	fmt.Println("Masukkan Username Pembeli yang ingin disetujui/ditolak:")
	fmt.Scan(&username)
	for i := 0; i < pb.totPb; i++ {
		if pb.arrPembeli[i].Usernamepb == username {
			fmt.Println("1. Setujui")
			fmt.Println("2. Tolak")
			fmt.Scan(&approve)
			if approve == 1 {
				pb.arrPembeli[i].Approved = true
				fmt.Println("Pembeli disetujui!")
			} else if approve == 2 {
				pb.arrPembeli[i] = pb.arrPembeli[pb.totPb-1]
				pb.totPb--
				fmt.Println("Pembeli ditolak!")
			} else {
				fmt.Println("Pilihan tidak valid")
			}
			return
		}
	}
	fmt.Println("Pembeli tidak ditemukan!")
}
