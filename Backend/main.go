package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Dirección string = "0.0.0.0:4000"
var DatosDeLaBaseDeDatos ParámetrosDeLaBaseDeDatos

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
	fmt.Println("Cargar parámetros de la base de datos..")
	LeerVariablesDeEntorno()
	fmt.Println("El servidor se está ejecutando en " + Dirección + "\n\n")
	mux := mux.NewRouter()
	mux.HandleFunc("/tasks", Controladores_Tareas).Methods("GET", "POST")
	mux.HandleFunc("/tasks/{id:[0-9]+}", Controladores_Tareas_IdentificaciónNúmero).Methods("GET", "PATCH", "DELETE")
	mux.HandleFunc("/login", Controladores_IniciarSesión).Methods("POST")
	// mux.HandleFunc("/register", Controladores_Registro).Methods("POST")
	http.ListenAndServe(Dirección, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(mux))

}
