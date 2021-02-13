package main

func main() {
	// message := "Hi"
	// sayHello(message)
	// println(message)
	// println(add(4, 5))
	// numTers, sum := addTerm(3, 4)
	// println("added : ", numTers, "terms forr a total of ", sum)

	// fmt.Println(split(17))

	addFunc := func(terms ...int) (numTers int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTers = len(terms)
		return
	}

	numTers, sum := addFunc(4, 6)
	println("Added :", numTers, "terms for a total of", sum)
}

// func sayHello(message string) {
// 	println("Hello World")
// 	message = "Hi Go"
// }

// func add(terms ...int) (int, int) {
// 	resullt := 0
// 	for _, term := range terms {
// 		resullt += term
// 	}
// 	return len(terms), resullt
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func addTerm(terms ...int) (numTers int, sum int) {
// 	for _, term := range terms {
// 		sum += term
// 	}
// 	numTers = len(terms)
// 	return
// }
