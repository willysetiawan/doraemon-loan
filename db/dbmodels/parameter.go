package dbmodels

import "time"

type (
	Installment struct {
		Key       string     `gorm:"column:key"`
		Value     int        `gorm:"column:value"`
		Type      string     `gorm:"column:type"`
		Interest  float32    `gorm:"column:interestRate"`
		Active    bool       `gorm:"column:active"`
		CreatedAt time.Time  `gorm:"column:createdAt"`
		CreatedBy string     `gorm:"column:createdBy"`
		UpdatedAt *time.Time `gorm:"column:updatedAt"`
		UpdatedBy string     `gorm:"column:updatedBy"`
	}
)

func (Installment) TableName() string {
	return "m_installment"
}
