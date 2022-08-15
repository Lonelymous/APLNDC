package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Dirección string = "0.0.0.0:4000"

func EnviarRespuesta(w http.ResponseWriter, aporte any) {
	datos, err := json.Marshal(aporte)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(datos)
}

func SolicitudDeDecodificación(w http.ResponseWriter, r *http.Request, aporte any) bool {
	descifrador := json.NewDecoder(r.Body)

	// Decodificar la entrada.
	err := descifrador.Decode(aporte)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func main() {
	fmt.Println("El servidor se está ejecutando en " + Dirección)
	mux := mux.NewRouter()
	mux.HandleFunc("/tasks", Controladores_Tareas).Methods("GET", "POST")
	mux.HandleFunc("/tasks/{id:[0-9]+}", Controladores_Tareas_IdentificaciónNúmero).Methods("GET", "PATCH", "DELETE")
	http.ListenAndServe(Dirección, mux)
}
