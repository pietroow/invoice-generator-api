package main

import (
	"github.com/pietroow/invoice-generator-api/company"
	"github.com/pietroow/invoice-generator-api/initializers"
	"github.com/pietroow/invoice-generator-api/invoice"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&company.Company{})
	initializers.DB.AutoMigrate(&invoice.Invoice{})
	initializers.DB.AutoMigrate(&invoice.InvoiceItem{})
}
