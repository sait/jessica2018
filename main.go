package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"sait.mx/DEMOSQLX/model"
	"sait.mx/DEMOSQLX/controllers"

)


func main() {
	fmt.Println("Pruebas con base de datos")
	model.OpenDB()

	r := gin.Default()
	r.GET("/api/v1/clientes", controllers.ListClientes)
	//este se pone en el postman mas el numero de puerto
	r.GET("/api/v1/clientes/:id", controllers.GetCliente)
	r.POST("/api/v1/clientes", controllers.InsertCliente)
	r.PUT("/api/v1/clientes/:id", controllers.UpdateCliente)
	r.DELETE("/api/v1/clientes/:id", controllers.DeleteCliente)

	r.Run(":9053")  //Nombre del puerto
}



