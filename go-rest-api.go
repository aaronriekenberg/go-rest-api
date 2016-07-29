package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

type SampleResponse struct {
	Id    string    `json:"id"`
	SubID string    `json:"subID"`
	Time  time.Time `json:"time"`
}

func TestTopLevelHandler(w http.ResponseWriter, r *http.Request) {
	logger.Printf("testTopLevelHandler %v", r)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TestIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logger.Printf("testIDHandler %v", r)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(vars["id"]))
}

func TestSubIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logger.Printf("testSubIDHandler %v", r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := SampleResponse{
		Id:    vars["id"],
		SubID: vars["subID"],
		Time:  time.Now(),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Printf("err %v", err)
		panic(err)
	}
}

func main() {
	logger.Printf("starting")
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/test/v1").Subrouter()
	s.HandleFunc("/", TestTopLevelHandler)
	s.HandleFunc("/{id}", TestIDHandler)
	s.HandleFunc("/{id}/sub/{subID}", TestSubIDHandler)
	logger.Fatal(http.ListenAndServe(":8080", r))
}
