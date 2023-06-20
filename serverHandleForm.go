package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello iskander 🚀!") // write data to response
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		//if err != nil {
		//	return
		//}
		// logic part of log in
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "3abbi el username ")
			return
		}
		pass, err := strconv.Atoi(r.Form.Get("username"))

		if err != nil {
			fmt.Println("error =>", err.Error())
			fmt.Fprintf(w, "put number password ")
			return
		}

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", pass)
		fmt.Println("gender", r.Form.Get("gender"))
		fmt.Fprintf(w, "thank you 🚀")
	}
}
func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
