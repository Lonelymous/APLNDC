package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Address string = "127.0.0.1:4000"

func SendResponse(w http.ResponseWriter, i any) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}
func main() {
	fmt.Println("The server is running on " + Address)
	mux := mux.NewRouter()
	mux.HandleFunc("/todos", Controllers_Todos).Methods("GET", "POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", Controllers_Todos_Id).Methods("GET", "PATCH", "DELETE")
	http.ListenAndServe(Address, mux)
}
