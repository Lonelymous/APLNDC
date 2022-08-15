package main

import (
	"database/sql"
	"strconv"
)

var BaseDeDatos *sql.DB

func ConectarALaBaseDeDatos() (db *sql.DB, err error) {
	ControladorDeBaseDeDatos := "mysql"
	NombreDeUsuarioDeLaBaseDeDatos := "Carlos"
	ContraseñaDeLaBaseDeDatos := "Carlo$123"
	NombreDeLaBaseDeDatos := "APLNDC"

	return sql.Open(ControladorDeBaseDeDatos, NombreDeUsuarioDeLaBaseDeDatos+":"+ContraseñaDeLaBaseDeDatos+"@/"+NombreDeLaBaseDeDatos)
}

//
func ObtenerElementosDeTareas() ([]ElementoDeTarea, error) {
	resultado := make([]ElementoDeTarea, 0)

	// Conectar a la base de datos.
	BaseDeDatos, err := ConectarALaBaseDeDatos()
	if err != nil {
		return nil, err
	}

	// SELECT * FROM `Tareas`;
	// Obtener datos de la base de datos
	filas, err := BaseDeDatos.Query("SELECT * FROM `Tareas`;")
	if err != nil {
		return nil, err
	}
	// :)
	for filas.Next() {
		var TareaIdentificaciónNúmero uint
		var UsuarioIdentificaciónNúmero uint
		var Nombre string
		var Descripción string
		var Prioridad PrioridadEnum
		var Hecho bool
		var Plazo string
		var CreadoEn string

		err := filas.Scan(
			&TareaIdentificaciónNúmero,
			&UsuarioIdentificaciónNúmero,
			&Nombre,
			&Descripción,
			&Prioridad,
			&Hecho,
			&Plazo,
			&CreadoEn)
		if err != nil {
			return nil, err
		}
		resultado = append(resultado, ElementoDeTarea{
			TareaIdentificaciónNúmero,
			UsuarioIdentificaciónNúmero,
			Nombre,
			Descripción,
			Prioridad,
			Hecho,
			Plazo,
			CreadoEn})
	}

	// Devolver datos
	return resultado, nil
}

//
func CrearElementoDeTarea(editor EditorDeElementosDeTareas) (uint, error) {

	// Conenct to database.
	BaseDeDatos, err := ConectarALaBaseDeDatos()
	if err != nil {
		return 0, err
	}

	valores := "'" + strconv.FormatUint(uint64(editor.UsuarioIdentificaciónNúmero), 10)
	valores += "', '" + editor.Nombre
	valores += "', '" + editor.Descripción
	valores += "', '" + string(editor.Prioridad)
	valores += "', " + strconv.FormatBool(editor.Hecho)
	valores += ", '" + editor.Plazo + "'"

	filas, err := BaseDeDatos.Exec("INSERT INTO `Tareas`(`UsuarioIdentificaciónNúmero`,`Nombre`,`Descripción`,`Prioridad`,`Hecho`,`Plazo`) VALUES(" + valores + ");")
	if err != nil {
		return 0, err
	}
	ClavePrimaria, err := filas.LastInsertId()
	return uint(ClavePrimaria), err
}
