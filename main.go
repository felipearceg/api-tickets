/*
Api para relizar CRUD sobre la tabla tiquetes, que tiene las filas: id, usuario, fechaCreado, fechaActualizado y estado.
Autor: Felipe Arce Giraldo.
Fecha: abril 2021
Versión: 1.0
*/
package main

/*
Importación de paquetes necesarios.
*/
import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
Estructura Tiquete que hace referencia a la tabla de la base de datos.
*/
type Ticket struct {
	Id               int
	Usuario          int
	FechaCreado      string
	FechaActualizado string
	Estado           string
}

/*
Constantes para la conexión a la base de datos.
*/
const (
	usuario         = "root"
	pass            = "password"
	hostname        = "127.0.0.1:3306"
	nombreBaseDatos = "apitickets"
)

// Variable de tipo plantilla que se usa para acceder a las hojas html.
var plantilla *template.Template

// Variable de tipo sql para realizar la conexión a la base de datos.
var baseDatos *sql.DB

/*
Función principal donde se realiza la conexión a la base de datos y se maneja la navegación entre páginas.
*/
func main() {

	plantilla, _ = template.ParseGlob("plantillas/*.html")
	var err error
	baseDatos, err = sql.Open("mysql", dsn())

	if err != nil {
		panic(err)
	}

	defer baseDatos.Close()

	http.HandleFunc("/", inicio)
	http.HandleFunc("/insertar", insertar)
	http.HandleFunc("/porid", porId)
	http.HandleFunc("/buscar", buscar)
	http.HandleFunc("/actualizar/", actualizar)
	http.HandleFunc("/actualizarresultado/", actualizarResultado)
	http.HandleFunc("/borrar/", borrar)
	http.ListenAndServe("localhost:8080", nil)

}

/*
Función para retornar el nombre con el que se va a realizar la conexión a la base de datos.
Usa los valores constantes creados anteriormente. DSN: Data Source Name.
*/
func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", usuario, pass, hostname, nombreBaseDatos)
}

/*
Función para redireccionar / a /buscar
*/
func inicio(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/buscar", 307)
}

/*
Función para insertar un nuevo tiquete.
*/
func insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		plantilla.ExecuteTemplate(w, "insertar.html", nil)
		return
	}

	//Se extraen los valores que ingresa el usuario.
	r.ParseForm()
	usuario := r.FormValue("usuario")
	estado := r.FormValue("estado")
	var err error

	//Si los valores son nulos, entonces se solicitan llenar todos los campos.
	if usuario == "" || estado == "" {
		plantilla.ExecuteTemplate(w, "insertar.html", "Error. Llene todos los campos")
		return
	}

	//Variable para ejecutar la inserción.
	var ins *sql.Stmt

	//Se inicializa la inserción.
	ins, err = baseDatos.Prepare("INSERT INTO apitickets.tiquetes (usuario, fechacreado, fechaactualizado, estado) VALUES (?, ?, ?, ?);")
	if err != nil {
		panic(err)
	}

	//Una vez terminada la función se cierra la variable que estaba insertando en la tabla.
	defer ins.Close()

	t := time.Now()
	fecha := t.Format("02 Jan 06 15:04:05")

	//Se ejecuta la inserción con los datos obtenidos.
	res, err := ins.Exec(usuario, fecha, fecha, estado)

	filasAfectadas, _ := res.RowsAffected()
	if err != nil || filasAfectadas != 1 {
		plantilla.ExecuteTemplate(w, "resultado.html", "Error insertando.")
		return
	}

	plantilla.ExecuteTemplate(w, "resultado.html", "Insertado.")

}

