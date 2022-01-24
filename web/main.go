package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "os"
  "golang_test/dbinterface"
  "golang_test/resources"
)

mysql = resources.MySQLClient{
          Host: os.Getenv("DBHOST"),
          Port: os.Getenv("DBPORT"),
          User: os.Getenv("MYSQLUSER"),
          Pass: os.Getenv("MYSQLPASS"),
          DB:   os.Getenv("DB"),
        }

// type Document struct {
//   Title string `json:"Title"`
//   Desc string `json:"desc"`
//   Contents string `json:"contents"`
// }

// type Documents []Document

func homePage(w http.ResponseWriter, r *http.Request) {
  uri, err := mysql.getURI()
  if err != nil {
    fmt.Fprintf(w, "Error")
  }
  fmt.Fprintf(w, uri)
}

func createAndRead(w http.ResponseWriter, r *http.Request) {
  dbinterface.Post()

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