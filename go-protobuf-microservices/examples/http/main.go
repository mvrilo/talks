package main

import (
	"log"
	"net/http"
)

func main() {
	res := `{ "message": "hi picpay" }`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(res))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
