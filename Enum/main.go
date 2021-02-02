package main

import "fmt"

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE    Brand = "Google"
	YAVUZ     Brand = "Yavuz"
)

func PrintBrand(brand Brand) {
	fmt.Println(brand)

}
func main() {

	PrintBrand(YAVUZ)
	PrintBrand(GOOGLE)

}
