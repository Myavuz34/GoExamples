package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	. "../models"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	page := Page{ID: 3, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	users := loadUsers()
	interests := loadInterests()
	interestMappings := loadInterestMappings()

	var newUsers []User

	for _, user := range users {

		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}

	t, _ := template.ParseFiles("template/page.html")
	t.Execute(w, viewModel)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadUsers() []User {
	bytes, _ := ioutil.ReadFile("json/users.json")
	var users []User
	json.Unmarshal(bytes, &users)
	return users
}

func loadInterests() []Interest {
	bytes, _ := ioutil.ReadFile("json/interests.json")
	var interests []Interest
	json.Unmarshal(bytes, &interests)
	return interests
}

func loadInterestMappings() []InterestMapping {
	bytes, _ := ioutil.ReadFile("json/userInterestMappings.json")
	var interestMappings []InterestMapping
	json.Unmarshal(bytes, &interestMappings)
	return interestMappings
}
