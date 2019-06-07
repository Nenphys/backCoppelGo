package database

import (
	"os"
	"strconv"
	"strings"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

// func main(){
//   QueryDelete("pedido","2")
// }
func printOK() {
	println("OK")
}

func checkError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func checkedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row,
	mysql.Result) {
	checkError(err)
	return rows, res
}

func Login(correo string, password string) string {
	user := "root"
	pass := "root"
	dbname := "coppel"
	proto := "tcp"
	addr := "localhost:3307"

	db := mysql.New(proto, "", addr, user, pass, dbname)
	checkError(db.Connect())
	var isAuthenticated string
	row, _ := checkedResult(db.Query("select * from usuario where correo = '" + correo + "' and password = '" + password + "'"))
	if len(row) != 0 {
		isAuthenticated = "Success"
	} else {
		isAuthenticated = "Fail"
	}
	return isAuthenticated
}

func QuerySelect(tabla string) ([]mysql.Row,
	mysql.Result) {
	user := "root"
	pass := "root"
	dbname := "coppel"
	proto := "tcp"
	addr := "localhost:3307"

	db := mysql.New(proto, "", addr, user, pass, dbname)
	checkError(db.Connect())
	rows, res := checkedResult(db.Query("select * from " + tabla))
	switch tabla {
	case "usuario":
		id := res.Map("idusuario")
		nombre := res.Map("nombre")
		correo := res.Map("correo")
		password := res.Map("password")
		for _, row := range rows {
			println(row.Int(id), row.Str(nombre), row.Str(correo), row.Str(password))
		}
		break
	case "pedido":
		id := res.Map("idpedido")
		direccion := res.Map("direccion")
		latlong := res.Map("latlong")
		fecha := res.Map("fecha")
		comprador := res.Map("comprador")
		for _, row := range rows {
			println(row.Int(id), row.Str(direccion), row.Str(latlong), row.Str(fecha), row.Str(comprador))
		}
		break
	case "articulo":
		id := res.Map("idarticulo")
		nombre := res.Map("nombre")
		imagen := res.Map("imagen")
		precio := res.Map("precio")
		descripcion := res.Map("descripcion")
		for _, row := range rows {
			println(row.Int(id), row.Str(nombre), row.Str(imagen), row.Str(precio), row.Str(descripcion))
		}
		break
	}
	printOK()
	db.Close()
	return rows, res

}

func QueryInsert(tabla string, param1 string, param2 string, param3 string, param4 string) {
	user := "root"
	pass := "root"
	dbname := "coppel"
	proto := "tcp"
	addr := "localhost:3307"
	db := mysql.New(proto, "", addr, user, pass, dbname)
	checkError(db.Connect())
	var query string
	switch tabla {
	case "usuario":
		query = ("INSERT INTO usuario (nombre,correo,password) values (?,?,?)")
		break
	case "pedido":
		query = ("INSERT INTO pedido (direccion,latlong,fecha,comprador) values (?,?,?,?)")
		break
	case "articulo":
		query = ("INSERT INTO articulo (nombre,imagen,precio,descripcion) values(?,?,?,?)")
		break
	}
	println(query)
	insForm, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	if strings.Contains(tabla, "usuario") {
		println("si entre a usuario")
		checkedResult(insForm.Exec(param1, param2, param3))
	} else {
		checkedResult(insForm.Exec(param1, param2, param3, param4))
	}
	printOK()
	db.Close()

}

func QueryDelete(tabla string, param string) {
	user := "root"
	pass := "root"
	dbname := "coppel"
	proto := "tcp"
	addr := "localhost:3307"

	db := mysql.New(proto, "", addr, user, pass, dbname)
	checkError(db.Connect())
	var query string
	switch tabla {
	case "usuario":
		query = "DELETE FROM usuario WHERE idusuario=?"
		break
	case "pedido":
		query = "DELETE FROM pedido WHERE idpedido=?"
		break
	case "articulo":
		query = "DELETE FROM articulo WHERE idarticulo=?"
		break
	}
	delForm, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	id, erro := strconv.Atoi(param)
	if erro != nil {
		panic(err.Error())
	}
	checkedResult(delForm.Exec(id))
	printOK()
	db.Close()

}

func QueryUpdate(tabla string, campoEditar string, param1 string, param2 string, param3 string, param4 string, param5 string) {
	user := "root"
	pass := "root"
	dbname := "coppel"
	proto := "tcp"
	addr := "localhost:3307"
	db := mysql.New(proto, "", addr, user, pass, dbname)
	checkError(db.Connect())
	var query string
	switch tabla {
	case "usuario":
		switch campoEditar {
		case "nombre":
			query = ("UPDATE usuario SET nombre=? WHERE idusuario=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "correo":
			query = ("UPDATE usuario SET correo=? WHERE idusuario=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "password":
			query = ("UPDATE usuario SET password=? WHERE idusuario=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		default:
			query = ("UPDATE usuario SET nombre=?, correo=?, password=? WHERE idusuario=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param4)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, param2, param3, id))
			break
		}

		break
	case "pedido":
		switch campoEditar {
		case "direccion":
			query = ("UPDATE pedido SET direccion=? WHERE idpedido=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "latlong":
			query = ("UPDATE pedido SET latlong=? WHERE idpedido=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "fecha":
			query = ("UPDATE pedido SET fecha=? WHERE idpedido=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "comprador":
			query = ("UPDATE pedido SET comprador=? WHERE idpedido=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		default:
			query = ("UPDATE pedido SET direccion=?, latlong=?, fecha=?, comprador=? WHERE idpedido=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param5)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, param2, param3, param4, id))
			break
		}
		break
	case "articulo":
		switch campoEditar {
		case "nombre":
			query = ("UPDATE articulo SET nombre=? WHERE idarticulo=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "imagen":
			query = ("UPDATE articulo SET imagen=? WHERE idarticulo=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "precio":
			query = ("UPDATE articulo SET precio=? WHERE idarticulo=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		case "descripcion":
			query = ("UPDATE articulo SET descripcion=? WHERE idarticulo=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param2)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, id))
			break
		default:
			query = ("UPDATE articulo SET nombre=?, imagen=?, precio=?, descripcion=? WHERE idarticulo=?")
			insForm, err := db.Prepare(query)
			if err != nil {
				panic(err.Error())
			}
			id, erro := strconv.Atoi(param5)
			if erro != nil {
				panic(err.Error())
			}
			checkedResult(insForm.Exec(param1, param2, param3, param4, id))
			break
		}
		break
	}
	printOK()
	db.Close()

}
