package main

import (
	//Para la pagina wed
	"database/sql"
	"log"
	"net/http"

	//buscar registros o plantillas atravez de carpetas
	"text/template"
	//drivers
	//funcion para conexion de la base de datos

	_ "github.com/go-sql-driver/mysql"
)

func conexionDB() *sql.DB {
	conexion, err := sql.Open("mysql", "root:"+""+"@tcp(localhost:3306)/sistema")
	if err != nil {
		panic(err.Error())
	}
	//defer conexion.Close()
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func main() {
	http.HandleFunc("/login", Sesion)
	http.HandleFunc("/registro", Registro)
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actulizar)
	log.Print("Servidor Corriendo..")
	http.ListenAndServe(":5500", nil)
}
func Sesion(w http.ResponseWriter, r *http.Request) {

}
func Registro(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		usuario := r.FormValue("usuario")
		contraseña := r.FormValue("coontraseña")
		conexionEstablecida := conexionDB()
		insertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO usuario (nombre,correo,usuario,contraseña) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistro.Exec(nombre, correo, usuario, contraseña)
		http.Redirect(w, r, "/login", 301)
		//defer conexionEstablecida.Close()
	}
}

//funciones para las llamadas de los templates
func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionDB()
	registros, err := conexionEstablecida.Query("SELECT * FROM empleado")
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	AregloEmpleado := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		//enmarcar los errores
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
		AregloEmpleado = append(AregloEmpleado, empleado)

	}
	//fmt.Println(AregloEmpleado)
	plantillas.ExecuteTemplate(w, "Inicio", AregloEmpleado)
}
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionDB()
		insertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO empleado (nombre,correo) VALUES (?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistro.Exec(nombre, correo)
		http.Redirect(w, r, "/", 301)
		//defer conexionEstablecida.Close()
	}

}

//funcion borrar un elemento de una lista
func Borrar(w http.ResponseWriter, r *http.Request) {
	//recepcion de datos
	idEmpleado := r.URL.Query().Get("id")
	//fmt.Println(idEmpleado)
	conexionEstablecida := conexionDB()
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleado WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)
}
func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	//fmt.Println(idEmpleado)
	conexionEstablecida := conexionDB()
	registro, err := conexionEstablecida.Query("SELECT * FROM empleado WHERE id=?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	for registro.Next() {
		var id int
		var nombre, correo string
		err = registro.Scan(&id, &nombre, &correo)
		//enmarcar los errores
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}
	//fmt.Println(empleado)
	plantillas.ExecuteTemplate(w, "editar", empleado)
}
func Actulizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := conexionDB()

		modificarRegistro, err := conexionEstablecida.Prepare("UPDATE empleado SET nombre=?,correo=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		modificarRegistro.Exec(nombre, correo, id)
		http.Redirect(w, r, "/", 301)
		//defer conexionEstablecida.Close()
	}
}
