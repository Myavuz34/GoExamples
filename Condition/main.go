package main

import "fmt"

func main() {

	// IF
	a := 10
	b := 10

	if b > a {
		fmt.Println("Büyüktür")
	} else if b == a {
		fmt.Println("Eşittir")
	} else {
		fmt.Println("küçüktür")
	}

	foo := 1
	if foo == 1 {
		println("bar")
	}

	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	// SWITCH
	foo1 := 33
	switch {
	case foo1 == 1:
		println("One")
	case foo1 == 2:
		println("two")
	case foo1 == 3:
		println("Three")
	default:
		println("Others")
	}

	var score float64
	fmt.Println("Enter score for your last exam :")
	fmt.Scanf("%v", &score)
	switch {
	case score <= 56:
		fmt.Println("Grade is F")
	case score <= 69:
		fmt.Println("Grade is C")
	case score <= 79:
		fmt.Println("Grade is D")
	case score <= 89:
		fmt.Println("Grade is B")
	case score <= 100:
		fmt.Println("Grade is A")
	default:
		fmt.Println("Please enter a score <=100")
	}
}
