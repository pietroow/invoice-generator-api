package main

import (
	"github.com/pietroow/invoice-generator-api/internal/company"
	"github.com/pietroow/invoice-generator-api/internal/config"
	"github.com/pietroow/invoice-generator-api/internal/invoice"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDatabase()
}

func main() {
	config.DB.AutoMigrate(&company.Company{})
	config.DB.AutoMigrate(&invoice.Invoice{})
	config.DB.AutoMigrate(&invoice.InvoiceItem{})
}
