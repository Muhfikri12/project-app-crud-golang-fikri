package main

import (
	"app/service"
	"app/utils"
	"app/view"
	"context"
	"fmt"
)

func main() {

	// Login Section
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		username, password := view.DisplayLoginPrompt()

		if service.Login(ctx, username, password) {
			view.DisplayMessage("Login berhasil!")
			view.DisplayMessage("Selamat datang, " + username)
			break
		} else {
			view.DisplayMessage("Login gagal. Coba lagi.")
		}

		// Cek Jika ingin keluar aplikasi
		if view.AskToExit() {
			return
		}
	}

	for {
		utils.ClearScreen()
		fmt.Println(utils.ColorMessage("blue", "=== Perpustakaan ==="))
		fmt.Println("1. Daftar Buku")
		fmt.Println("2. Pinjam Buku")
		fmt.Println("3. Riwayat Peminjaman")
		fmt.Println("4. Kembalikan Buku")
		fmt.Println(utils.ColorMessage("red","5. Keluar" ))
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			utils.ClearScreen()
			service.GetBook()
			utils.PromptToReturnToMenu()

		case 2:
			// Meminjam buku
			utils.ClearScreen()
			service.GetBook()
			fmt.Println(utils.ColorMessage("red","\n99. Kembali Ke Menu"))
		
			var bookID int
			var userName string
		
			fmt.Print("Masukkan ID buku yang ingin dipinjam (atau 99 untuk kembali ke menu): ")
			fmt.Scanln(&bookID)
		
			// Mengecek Kondisi untuk kembali ke Menu
			if bookID == 99 {
				continue 
			}
		
			fmt.Print("Masukkan nama peminjam: ")
			fmt.Scanln(&userName)
		
			// Memanggil fungsi Borrower dengan ID buku dan nama peminjam
			service.Borrower(bookID, userName)
			utils.PromptToReturnToMenu()		
		case 3:
			// Menampilkan riwayat peminjaman
			service.BookHistory()
			utils.PromptToReturnToMenu()
		case 4:
			utils.ClearScreen()
			fmt.Println("--- Daftar Nama Peminjam ---")
			service.ListBorrwer()
			var borrowerName string
			fmt.Print("Masukkan nama peminjam yang ingin mengembalikan buku: ")
			fmt.Scanln(&borrowerName)

			// Memanggil fungsi ReturnBook dengan nama peminjam
			service.ReturnBook(borrowerName)
			utils.PromptToReturnToMenu()
		case 5:
			// Keluar dari program
			fmt.Println("Terima kasih telah menggunakan perpustakaan.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			utils.PromptToReturnToMenu()
		}
	}
}
