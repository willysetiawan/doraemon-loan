package models

type (
	ResGetInstallment struct {
		Id           string  `example:"2020102900000000000001" json:"id"`
		Name         string  `example:"1 Bulan" json:"name"`
		InterestRate float32 `example:"2020102900000000000001" json:"interestRate"`
	}
)
