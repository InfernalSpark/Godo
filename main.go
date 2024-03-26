package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "niga re")
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}