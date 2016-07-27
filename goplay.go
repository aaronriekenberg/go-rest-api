package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"time"
)

type SampleResponse struct {
	Id    string    `json:"id"`
	SubID string    `json:"subID"`
	Time  time.Time `json:"time"`
}

func TestTopLevelHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("testTopLevelHandler %v", r)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TestIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("testIDHandler %v", r)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(vars["id"]))
}

func TestSubIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("testSubIDHandler %v", r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := SampleResponse{
		Id:    vars["id"],
		SubID: vars["subID"],
		Time:  time.Now(),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("err %v", err)
		panic(err)
	}
}

func main() {
	log.Printf("starting")
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/test/v1").Subrouter()
	s.HandleFunc("/", TestTopLevelHandler)
	s.HandleFunc("/{id}", TestIDHandler)
	s.HandleFunc("/{id}/sub/{subID}", TestSubIDHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
