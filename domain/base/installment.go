package base

import (
	"net/http"
	"process-loan/db"
	"process-loan/db/dbmodels"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/response"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type InstallmentServiceInterface interface {
	GetInstallment(res *response.Response)
	InsertInstallment(req models.ReqInstallment, res *response.Response)
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
				Id:           element.Key,
				Name:         strconv.Itoa(element.Value) + " " + element.Type,
				InterestRate: element.Interest,
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

func (domain *installmentService) InsertInstallment(req models.ReqInstallment, res *response.Response) {
	//Installment Payload
	reqInstallment := dbmodels.Installment{
		Key:       uuid.New().String(),
		Value:     req.Value,
		Type:      req.Type,
		Interest:  req.InterestRate,
		Active:    *req.Active,
		CreatedAt: time.Time.Local(time.Now().UTC()),
		CreatedBy: "Hardcode API",
		UpdatedAt: &time.Time{},
	}

	//Check Conflict
	resCheckData, errCheckData := db.CheckInstallment(domain.TraceID, req)
	if errCheckData != nil {
		res.Meta.DebugParam = "Failed check data"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	if resCheckData.Key != "" {
		res.ResponseMessage = "Installment already exist"
		res.ResponseCode = strconv.Itoa(http.StatusConflict)
		return
	}

	_, errInsertInstallment := db.InsertInstallment(domain.TraceID, reqInstallment)
	if errInsertInstallment != nil {
		res.Meta.DebugParam = "Failed process installment"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	resData := models.ResInsertInstallment{
		Id:           reqInstallment.Key,
		Name:         strconv.Itoa(reqInstallment.Value) + " " + reqInstallment.Type,
		InterestRate: reqInstallment.Interest,
		CreatedAt:    reqInstallment.CreatedAt.Format("2006-01-02T15:04:05+07:00"),
	}

	res.ResponseCode = strconv.Itoa(http.StatusCreated)
	res.ResponseMessage = constant.MESSAGE_CREATED
	res.Data = resData
	return

}
