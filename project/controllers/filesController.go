package controller

import (
	"fmt"
	"net/http"
	"saveDatabase/project/models"
	"saveDatabase/project/services"

	"github.com/gin-gonic/gin"
)

func FilesFindAll(c *gin.Context) {
	fmt.Println("Bateu na controller")
	files, err := services.FilesFindAll()
	fmt.Println("Saiu da Service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar os arquivos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Arquivos": files})
}

func FileFindByPhone(c *gin.Context) {
	phone := c.Param("phone")

	files, err := services.FileFindByPhone(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Arquivos não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Arquivos": files})
}

func FileFindByDate(c *gin.Context) {
	date := c.Param("date")

	files, err := services.FileFindByDate(date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Arquivos": files})
}

func FileFindByGroupAndDate(c *gin.Context) {
	group := c.Param("group")
	date := c.Param("date")

	files, err := services.FileFindByGroupAndDate(group, date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"group": group, "date": date})
	c.JSON(http.StatusOK, gin.H{"Arquivos": files})
}

func SaveFiles(c *gin.Context) {

	var newFile models.Geral
	if err := c.ShouldBindJSON(&newFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados do arquivo inválidos"})
		return
	}

	persistedFiles, err := services.SaveFiles(&newFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar os arquivos. " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Arquivos salvos com sucesso", "Arquivo": persistedFiles})
}

func FilesRoutes(r *gin.Engine) {
	r.GET("/api/files", FilesFindAll)
	r.GET("/api/files/:phone", FileFindByPhone)
	r.GET("/api/date/:date", FileFindByDate)
	r.GET("/api/group/:group/:date", FileFindByGroupAndDate)
	r.POST("/api/savefiles", SaveFiles)
}
