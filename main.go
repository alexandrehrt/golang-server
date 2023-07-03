package main

import (
	"go-crud/initializers"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
  router := gin.Default()

  routes.SetupRoutes(router)

  router.Run()
}