package routes

import (
	"net/http"

	controllers "github.com/Ahad-Parmar/Assignment_7/CRUD_GIN_POSTGRESQL/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/trucks", controllers.GetAllTrucks)
	router.POST("/truck", controllers.CreateTruck)
	router.GET("/truck/:truckid", controllers.GetSingleTruck)
	router.PUT("/truck/:truckid", controllers.EditTruck)
	router.DELETE("/truck/:truckid", controllers.DeleteTruck)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