/*
Función para buscar un tiquete por su id.
*/
func porId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		plantilla.ExecuteTemplate(w, "porid.html", nil)
		return
	}

	//Se extraen los valores que ingresa el usuario
	r.ParseForm()
	id := r.FormValue("id")

	//Si los valores son nulos, entonces se solicitan llenar todos los campos.
	if id == "" {
		plantilla.ExecuteTemplate(w, "resultado.html", "Error. Llene todos los campos")
		return
	}

	//Se hace una consulta de fila.
	consulta := baseDatos.QueryRow("SELECT * FROM apitickets.tiquetes WHERE id = ?;", id)

	//Se crea una variable de tipo tiqute
	var t Ticket

	//Se escanean los destinos de la variable t
	err := consulta.Scan(&t.Id, &t.Usuario, &t.FechaCreado, &t.FechaActualizado, &t.Estado)
	if err != nil {
		http.Redirect(w, r, "/buscar", 307)
		return
	}

	//El resultado de la consulta se almacena temporalmente para mostrarlo al usuario.
	//var salida string
	//salida = "ID: " + strconv.Itoa(t.Id) + ", ID Usuario: " + strconv.Itoa(t.Usuario) + ", Estado: " + t.Estado + ", Creado: " + t.FechaCreado + ", Actualizado: " + t.FechaActualizado
	plantilla.ExecuteTemplate(w, "porid.html", t)
}

/*
Función para buscar todos lo tiqutes.
*/
func buscar(w http.ResponseWriter, r *http.Request) {
	consulta := "SELECT * FROM apitickets.tiquetes;"
	filas, err := baseDatos.Query(consulta)
	if err != nil {
		panic(err)
	}

	defer filas.Close()
	var tiquetes []Ticket

	//Ciclo para iterar en todos los resultados de las consutlas
	for filas.Next() {

		//Se crea una variable tipo tiquete para agregarla a la lista de tiquetes
		var t Ticket
		err = filas.Scan(&t.Id, &t.Usuario, &t.FechaCreado, &t.FechaActualizado, &t.Estado)
		if err != nil {
			panic(err)
		}
		tiquetes = append(tiquetes, t)
	}

	//Se retorna una lista de tiquetes para mostrarle al usuario.
	plantilla.ExecuteTemplate(w, "buscar.html", tiquetes)
}

/*
Función para obtener la información del tiquete a editar.
*/
func actualizar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	fila := baseDatos.QueryRow("SELECT * FROM apitickets.tiquetes WHERE id = ?;", id)
	var t Ticket
	err := fila.Scan(&t.Id, &t.Usuario, &t.FechaCreado, &t.FechaActualizado, &t.Estado)
	if err != nil {
		http.Redirect(w, r, "/buscar", 307)
		return
	}
	plantilla.ExecuteTemplate(w, "actualizar.html", t)
}

/*
Función para editar un tiquete del cual ya se obtuvo su información.
*/
func actualizarResultado(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	usuario := r.FormValue("usuario")
	estado := r.FormValue("estado")
	t := time.Now()
	fecha := t.Format("02 Jan 06 15:04:05")

	consulta := "UPDATE apitickets.tiquetes SET usuario = ?, estado = ?, fechaactualizado = ? WHERE (id = ?);"

	salida, err := baseDatos.Prepare(consulta)
	if err != nil {
		panic(err)
	}
	defer salida.Close()
	var res sql.Result
	res, err = salida.Exec(usuario, estado, fecha, id)

	filasAfectadas, _ := res.RowsAffected()
	if err != nil || filasAfectadas != 1 {
		plantilla.ExecuteTemplate(w, "resultado.html", "Error actualizando.")
		return
	}

	plantilla.ExecuteTemplate(w, "resultado.html", "Actualizado.")
}

/*
Función para eliminar un tiquete.
*/
func borrar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	del, err := baseDatos.Prepare("DELETE FROM apitickets.tiquetes WHERE (id = ?);")
	if err != nil {
		panic(err)
	}
	defer del.Close()

	_, err = del.Exec(id)

	if err != nil {
		return
	}
	plantilla.ExecuteTemplate(w, "resultado.html", "Borrado.")
}
