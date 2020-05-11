package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	}
	if r.Method == "POST" {
		name := r.FormValue("name")
		address := r.FormValue("address")
		token := name + ":" + address
		cookie := http.Cookie{Name: "token", Value: token}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
