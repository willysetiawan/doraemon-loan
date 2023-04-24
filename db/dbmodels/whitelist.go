package dbmodels

import "time"

type (
	Whitelist struct {
		CompanyId string `gorm:"column:companyId"`
		// CompanyName             string     `gorm:"column:companyName"`
		CIF                     string     `gorm:"column:employeeCif"`
		EmployeeName            string     `gorm:"column:employeeName"`
		EmployeeId              string     `gorm:"column:employeeId"`
		EmployeeIdentityNo      string     `gorm:"column:employeeIdentityNo"`
		EmployeePhoneNumber     string     `gorm:"column:employeeMobilePhoneNo"`
		Email                   string     `gorm:"column:employeeEmail"`
		Salary                  float64    `gorm:"column:employeeSalary"`
		EmployeeStatus          string     `gorm:"column:employeeStatus"`
		MaxLoanAmount           float64    `gorm:"column:employeeMaxLoanAmount"`
		ParticipatePayrollMonth string     `gorm:"column:employeeParticipatePayrollMonth"`
		WhitelistCreatedAt      time.Time  `gorm:"column:whitelistCreatedAt"`
		WhitelistCreatedBy      string     `gorm:"column:whitelistCreatedBy"`
		WhitelistUpdatedAt      *time.Time `gorm:"column:whitelistUpdatedAt"`
		WhitelistUpdatedBy      string     `gorm:"column:whitelistUpdatedBy"`
	}
)

func (Whitelist) TableName() string {
	return "m_whitelist"
}
