package base

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"process-loan/db"
	"process-loan/db/dbmodels"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/response"
	"strconv"
	"strings"
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

	randNumber, _ := rand.Prime(rand.Reader, 32)
	referenceNo := time.Now().Format("20060102") + randNumber.String()
	coI := strings.Index(string(req.TaxIdentityImage), ",")
	var extFile string
	switch strings.TrimSuffix(req.TaxIdentityImage[5:coI], ";base64") {
	case "image/png":
		extFile = ".png"
	case "image/jpeg":
		extFile = ".jpg"
	}

	identityImage, _ := base64.StdEncoding.DecodeString(strings.Replace(req.TaxIdentityImage[coI:], ",", "", 1))

	fileData := identityImage

	//Check Whitelist
	reqCheckWhitelist := models.ReqCheckWhitelist{
		CustomerIdentityNo: req.CustomerIdentityNo,
		PhoneNumber:        req.PhoneNumber,
		EmployeeId:         req.EmployeeId,
	}
	resDataWhitelist, errCheckWhitelist := db.CheckWhitelist(domain.TraceID, reqCheckWhitelist)
	if errCheckWhitelist != nil {
		res.Meta.DebugParam = "Failed process check whitelist"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	if resDataWhitelist.EmployeeId == "" {
		res.ResponseMessage = constant.MESSAGE_UNAUTHORIZED
		res.ResponseCode = strconv.Itoa(http.StatusUnauthorized)
		return
	}

	renameFile := fmt.Sprintf("%s%s", "KTA"+referenceNo, extFile)
	if _, err := os.Stat("storage/"); os.IsNotExist(err) {
		errDir := os.Mkdir("storage/", 0777)
		if errDir != nil {
			res.Meta.DebugParam = "Failed create path image storage"
			res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
			res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
			return
		}
	}
	permissions := 0777 // or whatever you need
	err := os.WriteFile("storage/"+renameFile, fileData, fs.FileMode(permissions))
	if err != nil {

	}

	//Booking Loan
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
		Status:               0,
		TaxIdentityImagePath: "storage/" + renameFile,
		AgreeTermsCondition:  *req.AgreeTermsCondition,
		BookingCreatedAt:     time.Time.Local(time.Now().UTC()),
	}

	_, errInsertBooking := db.InsertBookingLoan(domain.TraceID, reqBookingLoan)
	if errInsertBooking != nil {
		res.Meta.DebugParam = "Failed process booking"
		res.ResponseMessage = constant.MESSAGE_BACKEND_SYSTEM_FAILURE
		res.ResponseCode = strconv.Itoa(http.StatusInternalServerError)
		return
	}

	resData := models.ResBookingLoan{
		BookingNo:   "KTA" + referenceNo,
		BookingTime: reqBookingLoan.BookingCreatedAt.Format("2006-01-02T15:04:05+07:00"),
	}

	res.ResponseCode = strconv.Itoa(http.StatusCreated)
	res.ResponseMessage = constant.MESSAGE_CREATED
	res.Data = resData
	return

}
