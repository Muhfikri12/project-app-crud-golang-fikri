package utils

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func ErrorMessage(text string) {
	red := color.New(color.FgRed).SprintFunc()
	err := errors.New(text)
	fmt.Printf("⚠️   %s%s\n", red("Error : "), err)
}

func SuccesMessage(text string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("✔️   %s%s\n", green("Succes : "), text)
}

func ColorMessage(color_ string, text string) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	switch color_ {
	case "red":
		return red(text)
	case "green":
		return green(text)
	case "blue":
		return blue(text)
	case "yellow":
		return yellow(text)
	}

	return "0"
}
