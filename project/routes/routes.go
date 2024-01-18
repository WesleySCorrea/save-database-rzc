package routes

import (
	"net/http"
	controller "saveDatabase/project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", homeHandler)
	controller.FilesRoutes(r)
	controller.ClientRoutes(r)
}

func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Projeto RZC - OKAY")
}
