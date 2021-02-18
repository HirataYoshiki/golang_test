package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Document struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Contents string `json:"contents"`
}

type Documents []Document

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "HomePage EndPoint")
}
func allDocuments(w http.ResponseWriter, r *http.Request) {
  documents := Documents{
    Document{Title: "FMEA", Desc: "some description", Contents: "some contents"},
    Document{Title: "D-FMEA", Desc: "some description", Contents: "some contents"},
  }

  json.NewEncoder(w).Encode(documents)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/documents", allDocuments)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main() {
  handleRequests()
}