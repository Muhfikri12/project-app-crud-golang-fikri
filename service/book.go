package service

import (
	"app/model"
	"app/utils"
	"encoding/json"
	"fmt"
	"os"
)

func GetBook() {
	file, err := os.Open("book.json")
	if err != nil {
		utils.Error("Error Opening File,", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var books []model.Books
	if err := decoder.Decode(&books); err != nil {
		fmt.Println(utils.ColorMessage("red", "Daftar Produk Tidak Tersedia !"))
		return
	}

	fmt.Println("Daftar Buku Tersedia:")
	for i, book := range books {
		// Only display the book if it has not been borrowed
		if !book.IsBorrowed {
			fmt.Printf("%v. %s, Ditulis Oleh: %s\n", i+1, book.Name, book.Author)
		}
	}
}

func GetBorrowHistory() ([]model.Borrower, error) {
	
	file, err := os.Open("history.json")
	if err != nil {
		return nil, fmt.Errorf("error opening history file: %w", err)
	}
	defer file.Close()

	var borrowHistory []model.Borrower
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&borrowHistory); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	var filterName string
	fmt.Print("\nMasukkan Nama Anda: ")
	fmt.Scan(&filterName)

	var filteredHistory []model.Borrower
	for _, entry := range borrowHistory {
		if entry.Name == filterName {
			filteredHistory = append(filteredHistory, entry)
			entry.Status = false
		}
	}

	if len(filteredHistory) == 0 {
		fmt.Print("Maaf, nama Anda tidak terdaftar dalam peminjaman. \n\n")
		return nil, nil
	}

	return filteredHistory, nil
}

