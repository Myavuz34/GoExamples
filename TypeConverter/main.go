package main

import (
	"fmt"
	"strconv"
)

func main() {

	var myString = "1"
	var x = 23
	var f float32 = 2.4

	fmt.Println(myString, x, f)

	//string'ten int'e dönüştürme
	number, _ := strconv.Atoi(myString)

	//int'ten string'e dönüştürme
	println("Sonuç: " + strconv.Itoa(number))
	result := number + 2
	println(result)

	var i int = 55
	var f1 float64 = float64(i)
	var u uint = uint(f1)

	println(i, f1, u)
}
