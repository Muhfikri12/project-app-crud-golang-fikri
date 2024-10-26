package utils

import "fmt"

func Error(text string,err error)  {
	if err != nil {
		fmt.Println(text, err)
		return
	}
}