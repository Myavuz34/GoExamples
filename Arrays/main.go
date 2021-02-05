package main

import "fmt"

func main() {

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	array := [...]int{1, 2, 3, 4}
	fmt.Println(array)

	var cars = [...]string{"Tesla", "Mercedes", "Audi"}

	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	i := 0
	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}

	for i, val := range cars {
		fmt.Println("i = ", i, "value = ", val)
	}

}
