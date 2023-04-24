package dbmodels

import "time"

type (
	Company struct {
		Id               string     `gorm:"column:id"`
		CompanyName      string     `gorm:"column:companyName"`
		CompanyActive    bool       `gorm:"column:companyActive"`
		CompanyCreatedAt time.Time  `gorm:"column:companyCreatedAt"`
		CompanyCreatedBy string     `gorm:"column:companyCreatedBy"`
		CompanyUpdatedAt *time.Time `gorm:"column:companyUpdatedAt"`
		CompanyUpdatedBy string     `gorm:"column:companyUpdatedBy"`
	}
)

func (Company) TableName() string {
	return "m_company"
}
