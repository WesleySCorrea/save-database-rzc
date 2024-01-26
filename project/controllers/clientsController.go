package controller

import (
	"fmt"
	"net/http"
	"saveDatabase/project/models"
	"saveDatabase/project/services"

	"github.com/gin-gonic/gin"
)

func ClientFindAll(c *gin.Context) {
	fmt.Println("Bateu na controller")
	clients, err := services.ClientFindAll()
	fmt.Println("Saiu da Service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar os clientes: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Clientes": clients})
}

func ClientFindByID(c *gin.Context) {
	id := c.Param("id")

	client, err := services.ClientFindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Cliente": client})
}

func ClientFindByIDAndGroup(c *gin.Context) {
	id := c.Param("id")
	group := c.Param("group")

	files, err := services.ClientFindByIDAndGroup(id, group)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Arquivos": files})
}

func ClientFindByIDAndGroupWithDate(c *gin.Context) {
	id := c.Param("id")
	group := c.Param("group")
	date := c.Param("date")

	files, err := services.ClientFindByIDAndGroupWithDate(id, group, date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Arquivos": files})
}

func SaveClient(c *gin.Context) {

	var newClient models.Client
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	persistedClient, err := services.SaveClient(&newClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o cliente. " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Cliente salvos com sucesso", "Cliente": persistedClient})
}

func UpdateClient(c *gin.Context) {
	id := c.Param("id")

	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	clientUpdate, updateErr := services.UpdateClient(id, &client)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente atualizado com sucesso", "cliente": clientUpdate})
}

func ClientRoutes(r *gin.Engine) {
	r.GET("/api/clients", ClientFindAll)
	r.GET("/api/clients/:id", ClientFindByID)
	r.GET("/api/clients/:id/:group", ClientFindByIDAndGroup)
	r.GET("/api/clients/:id/:group/:date", ClientFindByIDAndGroupWithDate)
	r.POST("/api/saveclients", SaveClient)
	r.PUT("/api/updateclients/:id", UpdateClient)
}
