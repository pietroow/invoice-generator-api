package main

import (
	"time"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.POST("/companies", company.CompaniesCreate)
	router.GET("/companies", company.CompaniesFind)
	router.GET("/companies/:id", company.CompaniesFindOne)
	router.PUT("/companies/:id", company.CompanyUpdate)
	router.DELETE("/companies/:id", company.CompanyDelete)
	router.Run()
}
