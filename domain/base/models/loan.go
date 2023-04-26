package models

type (
	ReqCheckWhitelist struct {
		CustomerIdentityNo string `example:"2020102900000000000001" json:"customerIdentityNo" validate:"max=16,numeric" maxLength:"16"`
		PhoneNumber        string `example:"2020102900000000000001" json:"phoneNumber" validate:"max=15,numeric" maxLength:"15"`
		EmployeeId         string `example:"2020102900000000000001" json:"employeeId" validate:"max=20,alphaunicode" maxLength:"20"`
	}
	ReqBookingLoan struct {
		CustomerName        string  `example:"2020102900000000000001" json:"customerName" validate:"required,max=100,alphaunicode" maxLength:"100"`
		CustomerIdentityNo  string  `example:"2020102900000000000001" json:"customerIdentityNo" validate:"required,max=16,numeric" maxLength:"16"`
		PhoneNumber         string  `example:"2020102900000000000001" json:"phoneNumber" validate:"required,max=15,numeric" maxLength:"15"`
		Email               string  `example:"Palembang" json:"email,omitempty" validate:"max=100" maxLength:"100"`
		CompanyName         string  `example:"2020102900000000000001" json:"companyName" validate:"required,max=50,alphaunicode" maxLength:"50"`
		EmployeeId          string  `example:"2020102900000000000001" json:"employeeId" validate:"required,max=20,alphaunicode" maxLength:"20"`
		InstallmentId       int     `example:"2020102900000000000001" json:"installmentId" validate:"required,max=2,numeric" maxLength:"2"`
		Amount              float64 `example:"2020102900000000000001" json:"amount" validate:"required,max=16,numeric" maxLength:"16"`
		TaxIdentityImage    string  `example:"2020102900000000000001" json:"taxIdentityImage" validate:"required"`
		AgreeTermsCondition *bool   `example:"2020102900000000000001" json:"agreeTermsCondition" validate:"required"`
	}

	ResBookingLoan struct {
		BookingNo   string `example:"2020102900000000000001" json:"bookingNo"`
		BookingTime string `example:"2020102900000000000001" json:"bookingTime"`
	}

	ReqProcessBookingLoan struct {
		BookingId string `example:"2020102900000000000001" json:"bookingId" validate:"required"`
		Action    *bool  `example:"2020102900000000000001" json:"action" validate:"required"`
	}

	ResProcessBookingLoan struct {
		BookingNo string `example:"2020102900000000000001" json:"bookingNo"`
		Status    string `example:"2020102900000000000001" json:"status"`
	}

	ResGetListBookingLoan struct {
		BookingId    string  `example:"KTA2002043423097" json:"bookingId"`
		CustomerName string  `example:"Budi" json:"customerName"`
		PhoneNumber  string  `example:"08527427327" json:"phoneNumber"`
		CompanyName  string  `example:"PT. AAAA" json:"companyName"`
		Installment  string  `example:"5 Bulan" json:"installment"`
		Amount       float64 `example:"50000000" json:"amount"`
		Status       string  `example:"Approved" json:"status"`
		CreatedAt    string  `example:"2023-04-18 17:14:31.256179" json:"createdAt"`
	}
)
