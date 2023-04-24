package models

type (
	ReqWhitelist struct {
		CompanyId             string  `example:"2020102900000000000001" json:"companyId" validate:"required,numeric"`
		CIF                   string  `example:"2020102900000000000001" json:"cif" validate:"required,max=16,numeric" maxLength:"16"`
		EmployeeName          string  `example:"2020102900000000000001" json:"employeeName" validate:"required,max=15,numeric" maxLength:"15"`
		EmployeeId            string  `example:"Palembang" json:"employeeId" validate:"max=100,numeric" maxLength:"100"`
		EmployeeIdentityNo    string  `example:"Palembang" json:"employeeIdentityNo" validate:"max=100,numeric" maxLength:"100"`
		EmployeeMobilePhoneNo string  `example:"2020102900000000000001" json:"employeeMobilePhoneNo" validate:"required,max=50,numeric" maxLength:"50"`
		EmployeeEmail         string  `example:"2020102900000000000001" json:"employeeEmail" validate:"required,max=20,alphaunicode" maxLength:"20"`
		EmployeeSalary        float64 `example:"2020102900000000000001" json:"employeeSalary" validate:"required,max=16,numeric" maxLength:"16"`
		EmployeeStatus        string  `example:"2020102900000000000001" json:"employeeStatus" validate:"required"`
		PayrollMonth          string  `example:"2020102900000000000001" json:"participatePayrollMonth" validate:"required"`
		MaxLoan               float64 `example:"2020102900000000000001" json:"maxLoanAmount" validate:"required"`
	}

	ResGetWhitelist struct {
		Id                    string  `example:"2020102900000000000001" json:"id,omitempty"`
		Company               string  `example:"2020102900000000000001" json:"company,omitempty"`
		CIF                   string  `example:"2020102900000000000001" json:"cif,omitempty"`
		EmployeeName          string  `example:"2020102900000000000001" json:"employeeName,omitempty"`
		EmployeeId            string  `example:"2020102900000000000001" json:"employeeId,omitempty"`
		EmployeeIdentityNo    string  `example:"Palembang" json:"employeeIdentityNo,omitempty,omitempty"`
		EmployeeMobilePhoneNo string  `example:"2020102900000000000001" json:"employeeMobilePhoneNo,omitempty"`
		EmployeeEmail         string  `example:"2020102900000000000001" json:"employeeEmail,omitempty"`
		EmployeeSalary        float64 `example:"2020102900000000000001" json:"employeeSalary,omitempty"`
		EmployeeStatus        string  `example:"2020102900000000000001" json:"employeeStatus,omitempty"`
		PayrollMonth          string  `example:"2020102900000000000001" json:"payrollMonth,omitempty"`
		MaxLoan               float64 `example:"2020102900000000000001" json:"maxLoan,omitempty"`
	}
)
