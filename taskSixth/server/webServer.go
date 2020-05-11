package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type request struct {
	Host       string  `json:"host"`
	UserAgent  string  `json:"user_agent"`
	RequestURI string  `json:"request_uri"`
	Headers    headers `json:"headers"`
}

type headers struct {
	Accept    string `json:"Accept"`
	UserAgent string `json:"User-Agent"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	req := &request{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Headers:    headers{r.Header.Get("Accept"), r.Header.Get("User-Agent")},
	}
	request, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	w.Write(request)
}
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
