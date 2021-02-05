package main

import "fmt"

func main() {

	myArray1 := [...]int{3, 5, 7}
	mySlice1 := myArray1[:]
	fmt.Println(mySlice1)
	mySlice1[0] = 11
	fmt.Println(mySlice1)
	fmt.Println(myArray1)

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "Yellow")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

}
