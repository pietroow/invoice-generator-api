package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pietroow/invoice-generator-api/controllers"
	"github.com/pietroow/invoice-generator-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/companies", controllers.CompaniesCreate)
	router.GET("/companies", controllers.CompaniesFind)
	router.GET("/companies/:id", controllers.CompaniesFindOne)
	router.Run()
}
