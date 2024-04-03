package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HelloHandler endpoint")

	fmt.Fprint(w, "Hello endpoint")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("FormHandler endpoint")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request successful")

	fmt.Printf("Form: %v\n", r.Form)

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
