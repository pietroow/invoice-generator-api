package main

import (
	"github.com/gin-gonic/gin"
	"invoice-generator/controllers"
	"invoice-generator/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/invoices", controllers.InvoicesCreate)
	router.Run()
}
