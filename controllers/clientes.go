package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"sait.mx/DEMOSQLX/model"
)

func ListClientes(c *gin.Context) {
	clientes, err := model.ListClientes() //clientes listclientes
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, clientes)
}

func GetCliente(c *gin.Context)  {                                     //OBTENER CLIENTE
	id := c.Param("id")
	cliente, err := model.GetCliente(id)  //cliente getcliente
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, cliente)
}

func InsertCliente(c *gin.Context){									//INSERTA A CLIENTE
	body, err := ioutil.ReadAll(c.Request.Body)   //SIRVE PARA LEER EL BODY
	if err != nil{
		c.JSON(500, err.Error())
		return
	}
	var cliente model.Cliente
	err = json.Unmarshal(body, &cliente)
	if err != nil{
		c.JSON(500, err.Error())
		return
	}
	err = model.InsertCliente(cliente)
	if err != nil{
		c.JSON(500, err.Error())
		return

	}
	c.JSON(200, "Insertado correctamente")
}


func UpdateCliente(c *gin.Context){
	id := c.Param("id")//USA  PUT
	body, err := ioutil.ReadAll(c.Request.Body)   //SIRVE PARA LEER EL BODY
	if err != nil{
		c.JSON(500, err.Error())
		return
	}
	var cliente model.Cliente
	err = json.Unmarshal(body, &cliente)
	if err != nil{
		c.JSON(500, err.Error())
		return
	}
	cliente.ID = id
	err = model.UpdateCliente(cliente)
	if err != nil{
		c.JSON(500, err.Error())
		return

	}
	c.JSON(200, "Actualizado correctamente")
}


func DeleteCliente(c *gin.Context)  {  //eliminar CLIENTE
	id := c.Param("id")
	err := model.DeleteCliente(id)  //cliente getcliente
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Eliminado correctamente")
}