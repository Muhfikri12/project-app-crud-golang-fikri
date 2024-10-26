package view

import (
	"app/utils"
	"fmt"
)

func DisplayLoginPrompt() (string, string) {
	utils.ClearScreen()
	fmt.Print("Masukkan username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Masukkan password: ")
	var password string
	fmt.Scanln(&password)

	return username, password
}

func DisplayMessage(message string) {
	fmt.Println(message)
}

func AskToExit() bool {
	fmt.Print("Ingin keluar? (y/n): ")
	var exitChoice string
	fmt.Scanln(&exitChoice)
	return exitChoice == "y"
}

