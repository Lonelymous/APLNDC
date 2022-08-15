package main

import (
	"net/http"
)

type PrioridadEnum string

const (
	Critical            PrioridadEnum = "critical"
	Importante          PrioridadEnum = "importante"
	Sérieux             PrioridadEnum = "sérieux"
	AnnyiraNemQvaFontos PrioridadEnum = "annyira nem qva fontos"
	Unwichtig           PrioridadEnum = "unwichtig"
)

type ElementoDeTarea struct {
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
	switch r.Method {
	case http.MethodGet:
		resultado, err := ObtenerElementosDeTareas()
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
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Controllers_Tasks_Id Get"))
	case http.MethodPatch:
		w.Write([]byte("Controllers_Tasks_Id Patch"))
	case http.MethodDelete:
		w.Write([]byte("Controllers_Tasks_Id Delete"))
	}
}
