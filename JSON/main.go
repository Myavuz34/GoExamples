package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personel string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("*************************")
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fata Error: ", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(fileName string, key interface{}) {

	outFile, err := os.Create(fileName)
	CheckError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	CheckError(err)
	outFile.Close()
}

func main() {

	person := Person{
		ID:        10,
		FirstName: "Mustafa",
		LastName:  "Yavuz",
		UserName:  "Myavuz",
		Gender:    "E",
		Name:      Name{Family: "bididibid", Personel: "Mustafa"},
		Email: []Email{
			Email{ID: 1, Kind: "Work", Address: "myavuz53@gmail.com"},
			Email{ID: 2, Kind: "Home", Address: "mustafayavuz34@hotmail.com.tr"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Go"},
			Interest{ID: 2, Name: "C#"},
			Interest{ID: 3, Name: "Python"},
		},
	}

	WriteMessage("Reading Operation Started")

	WriteMessage("Personel Fullname")
	WriteStarLine()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personel Emil With Index")
	WriteStarLine()
	resEmail := GetPersonEmailAddress(&person, 1)
	WriteMessage(resEmail)
	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personel Email Object winth Index")
	WriteStarLine()
	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarLine()

	WriteMessage("Reading Operation Ended")
	WriteMessage("\n")
	WriteMessage("Writing Operation Started")
	SaveJSON("person.json", person)
	WriteMessage("Writing Operation Ended")
}
