package main

import (
	"invoice-generator/initializers"
	"invoice-generator/models"
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
