package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FName string
	LName string
	Age   int
}

func (hmn Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hmn.FName = "Mustafa"
	hmn.LName = "Yavuz"
	hmn.Age = 30

	//Formu parse etmek için kullanılır.
	r.ParseForm()

	//Sunucudan Form bilgisini almak için kullanılır
	fmt.Println(r.Form)

	//Url'den bilgi almak için kullanılır.
	fmt.Println("path", r.URL.Path)

	fmt.Fprint(w, "<table><tr><td>Ad</td><td>SoyAd</td><td>Yaş</td></tr><tr><td>"+hmn.FName+"</td><td>"+hmn.LName+"</td><td>"+strconv.Itoa(hmn.Age)+"</td></tr></table><br>Path: "+r.URL.Path+"")

}

func main() {

	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
