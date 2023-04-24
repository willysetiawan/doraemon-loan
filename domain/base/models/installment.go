package models

type (
	ReqInstallment struct {
		Type         string  `example:"2020102900000000000001" json:"type" validate:"required,max=100,alphaunicode" maxLength:"100"`
		Value        int     `example:"2020102900000000000001" json:"value" validate:"required,max=16,numeric" maxLength:"16"`
		InterestRate float32 `example:"2020102900000000000001" json:"interestRate" validate:"required,max=15,numeric" maxLength:"15"`
		Active       *bool   `example:"2020102900000000000001" json:"active" validate:"required"`
	}

	ResGetInstallment struct {
		Id           string  `example:"2020102900000000000001" json:"id"`
		Name         string  `example:"1 Bulan" json:"name"`
		InterestRate float32 `example:"2020102900000000000001" json:"interestRate"`
	}

	ResInsertInstallment struct {
		Id           string  `example:"2020102900000000000001" json:"id"`
		Name         string  `example:"2020102900000000000001" json:"name"`
		InterestRate float32 `example:"2020102900000000000001" json:"interestRate"`
		CreatedAt    string  `example:"2020102900000000000001" json:"createdAt"`
	}
)
