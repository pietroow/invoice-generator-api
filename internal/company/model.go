package company

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name         string
	Email        string
	PhoneNumber  string
	AddressLine1 string
	City         string
	State        string
	PostalCode   string
	Country      string

	IsSelf      bool
	ContractEnd *time.Time

	BillingType        string  // "FIXED" or "HOURLY"
	HourlyRate         float64 `gorm:"type:numeric(10,2)"`
	DefaultMonthlyRate float64 `gorm:"type:numeric(10,2)"`
	InvoiceCounter     int
}
