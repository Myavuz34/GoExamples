package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//Default ve boş yapıcı metot
func NewHuman() *Human {

	h := new(Human)
	return h
}

func NewHumanWithParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age

	return h
}

func main() {

	//Nesne oluşturma

	//V1
	// x := Human{FirstName: "Mustafa", LastName: "Yavuz", Age: 30}
	// fmt.Println(x.FirstName, x.LastName, x.Age)

	//V2
	// x := new(Human)
	// x.FirstName = "Mustafa"
	// x.LastName = "Yavuz"
	// x.Age = 30

	///V3
	// x := NewHuman()
	// x.FirstName = "Mustafa"
	// x.LastName = "Yavuz"
	// x.Age = 30
	// fmt.Println(x.Age)

	//V4
	x := NewHumanWithParams("Mustafa", "Yavuz", 30)
	fmt.Println(x.FirstName, x.LastName, x.Age)

}
