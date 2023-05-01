package dbmodels

import "time"

type (
	Partner struct {
		Id               string     `gorm:"column:id"`
		PartnerName      string     `gorm:"column:partnerName"`
		PartnerActive    bool       `gorm:"column:partnerActive"`
		PartnerCreatedAt time.Time  `gorm:"column:partnerCreatedAt"`
		PartnerCreatedBy string     `gorm:"column:partnerCreatedBy"`
		PartnerUpdatedAt *time.Time `gorm:"column:partnerUpdatedAt"`
		PartnerUpdatedBy string     `gorm:"column:partnerUpdatedBy"`
	}
)

func (Partner) TableName() string {
	return "m_partner"
}
