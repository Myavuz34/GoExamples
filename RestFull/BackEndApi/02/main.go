package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {

	apiRoot := "/api"
	http.HandleFunc(apiRoot, func(rw http.ResponseWriter, r *http.Request) {
		message := API{"API HOME"}
		output, err := json.Marshal(message)
		CheckError(err)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(rw, string(output))
	})

	http.HandleFunc(apiRoot+"/users", func(rw http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, FirstName: "Mustafa", LastName: "Yavuz", Age: 30},
			{ID: 2, FirstName: "Kerem", LastName: "Yavuz", Age: 45},
			{ID: 3, FirstName: "Ali", LastName: "Yavuz", Age: 10},
		}
		message := users
		output, err := json.Marshal(message)
		CheckError(err)
		fmt.Fprintf(rw, string(output))
	})

	http.HandleFunc(apiRoot+"/me", func(rw http.ResponseWriter, r *http.Request) {
		user := User{1, "Mustafa", "Yavuz", 30}
		message := user
		output, err := json.Marshal(message)
		CheckError(err)
		fmt.Fprintf(rw, string(output))
	})

	http.ListenAndServe(":8080", nil)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
}
