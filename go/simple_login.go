package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)

	http.ListenAndServe("localhost:8000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		if len(r.Form["username"]) == 0 {
			fmt.Println("invalid")
			return
		}

		if _, err := strconv.Atoi(r.Form.Get("username")); err == nil {
			fmt.Println("all nums")
			return
		}

		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("username")); m {
			fmt.Println("username chinese word")
			return
		}

		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("password")); m {
			fmt.Println("password all nums")
			return
		}
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}
