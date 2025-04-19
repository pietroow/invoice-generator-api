package invoice

import (
	"time"

	"github.com/pietroow/invoice-generator-api/internal/company"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	InvoiceNumber int
	Date          time.Time

	CompanyID uint // your company (is_self = true)
	ClientID  uint // client (is_self = false)

	BillingType   string
	HourlyRate    float64 `gorm:"type:numeric(10,2)"`
	WorkedHours   float64 `gorm:"type:numeric(10,2)"`
	OverrideTotal bool

	Subtotal   float64 `gorm:"type:numeric(10,2)"`
	TaxPercent float64 `gorm:"type:numeric(5,2)"`
	TaxAmount  float64 `gorm:"type:numeric(10,2)"`
	Total      float64 `gorm:"type:numeric(10,2)"`

	Company company.Company `gorm:"foreignKey:CompanyID"`
	Client  company.Company `gorm:"foreignKey:ClientID"`
}

type InvoiceItem struct {
	gorm.Model

	InvoiceID   uint
	Description string
	Amount      float64 `gorm:"type:numeric(10,2)"`

	Invoice Invoice `gorm:"foreignKey:InvoiceID"`
}
