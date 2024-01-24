package controller

import (
	"net/http"
	"saveDatabase/project/models"
	"saveDatabase/project/services"

	"github.com/gin-gonic/gin"
)

func GroupFindByPhone(c *gin.Context) {
	phone := c.Param("phone")

	group, err := services.GroupFindByPhone(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grupo não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Grupo": group})
}

func GroupFindByID(c *gin.Context) {
	id := c.Param("id")

	group, err := services.GroupFindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Grupo": group})
}

func SaveGroup(c *gin.Context) {

	var newGroup models.Groups
	if err := c.ShouldBindJSON(&newGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	persistedGroup, err := services.SaveGroup(&newGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar um grupo. " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Grupo salvo com sucesso", "Grupo": persistedGroup})
}

func GroupsRoutes(r *gin.Engine) {
	r.GET("/api/groups/group/:phone", GroupFindByPhone)
	r.GET("/api/groups/id/:id", GroupFindByID)
	r.POST("/api/savegroups", SaveGroup)
}
