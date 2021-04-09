package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// String birleştirme işlemi
	var builder bytes.Buffer
	builder.WriteString("Merhaba Dünya;\n")
	builder.WriteString("Bu benim yeni sayfam\n")
	builder.WriteString("Bu bir test messajıdır\n")

	uri := "www.myavuz.net"

	page := Page{
		Title:           "Mustafa Yavuz Blog Sayfası",
		Author:          "Mustafa Yavuz",
		Header:          "Golang Programlama Dili",
		PageDescription: "Go Programlama dili öğreniyorum",
		Content:         builder.String(),
		URI:             "http://" + uri}
	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
