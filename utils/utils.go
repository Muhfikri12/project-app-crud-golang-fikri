package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		module := exec.Command("cmd", "/c", "cls")
		module.Stdout = os.Stdout
		module.Run()
	} else {
		module := exec.Command("clear")
		module.Stdout = os.Stdout
		module.Run()
	}
}

func PromptToReturnToMenu() {
	fmt.Println(ColorMessage("yellow", "\nKlik Enter untuk kembali ke menu..."))
	fmt.Scanln()
}
