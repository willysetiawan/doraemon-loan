package base

import (
	"net/http"
	"process-loan/db"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/response"
	"strconv"
)

type InstallmentServiceInterface interface {
	GetInstallment(res *response.Response)
}

type installmentService struct {
	Name    string
	TraceID string
}

func InitInstallmentService(logName, traceID string) InstallmentServiceInterface {
	name := logName + "InstallmentService - "
	return &installmentService{
		Name:    name,
		TraceID: traceID,
	}
}

func (domain *installmentService) GetInstallment(res *response.Response) {
	resInstallment, errGetInstallment := db.GetActiveInstallment(domain.TraceID)
	if errGetInstallment != nil {
		res.Meta.DebugParam = "Failed get installment"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	var resData []models.ResGetInstallment
	if len(resInstallment) > 0 {
		for _, element := range resInstallment {
			resData = append(resData, models.ResGetInstallment{
				Id:           element.InstallmentKey,
				Name:         strconv.Itoa(element.InstallmentValue) + " " + element.InstallmentType,
				InterestRate: element.InstallmentInterest,
			})
		}
		res.Data = resData
		res.ResponseCode = strconv.Itoa(http.StatusOK)
		res.ResponseMessage = constant.MESSAGE_SUCCESS
	} else {
		res.ResponseCode = strconv.Itoa(http.StatusNotFound)
		res.ResponseMessage = constant.MESSAGE_NOT_FOUND
		res.Data = make([]interface{}, 0)
	}
	return

}
