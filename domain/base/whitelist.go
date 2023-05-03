package base

import (
	"net/http"
	"process-loan/db"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/response"
	"strconv"
)

type WhitelistServiceInterface interface {
	GetWhitelist(res *response.Response)
	CheckWhitelist(req models.ReqCheckWhitelist, res *response.Response)
}

type whitelistService struct {
	Name    string
	TraceID string
}

func InitWhitelistService(logName, traceID string) WhitelistServiceInterface {
	name := logName + "WhitelistService - "
	return &whitelistService{
		Name:    name,
		TraceID: traceID,
	}
}

func (domain *whitelistService) CheckWhitelist(req models.ReqCheckWhitelist, res *response.Response) {
	resData, errCheckWhitelist := db.CheckWhitelist(domain.TraceID, req)
	if errCheckWhitelist != nil {
		res.Meta.DebugParam = "Failed process check whitelist"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	if resData.EmployeeId == "" {
		res.ResponseMessage = constant.MESSAGE_UNAUTHORIZED
		res.ResponseCode = strconv.Itoa(http.StatusUnauthorized)
		return
	} else {
		res.ResponseMessage = constant.MESSAGE_SUCCESS
		res.ResponseCode = strconv.Itoa(http.StatusOK)
		return
	}
}

func (domain *whitelistService) GetWhitelist(res *response.Response) {
	resWhitelist, errGetWhitelist := db.GetWhitelist(domain.TraceID)
	if errGetWhitelist != nil {
		res.Meta.DebugParam = "Failed get whitelist"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	var resData []models.ResGetWhitelist
	if len(resWhitelist) > 0 {
		for _, element := range resWhitelist {
			resData = append(resData, models.ResGetWhitelist{
				// Id:           element.I,
				// Name:         strconv.Itoa(element.Value) + " " + element.Type,
				Partner:               element.PartnerName,
				CIF:                   element.CIF,
				EmployeeName:          element.EmployeeName,
				EmployeeId:            element.EmployeeId,
				EmployeeIdentityNo:    element.EmployeeIdentityNo,
				EmployeeMobilePhoneNo: element.EmployeePhoneNumber,
				EmployeeEmail:         element.Email,
				EmployeeSalary:        element.Salary,
				EmployeeStatus:        element.EmployeeStatus,
				PayrollMonth:          element.ParticipatePayrollMonth,
				MaxLoan:               element.MaxLoanAmount,
			})
		}
		res.Data = resData
	} else {
		res.Data = make([]interface{}, 0)
	}

	res.ResponseCode = strconv.Itoa(http.StatusOK)
	res.ResponseMessage = constant.MESSAGE_SUCCESS

	return

}
