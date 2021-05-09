package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome home!")
}

func main(){
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/testing/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}