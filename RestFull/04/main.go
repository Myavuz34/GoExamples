package main

import "net/http"

func main() {

	http.HandleFunc("/watch", search)
	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {

	v := r.FormValue("v")

	data := "/watch?v=" + v

	w.Write([]byte(data))
}

//https://www.youtube.com/watch?v=tIfJh88EYPE
