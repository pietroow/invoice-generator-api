package main

import (
	"github.com/pietroow/invoice-generator-api/initializers"
	"github.com/pietroow/invoice-generator-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Invoice{})
	initializers.DB.AutoMigrate(&models.Company{})
	initializers.DB.AutoMigrate(&models.InvoiceItem{})
}
