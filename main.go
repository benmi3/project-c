package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Name string `json:"name"`
}

type ItemInfo struct {
	ID uint64 `json:"id"`
}

func handleInsert(w http.ResponseWriter, r *http.Request) {
	var req Item
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error: decoding json: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := ItemInfo{ID: 1}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("Error: encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/insert", handleInsert).Methods("POST")

	http.ListenAndServe(":8090", nil)
}
