package dbmodels

import "time"

type (
	BookingLoan struct {
		BookingId            string    `gorm:"column:bookingId"`
		CustomerName         string    `gorm:"column:customerName"`
		CustomerIdentityNo   string    `gorm:"column:customerIdentityNo"`
		PhoneNumber          string    `gorm:"column:phoneNumber"`
		Email                string    `gorm:"column:email"`
		CompanyName          string    `gorm:"column:companyName"`
		EmployeeId           string    `gorm:"column:employeeId"`
		InstallmentId        int       `gorm:"column:installmentId"`
		Amount               float64   `gorm:"column:amount"`
		Status               int       `gorm:"column:status"`
		TaxIdentityImagePath string    `gorm:"column:taxIdentityImagePath"`
		AgreeTermsCondition  bool      `gorm:"column:agreeTermsCondition"`
		BookingCreatedAt     time.Time `gorm:"column:bookingCreatedAt"`
	}
)

func (BookingLoan) TableName() string {
	return "t_booking_loan"
}
