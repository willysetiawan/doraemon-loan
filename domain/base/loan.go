package base

import (
	"crypto/rand"
	"net/http"
	"process-loan/db"
	"process-loan/db/dbmodels"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/response"
	"strconv"
	"time"
)

type LoanServiceInterface interface {
	BookingLoan(req models.ReqBookingLoan, res *response.Response)
}

type loanService struct {
	Name    string
	TraceID string
}

func InitLoanService(logName, traceID string) LoanServiceInterface {
	name := logName + "LoanService - "
	return &loanService{
		Name:    name,
		TraceID: traceID,
	}
}

func (domain *loanService) BookingLoan(req models.ReqBookingLoan, res *response.Response) {
	//Check Duplicate Partner Reference No
	randNumber, _ := rand.Prime(rand.Reader, 64)
	referenceNo := time.Now().Format("20060102") + randNumber.String()
	//Debit Account No Digital
	reqBookingLoan := dbmodels.BookingLoan{
		BookingId:            "KTA" + referenceNo,
		CustomerName:         req.CustomerName,
		CustomerIdentityNo:   req.CustomerIdentityNo,
		PhoneNumber:          req.PhoneNumber,
		Email:                req.Email,
		CompanyName:          req.CompanyName,
		EmployeeId:           req.EmployeeId,
		InstallmentId:        req.InstallmentId,
		Amount:               req.Amount,
		TaxIdentityImagePath: req.TaxIdentityImage,
		AgreeTermsCondition:  req.AgreeTermsCondition,
		BookingCreatedAt:     time.Time.Local(time.Now().UTC()),
	}

	_, errInsertBooking := db.InsertBookingLoan(domain.TraceID, reqBookingLoan)
	if errInsertBooking != nil {
		res.Meta.DebugParam = "Failed process interbank transfer"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	resData := models.ResBookingLoan{
		BookingNo: "KTA" + referenceNo,
	}

	res.ResponseCode = strconv.Itoa(http.StatusOK)
	res.ResponseMessage = constant.MESSAGE_SUCCESS
	res.Data = resData
	return

}
