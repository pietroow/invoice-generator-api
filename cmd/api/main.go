package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pietroow/invoice-generator-api/internal/company"
	"github.com/pietroow/invoice-generator-api/internal/config"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/companies", company.CompaniesCreate)
	router.GET("/companies", company.CompaniesFind)
	router.GET("/companies/:id", company.CompaniesFindOne)
	router.PUT("/companies/:id", company.CompanyUpdate)
	router.DELETE("/companies/:id", company.CompanyDelete)
	router.Run()
}
