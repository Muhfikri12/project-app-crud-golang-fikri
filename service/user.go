package service

import (
	"app/model"
	"app/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"
)



func Borrower(bookID int, borrowerName string) {
	bookFile, err := os.OpenFile("book.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening book file:", err)
		return
	}
	defer bookFile.Close()

	// Mendekode daftar buku dari book.json
	var books []model.Books
	decoder := json.NewDecoder(bookFile)
	if err := decoder.Decode(&books); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Mencari buku berdasarkan ID dan memeriksa ketersediaannya
	var borrowedBook model.Books
	for i, book := range books {
		if book.ID == bookID {
			if book.IsBorrowed {
				fmt.Printf("Buku '%s' sudah dipinjam.\n", book.Name)
				return
			}
			
			// Menandai buku sebagai dipinjam
			books[i].IsBorrowed = true

			borrowedBook = book
			fmt.Printf("Anda telah berhasil meminjam buku '%s'.\n", book.Name)

			// Memperbarui book.json dengan status terbaru
			bookFile.Seek(0, 0) // Kembali ke awal file
			bookFile.Truncate(0) // Menghapus konten sebelumnya
			encoder := json.NewEncoder(bookFile)
			if err := encoder.Encode(&books); err != nil {
				fmt.Println("Error encoding JSON to book.json:", err)
				return
			}
			
			// Menambahkan data peminjaman ke history.json
			addBorrowerToHistory( borrowerName, borrowedBook)
			return
		}
	}
	fmt.Println("Buku dengan ID tersebut tidak ditemukan.")
}



// Fungsi untuk menambahkan data peminjaman ke history.json
func addBorrowerToHistory(borrowerName string, book model.Books) {
	historyFile, err := os.OpenFile("history.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening history file:", err)
		return
	}
	defer historyFile.Close()

	// Membaca data riwayat peminjaman sebelumnya dari history.json
	var borrowHistory []model.Borrower
	decoder := json.NewDecoder(historyFile)
	if err := decoder.Decode(&borrowHistory); err != nil && err.Error() != "EOF" {
		fmt.Println("Error decoding JSON from history.json:", err)
		return
	}

	// Menemukan ID terbesar saat ini untuk menentukan ID selanjutnya
	var lastID int
	for _, user := range borrowHistory {
		if user.ID > lastID {
			lastID = user.ID
		}
	}

	// Menambahkan data peminjaman baru ke riwayat dengan ID auto increment
	newEntry := model.Borrower{
		ID:         lastID + 1, // Auto-increment ID
		Name:       borrowerName,
		Book:       book,
		Status:     true,
		Created_at: time.Now(),
	}
	borrowHistory = append(borrowHistory, newEntry)

	// Menulis ulang data peminjaman ke history.json
	historyFile.Seek(0, 0)  // Kembali ke awal file
	historyFile.Truncate(0) // Menghapus konten sebelumnya
	encoder := json.NewEncoder(historyFile)
	if err := encoder.Encode(&borrowHistory); err != nil {
		fmt.Println("Error encoding JSON to history.json:", err)
	}
}

func BookHistory() {
	history, err := GetBorrowHistory()
	if err != nil {
		fmt.Println("Error retrieving borrow history:", err)
		return
	}

	jsonData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println("Daftar Peminjaman Buku :")
	fmt.Println("\n", string(jsonData))
}

func ReturnBook(borrowerName string) {
	historyFile, err := os.OpenFile("history.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening history file:", err)
		return
	}
	defer historyFile.Close()

	// Read the existing borrowing history from history.json
	var borrowHistory []model.Borrower
	decoder := json.NewDecoder(historyFile)
	if err := decoder.Decode(&borrowHistory); err != nil {
		fmt.Println("Error decoding JSON from history.json:", err)
		return
	}

	var borrowedBooks []model.Borrower
	bookFound := false // To track if any books were found for the borrower

	// Collect borrowed books for the borrower with Status true
	for _, entry := range borrowHistory {
		if entry.Name == borrowerName && entry.Status {
			bookFound = true // Set to true when a borrowed book is found
			borrowedBooks = append(borrowedBooks, entry) // Collect all borrowed books that are currently borrowed
		}
	}

	if !bookFound {
		// If no borrowing records found for the borrower
		fmt.Printf("Tidak ada peminjaman yang ditemukan untuk nama '%s'.\n", borrowerName)
		return
	}

	// Display only the currently borrowed books
	jsonData, err := json.MarshalIndent(borrowedBooks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling borrowed books to JSON:", err)
		return
	}

	fmt.Println("Daftar buku yang dipinjam:")
	fmt.Println(string(jsonData))

	// Allow borrower to return a book by ID
	var bookID int
	fmt.Print("Masukkan ID buku yang ingin dikembalikan: ")
	fmt.Scanln(&bookID)

	// Process the return of the book
	for i, entry := range borrowHistory {
		if entry.Name == borrowerName && entry.Book.ID == bookID {
			if !entry.Status {
				fmt.Println("Buku sudah dikembalikan sebelumnya.")
				return
			}

			// Mark the borrowing as inactive
			borrowHistory[i].Status = false
			fmt.Printf("Anda telah berhasil mengembalikan buku '%s'.\n", entry.Book.Name)

			// Write back the borrowing data to history.json
			historyFile.Seek(0, 0) // Go back to the beginning of the file
			historyFile.Truncate(0) // Clear previous content
			encoder := json.NewEncoder(historyFile)
			if err := encoder.Encode(&borrowHistory); err != nil {
				fmt.Println("Error encoding JSON to history.json:", err)
			}
			return
		}
	}

	fmt.Printf("Buku dengan ID '%d' tidak ditemukan dalam riwayat peminjaman untuk nama '%s'.\n", bookID, borrowerName)
}

func ListBorrwer()  {
	file, err := os.Open("history.json")
	if err != nil {
		fmt.Println("Error Opening history.json,", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var listBorrwer []model.Borrower

	if err := decoder.Decode(&listBorrwer); err != nil {
		fmt.Println(utils.ColorMessage("red", "Daftar Peminjam Tidak Tersedia !"))
		return
	}

	// Map to store counts of each borrower's name
	nameCount := make(map[string]int)

	// Count occurrences of each borrower's name
	for _, user := range listBorrwer {
		nameCount[user.Name]++
	}

	// Print the list of borrowers with their counts
	for i, count := range nameCount {
		fmt.Printf("Nama: %s, Jumlah Peminjaman: %d\n", i, count)
	}
}