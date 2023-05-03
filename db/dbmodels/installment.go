package dbmodels

import "time"

type (
	Installment struct {
		InstallmentKey       string     `gorm:"column:installmentKey"`
		InstallmentValue     int        `gorm:"column:installmentValue"`
		InstallmentType      string     `gorm:"column:installmentType"`
		InstallmentInterest  float32    `gorm:"column:installmentInterestRate"`
		InstallmentActive    bool       `gorm:"column:installmentActive"`
		InstallmentCreatedAt time.Time  `gorm:"column:installmentCreatedAt"`
		InstallmentCreatedBy string     `gorm:"column:installmentCreatedBy"`
		InstallmentUpdatedAt *time.Time `gorm:"column:installmentUpdatedAt"`
		InstallmentUpdatedBy string     `gorm:"column:installmentUpdatedBy"`
	}
)

func (Installment) TableName() string {
	return "m_installment"
}
