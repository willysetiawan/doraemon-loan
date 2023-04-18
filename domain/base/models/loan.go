package models

type (
	ReqBookingLoan struct {
		CustomerName        string  `example:"2020102900000000000001" json:"partnerReferenceNo" validate:"required"`
		CustomerIdentityNo  string  `example:"2020102900000000000001" json:"beneficiaryAccountName" validate:"required,max=100" maxLength:"100"`
		PhoneNumber         string  `example:"2020102900000000000001" json:"beneficiaryAccountNo" validate:"required,max=34" maxLength:"34"`
		Email               string  `example:"Palembang" json:"beneficiaryAddress,omitempty" validate:"max=100" maxLength:"100"`
		CompanyName         string  `example:"2020102900000000000001" json:"beneficiaryBankCode" validate:"required,max=8" maxLength:"8"`
		EmployeeId          string  `example:"2020102900000000000001" json:"beneficiaryBankName,omitempty" validate:"max=50" maxLength:"50"`
		InstallmentId       int8    `example:"2020102900000000000001" json:"installmentId" validate:"required,max=50" maxLength:"50"`
		Amount              float64 `example:"2020102900000000000001" json:"currency,omitempty" validate:"max=3" maxLength:"3"`
		TaxIdentityImage    string  `example:"2020102900000000000001" json:"sourceAccountNo" validate:"required,max=19" maxLength:"19"`
		AgreeTermsCondition bool    `example:"2020102900000000000001" json:"transactionDate" validate:"required,max=25" maxLength:"25"`
	}

	ResBookingLoan struct {
		BookingNo string `example:"2020102900000000000001" json:"bookingNo"`
	}
)
