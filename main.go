package main

import (
	"fmt"
	"os"
)

func main() {

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	projessorArchiterture := os.Getenv("PROCESSOR_ARCHITECTURE")
	// projessorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	// processorLeve := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("Username :" + uName)
	fmt.Println("Domain : " + uDomain)
	fmt.Println("İşlemci Mimarisi : " + projessorArchiterture)
	fmt.Println("HOMEPATH : " + homePath)
	fmt.Println("GOPATH : " + goPath)
	fmt.Println("GOROOT :" + goRoot)
}
