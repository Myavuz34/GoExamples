package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("Connection open: %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection Open: %v\n", isConnected)

}

func databaseProcessing() {
	defer disConnect()
	connect()
	fmt.Println("Deferring disconnect!")
	fmt.Printf("Connection Open: %v\n", isConnected)
	fmt.Println("Doing something")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to databased")
}

func disConnect() {
	isConnected = false
	fmt.Println("Disconnected")
}
