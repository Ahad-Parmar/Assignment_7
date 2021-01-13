package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "github.com/Ahad-Parmar/Assignment_7/CRUD_GIN_POSTGRESQL/config"
	routes "github.com/Ahad-Parmar/Assignment_7/CRUD_GIN_POSTGRESQL/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
