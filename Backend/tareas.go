package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var BaseDeDatos *sql.DB

type ParámetrosDeLaBaseDeDatos struct {
	Conductor       string
	NombreDeUsuario string
	Contraseña      string
	Nombre          string
}

func LeerVariablesDeEntorno() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	DatosDeLaBaseDeDatos.Conductor = os.Getenv("ControladorDeBaseDeDatos")
	DatosDeLaBaseDeDatos.NombreDeUsuario = os.Getenv("NombreDeUsuarioDeLaBaseDeDatos")
	DatosDeLaBaseDeDatos.Contraseña = os.Getenv("ContraseñaDeLaBaseDeDatos")
	DatosDeLaBaseDeDatos.Nombre = os.Getenv("NombreDeLaBaseDeDatos")
}

func ConectarALaBaseDeDatos() (db *sql.DB, err error) {
	return sql.Open(DatosDeLaBaseDeDatos.Conductor, DatosDeLaBaseDeDatos.NombreDeUsuario+":"+DatosDeLaBaseDeDatos.Contraseña+"@/"+DatosDeLaBaseDeDatos.Nombre)
}

func RegistrarSolicitudes(r *http.Request) {
	log.Printf("%s %s %s\n\n\n", r.RemoteAddr, r.Method, r.URL)
}

//
func ObtenerElementosDeTareas(sesiónIdentificaciónNúmero string) ([]ElementoDeTarea, error) {
	resultado := make([]ElementoDeTarea, 0)

	// Conectar a la base de datos.
	BaseDeDatos, err := ConectarALaBaseDeDatos()
	if err != nil {
		return nil, err
	}

	// Obtener datos de la base de datos
	filas, err := BaseDeDatos.Query("SELECT `tareas`.`Nombre`, `tareas`.`Descripción`, `tareas`.`Prioridad`, `tareas`.`Hecho`, `tareas`.`Plazo`, `tareas`.`CreadoEn` FROM `tareas` LEFT JOIN `usuarios` ON `tareas`.`usuarioIdentificaciónNúmero` = `usuarios`.`usuarioIdentificaciónNúmero` WHERE `sesiónIdentificaciónNúmero` = '" + sesiónIdentificaciónNúmero + "';")
	if err != nil {
		return nil, err
	}
	// :)
	for filas.Next() {
		var Nombre string
		var Descripción string
		var Prioridad PrioridadEnum
		var Hecho bool
		var Plazo string
		var CreadoEn string

		err := filas.Scan(
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
