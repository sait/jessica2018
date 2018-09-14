package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cliente struct {
	ID       string `db:"id"       json:"id"`
	Nombre   string `db:"nombre"   json:"nombre"`
	Apellido string `db:"apellido" json:"apellido"`
}

var DB *sqlx.DB

func OpenDB() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/pruebasgo")
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ListClientes() (clientes []Cliente, err error) {
	err = DB.Select(&clientes, `SELECT id, nombre, apellido FROM clientes`)
	return
}

func GetCliente(clientID string) (cliente Cliente, err error) {
	err = DB.Get(&cliente, `SELECT id, nombre, apellido FROM clientes
		WHERE id=?`, clientID)
	return
}

func InsertCliente(cliente Cliente) (err error) {
	_, err = DB.NamedExec(`INSERT INTO clientes (nombre, apellido)
		VALUES(:nombre, :apellido)`, &cliente)
	return
}

func UpdateCliente(cliente Cliente) (err error) {
	_, err = DB.NamedExec(`UPDATE clientes
		SET nombre=:nombre, apellido=:apellido
		WHERE id=:id`, cliente)
	return
}

func DeleteCliente(clienteID string) (err error) {
	_, err = DB.Exec(`DELETE FROM clientes
		WHERE id=?`, clienteID)
	return
}

