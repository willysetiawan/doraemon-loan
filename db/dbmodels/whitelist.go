package dbmodels

import "time"

type (
	Whitelist struct {
		CompanyId               string     `gorm:"column:companyId"`
		CompanyName             string     `gorm:"column:companyName"`
		CIF                     string     `gorm:"column:cif"`
		CustomerName            string     `gorm:"column:customerName"`
		EmployeeId              string     `gorm:"column:employeeId"`
		CustomerIdentityNo      string     `gorm:"column:customerIdentityNo"`
		CustomerPhoneNumber     string     `gorm:"column:customerPhoneNumber"`
		Email                   string     `gorm:"column:email"`
		Salary                  string     `gorm:"column:salary"`
		EmployeeStatus          string     `gorm:"column:employeeStatus"`
		MaxLoanAmount           string     `gorm:"column:maxLoanAmount"`
		ParticipatePayrollMonth string     `gorm:"column:participatePayrollMonth"`
		WhitelistCreatedAt      time.Time  `gorm:"column:whitelistCreatedAt"`
		WhitelistCreatedBy      string     `gorm:"column:whitelistCreatedBy"`
		WhitelistUpdatedAt      *time.Time `gorm:"column:whitelistUpdatedAt"`
		WhitelistUpdatedBy      string     `gorm:"column:whitelistUpdatedBy"`
	}
)

func (Whitelist) TableName() string {
	return "m_whitelist"
}
