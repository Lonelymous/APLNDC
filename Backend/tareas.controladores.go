package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type PrioridadEnum string

const (
	Critical            PrioridadEnum = "critical"
	Importante          PrioridadEnum = "importante"
	Sérieux             PrioridadEnum = "sérieux"
	AnnyiraNemQvaFontos PrioridadEnum = "annyira nem qva fontos"
	Unwichtig           PrioridadEnum = "unwichtig"
)

type ElementoDeUsuario struct {
	UsuarioIdentificaciónNúmero uint   `json:"usuarioIdentificaciónNúmero"`
	Nombre                      string `json:"nombre"`
	ContraseñaHash              string `json:"contraseñaHash"`
	Sal                         string `json:"sal"`
}

type EditorDeElementoDeUsuario struct {
	Nombre     string `json:"nombre"`
	Contraseña string `json:"contraseña"`
}

type ElementoDeTarea struct {
	Nombre      string        `json:"nombre"`
	Descripción string        `json:"descripción"`
	Prioridad   PrioridadEnum `json:"prioridad"`
	Hecho       bool          `json:"hecho"`
	Plazo       string        `json:"plazo"`
	CreadoEn    string        `json:"creadoEn"`
}

type ElementoDeTarea2 struct {
	TareaIdentificaciónNúmero   uint          `json:"tareaIdentificaciónNúmero"`
	UsuarioIdentificaciónNúmero uint          `json:"usuarioIdentificaciónNúmero"`
	Nombre                      string        `json:"nombre"`
	Descripción                 string        `json:"descripción"`
	Prioridad                   PrioridadEnum `json:"prioridad"`
	Hecho                       bool          `json:"hecho"`
	Plazo                       string        `json:"plazo"`
	CreadoEn                    string        `json:"creadoEn"`
}

type EditorDeElementosDeTareas struct {
	UsuarioIdentificaciónNúmero uint          `json:"usuarioIdentificaciónNúmero"`
	Nombre                      string        `json:"nombre"`
	Descripción                 string        `json:"descripción"`
	Prioridad                   PrioridadEnum `json:"prioridad"`
	Hecho                       bool          `json:"hecho"`
	Plazo                       string        `json:"plazo"`
}

func Controladores_Tareas(w http.ResponseWriter, r *http.Request) {
	RegistrarSolicitudes(r)
	switch r.Method {
	case http.MethodGet:

		sesiónIdentificaciónNúmero := r.Header.Get("sessionId")
		if len(sesiónIdentificaciónNúmero) != 16 {
			fmt.Println("Invalid sesssionId")
			http.Error(w, "Invalid sesssionId", http.StatusNotAcceptable)
		}

		resultado, err := ObtenerElementosDeTareas(sesiónIdentificaciónNúmero)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		EnviarRespuesta(w, resultado)
	case http.MethodPost:
		var solicitud EditorDeElementosDeTareas
		//
		if !SolicitudDeDecodificación(w, r, &solicitud) {
			return
		}
		//
		ClavePrimaria, err := CrearElementoDeTarea(solicitud)
		//
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//
		EnviarRespuesta(w, ClavePrimaria)
	}
}

func Controladores_Tareas_IdentificaciónNúmero(w http.ResponseWriter, r *http.Request) {
	RegistrarSolicitudes(r)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Controllers_Tasks_Id Get"))
	case http.MethodPatch:
		w.Write([]byte("Controllers_Tasks_Id Patch"))
	case http.MethodDelete:
		w.Write([]byte("Controllers_Tasks_Id Delete"))
	}
}

func Controladores_IniciarSesión(w http.ResponseWriter, r *http.Request) {
	RegistrarSolicitudes(r)

	// password + salt => hashed password
	// hashed password - salt => password

	var solicitud EditorDeElementoDeUsuario
	//
	if !SolicitudDeDecodificación(w, r, &solicitud) {
		return
	}

	fmt.Println(solicitud.Nombre)
	fmt.Println(solicitud.Contraseña)

	/*

		{
			"nombre":"admin",
			"contraseña":"admin123"
		}
	*/

	// átkéne konvertálni
	BaseDeDatos, err := ConectarALaBaseDeDatos()
	if err != nil {
		return
	}
	filas, err := BaseDeDatos.Query("SELECT * FROM `Usuarios` WHERE `Nombre` = '" + solicitud.Nombre + "';")
	if err != nil {
		return
	}
	for filas.Next() {
		var Nombre string
		var UsuarioIdentificaciónNúmero uint
		filas.Scan(&Nombre, &UsuarioIdentificaciónNúmero)
		fmt.Println("Nombre\t" + Nombre)
		fmt.Println("UsuarioIdentificaciónNúmero\t" + strconv.FormatUint(uint64(UsuarioIdentificaciónNúmero), 10))

	}
}

// func Controladores_Registro(w http.ResponseWriter, r *http.Request) {
// 	RegistrarSolicitudes(r)
// 	// Conenct to database.
// 	BaseDeDatos, err := ConectarALaBaseDeDatos()
// 	if err != nil {
// 		return 0, err
// 	}

// 	valores := "'" + strconv.FormatUint(uint64(editor.UsuarioIdentificaciónNúmero), 10)
// 	valores += "', '" + editor.Nombre
// 	valores += "', '" + editor.Descripción
// 	valores += "', '" + string(editor.Prioridad)
// 	valores += "', " + strconv.FormatBool(editor.Hecho)
// 	valores += ", '" + editor.Plazo + "'"

// 	filas, err := BaseDeDatos.Exec("INSERT INTO `Usuarios`(`UsuarioIdentificaciónNúmero`,`Nombre`,`Descripción`,`Prioridad`,`Hecho`,`Plazo`) VALUES(" + valores + ");")
// 	if err != nil {
// 		return 0, err
// 	}
// 	ClavePrimaria, err := filas.LastInsertId()
// 	return uint(ClavePrimaria), err
// }
