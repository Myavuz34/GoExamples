package main

import "net/http"

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/Index", IndexHandler)
	http.HandleFunc("/Contact", AboutHandler)

	http.ListenAndServe(":8080", nil)

}

//IndexHandler sayfası
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MErhaba Dünya"))
	x := r.URL.Path[1:]
	data := ""
	if len(x) > 0 {
		data = " Naaber " + x + "!"
	} else {
		data = " Index Page"
	}
	w.Write([]byte(data))

	w.WriteHeader(http.StatusOK)
}

//AboutHandler sayfası
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))

}
