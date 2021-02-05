package main

import (
	"fmt"
	"os"
)

func main() {

	// myError := errors.New("Bu bir hata mesajıdır")
	// fmt.Println(myError)

	_, err := os.Open("abc.rar")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		// log.Fatal("Error : ", err.Error)
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
