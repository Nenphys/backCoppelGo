package usuario

import (
	"DB"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Usuario struct {
	ID             string `json:"id,omitempty"`
	NombreCompleto string `json:"nombrecompleto,omitempty"`
	Correo         string `json:"correo,omitempty"`
	Password       string `json:"password,omitempty"`
}

type Articulo struct {
	ID          string `json:"id,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Imagen      string `json:"imagen,omitempty"`
	Precio      string `json:"precio,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

type Pedido struct {
	ID        string     `json:"id,omitempty"`
	Direccion string     `json:"direccion,omitempty"`
	LatLog    *Ubicacion `json:"latlog,omitempty"`
	Fecha     string     `json:"fecha,omitempty"`
	Comprador string     `json:"comprador,omitempty"`
}

type Ubicacion struct {
	Latitud  string `json:"latitud,omitempty"`
	Longitud string `json:"longitud,omitempty"`
}

type LoginStatus struct {
	Status string `json:"status,omitempty"`
}

var usuario []Usuario
var people []Usuario
var articulo []Articulo
var pedido []Pedido

func Init() {

	router := mux.NewRouter()
	// adding example data
	people = append(people, Usuario{ID: "1", NombreCompleto: "row.Str(nombre)", Correo: "row.Str(correo)", Password: "row.Str(password)"})
	articulo = append(articulo, Articulo{ID: "1", Nombre: "Ryan", Imagen: "Ray", Precio: "California", Descripcion: "kdsjgfdjfg"})
	pedido = append(pedido, Pedido{ID: "1", Direccion: "Ryan", LatLog: &Ubicacion{Latitud: "Dubling", Longitud: "California"}, Fecha: "Ray", Comprador: "Ray"})

	//Login
	router.HandleFunc("/login/{correo},{pass}", GetLoginEndpoint).Methods("GET")

	// Usuarios
	router.HandleFunc("/usuario", GetUsuariosEndpoint).Methods("GET")
	router.HandleFunc("/usuario/{id}", GetUsuarioEndpoint).Methods("GET")
	router.HandleFunc("/usuario/{id}", CreateUsuarioEndpoint).Methods("POST")
	router.HandleFunc("/usuario/{id}", DeleteUsuarioEndpoint).Methods("DELETE")

	//Articulos
	router.HandleFunc("/articulo", GetArticuloEndpoint).Methods("GET")
	router.HandleFunc("/articulo/{id}", GetArticulosEndpoint).Methods("GET")
	router.HandleFunc("/articulo/{id}", CreateArticuloEndpoint).Methods("POST")
	router.HandleFunc("/articulo/{id}", DeleteArticuloEndpoint).Methods("DELETE")

	//pedidos
	router.HandleFunc("/pedido", GetPedidoEndpoint).Methods("GET")
	router.HandleFunc("/pedido/{id}", GetPedidoEndpoint).Methods("GET")
	router.HandleFunc("/pedido/{id}", CreatePedidoEndpoint).Methods("POST")
	router.HandleFunc("/pedido/{id}", DeletePedidoEndpoint).Methods("DELETE")

	//Start Listen
	log.Fatal(http.ListenAndServe(":8080", router))

}

//login
func GetLoginEndpoint(w http.ResponseWriter, req *http.Request) {
	login := mux.Vars(req)
	var correo, pass string
	correo = login["correo"]
	println(correo)
	pass = login["pass"]
	println(pass)
	status := database.Login(correo, pass)
	println(status)
	json.NewEncoder(w).Encode(&LoginStatus{Status: status})
}

// EndPoints Usuario
func GetUsuarioEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Usuario{})
}

func GetUsuariosEndpoint(w http.ResponseWriter, req *http.Request) {
	rows, res := database.QuerySelect("usuario")
	id := res.Map("idusuario")
	nombre := res.Map("nombre")
	correo := res.Map("correo")
	password := res.Map("password")
	for _, row := range rows {
		id := strconv.Itoa(row.Int(id))
		usuario = append(usuario, Usuario{ID: id, NombreCompleto: row.Str(nombre), Correo: row.Str(correo), Password: row.Str(password)})
	}
	json.NewEncoder(w).Encode(usuario)
	usuario = nil
}

func CreateUsuarioEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Usuario
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}

func DeleteUsuarioEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

// EndPoints Articulos
func GetArticulosEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range articulo {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Articulo{})
}

func GetArticuloEndpoint(w http.ResponseWriter, req *http.Request) {
	rows, res := database.QuerySelect("articulo")
	id := res.Map("idarticulo")
	nombre := res.Map("nombre")
	imagen := res.Map("imagen")
	precio := res.Map("precio")
	descripcion := res.Map("descripcion")
	for _, row := range rows {
		id := strconv.Itoa(row.Int(id))
		articulo = append(articulo, Articulo{ID: id, Nombre: row.Str(nombre), Imagen: row.Str(imagen), Precio: row.Str(precio), Descripcion: row.Str(descripcion)})
	}
	json.NewEncoder(w).Encode(articulo)
	articulo = nil
}

func CreateArticuloEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Articulo
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	articulo = append(articulo, person)
	json.NewEncoder(w).Encode(articulo)

}

func DeleteArticuloEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range articulo {
		if item.ID == params["id"] {
			articulo = append(articulo[:index], articulo[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articulo)
}

// EndPoints Pedido
func GetPedidosEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range articulo {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pedido{})
}

func GetPedidoEndpoint(w http.ResponseWriter, req *http.Request) {
	rows, res := database.QuerySelect("pedido")
	id := res.Map("idpedido")
	direccion := res.Map("direccion")
	latlong := res.Map("latlong")
	fecha := res.Map("fecha")
	comprador := res.Map("comprador")
	for _, row := range rows {
		id := strconv.Itoa(row.Int(id))
		latLong := strings.Split(row.Str(latlong), ",")
		pedido = append(pedido, Pedido{ID: id, Direccion: row.Str(direccion), LatLog: &Ubicacion{Latitud: latLong[0], Longitud: latLong[1]}, Fecha: row.Str(fecha), Comprador: row.Str(comprador)})
	}
	json.NewEncoder(w).Encode(pedido)
	pedido = nil
}

func CreatePedidoEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var crearPedido Pedido
	_ = json.NewDecoder(req.Body).Decode(&crearPedido)
	crearPedido.ID = params["id"]
	pedido = append(pedido, crearPedido)
	json.NewEncoder(w).Encode(pedido)

}

func DeletePedidoEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range articulo {
		if item.ID == params["id"] {
			pedido = append(pedido[:index], pedido[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(pedido)
}
