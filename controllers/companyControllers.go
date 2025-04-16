package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pietroow/invoice-generator-api/initializers"
	"github.com/pietroow/invoice-generator-api/models"
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

	companyCreate := models.Company{
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

	initializers.DB.Create(&companyCreate)

	fmt.Println(companyCreate)

	//create a post

	// return it
	c.JSON(200, gin.H{
		"message": "ponggg3x",
	})
}
