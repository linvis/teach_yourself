package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")

	cookie := http.Cookie{Name: "fuckyou"}
	fmt.Println()
	http.SetCookie(w, &cookie)
}
