package company

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pietroow/invoice-generator-api/internal/config"
)

type CompanyCreate struct {
	Name               string  `form:"name" json:"name" binding:"required"`
	Email              string  `form:"email" json:"email" binding:"required"`
	PhoneNumber        string  `form:"phone_number" json:"phoneNumber" binding:"required"`
	AddressLine1       string  `form:"address_line_1" json:"addressLine1" binding:"required"`
	City               string  `form:"city" json:"city" binding:"required"`
	State              string  `form:"state" json:"state" binding:"required"`
	PostalCode         string  `form:"postal_code" json:"postalCode" binding:"required"`
	Country            string  `form:"country" json:"country" binding:"required"`
	IsSelf             bool    `json:"isSelf"`
	ContractEnd        string  `json:"contractEnd"`
	BillingType        string  `json:"billingType"`
	HourlyRate         float64 `json:"hourlyRate"`
	DefaultMonthlyRate float64 `json:"defaultMonthlyRate"`
	InvoiceCounter     int     `json:"invoiceCounter"`
}

func CompaniesCreate(c *gin.Context) {
	var body CompanyCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contractEndTime, err := time.Parse(time.RFC3339, body.ContractEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for contractEnd"})
		return
	}

	companyCreate := Company{
		Name:               body.Name,
		Email:              body.Email,
		PhoneNumber:        body.PhoneNumber,
		AddressLine1:       body.AddressLine1,
		City:               body.City,
		State:              body.State,
		PostalCode:         body.PostalCode,
		Country:            body.Country,
		IsSelf:             body.IsSelf,
		ContractEnd:        &contractEndTime,
		BillingType:        body.BillingType,
		HourlyRate:         body.HourlyRate,
		DefaultMonthlyRate: body.DefaultMonthlyRate,
		InvoiceCounter:     body.InvoiceCounter,
	}

	result := config.DB.Create(&companyCreate)
	if result != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"company": companyCreate,
	})
}

func CompaniesFind(c *gin.Context) {
	var companies []Company
	config.DB.Find(&companies)

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

func CompaniesFindOne(c *gin.Context) {
	var company Company
	id := c.Param("id")
	config.DB.Find(&company, id)
	c.JSON(200, gin.H{
		"company": company,
	})
}

func CompanyUpdate(c *gin.Context) {
	var body CompanyCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var company Company
	id := c.Param("id")
	config.DB.Find(&company, id)

	contractEndTime, err := time.Parse(time.RFC3339, body.ContractEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for contractEnd"})
		return
	}

	config.DB.Model(&company).Updates(Company{
		Name:               body.Name,
		Email:              body.Email,
		PhoneNumber:        body.PhoneNumber,
		AddressLine1:       body.AddressLine1,
		City:               body.City,
		State:              body.State,
		PostalCode:         body.PostalCode,
		Country:            body.Country,
		IsSelf:             body.IsSelf,
		ContractEnd:        &contractEndTime,
		BillingType:        body.BillingType,
		HourlyRate:         body.HourlyRate,
		DefaultMonthlyRate: body.DefaultMonthlyRate,
		InvoiceCounter:     body.InvoiceCounter,
	})

	c.JSON(204, gin.H{
		"company": company,
	})

}

func CompanyDelete(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&Company{}, id)
	c.Status(200)
}
