package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"process-loan/domain/base"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/env"
	"process-loan/lib/log"
	"process-loan/lib/response"

	"github.com/gin-gonic/gin"
)

type loanController struct {
	Name    string
	BaseUrl string
}

func InitLoanController() *loanController {
	baseUrl := env.String("MainSetup.ServerHost", "")
	return &loanController{
		Name:    "LoanController - ",
		BaseUrl: baseUrl,
	}
}

// Example Send godoc
// @Summary Example Send
// @Description Example Deposit List Send
// @Tags Example
// @ID Example Send
// @Accept  json
// @Produce  json
// @Param body body models.ReqDepositList true "request body"
// @Success 200 {object} models.ResExampleSendSuccess
// @Failure 400 {object} models.ResExampleSendError
// @Router /v1/partner/check-data [post]
func (ctrl *loanController) BookingLoan(c *gin.Context) {
	logName := ctrl.Name + " Booking Loan "
	res := response.Response{
		ResponseCode:    constant.CODE_FAILED,
		ResponseMessage: http.StatusText(http.StatusBadRequest),
	}
	traceID := c.GetHeader("Trace-ID")
	// authorization := strings.Split(c.Request.Header.Get("Authorization"), " ")
	var req models.ReqBookingLoan
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/booking-loan", nil, "POST", &req, &res)
	log.LogFmtTemp("Start " + logName)
	defer log.LogFmtTemp("End " + logName)
	// res.Meta.TraceID = traceID
	// ini buat POST
	dataBody, _ := c.Get("body")
	if err := json.Unmarshal([]byte(fmt.Sprint(dataBody)), &req); err != nil {
		log.LogData(traceID, logName, "json.Unmarshal", constant.LEVEL_LOG_ERROR, err.Error())
		res.Meta.DebugParam = err.Error()
		response.JsonGen(c, res)
		return
	}

	// validate := validator.New()
	// //Validate Required Struct Field
	// if err := validate.Struct(req); err != nil {
	// 	var jsonField bytes.Buffer
	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		if err.ActualTag() == "required" {
	// 			arrayString := strings.Split(strings.Replace(err.StructNamespace(), "ReqBookingLoan.", "", -1), ".")
	// 			if len(arrayString) == 1 {
	// 				structRequired := strings.Replace(err.StructNamespace(), "ReqBookingLoan.", "", -1)
	// 				jsonField.WriteString(strings.ToLower(string([]rune(structRequired)[0])) + string([]rune(structRequired)[1:]))
	// 				jsonField.WriteString(" ")
	// 			} else {

	// 				for index, element := range arrayString {
	// 					if index == (len(arrayString) - 1) {
	// 						jsonField.WriteString(strings.ToLower(string([]rune(element)[0])) + string([]rune(element)[1:]))
	// 						jsonField.WriteString(" ")
	// 					} else {
	// 						jsonField.WriteString(strings.ToLower(string([]rune(element)[0])) + string([]rune(element)[1:]) + ".")
	// 					}
	// 				}

	// 			}
	// 		}
	// 	}

	// 	if jsonField.String() != "" {
	// 		res.ResponseMessage = constant.MESSAGE_MISSING_MANDATORY_FIELD + " " + jsonField.String()
	// 		res.ResponseCode = strconv.Itoa(http.StatusBadRequest)
	// 		c.JSON(http.StatusBadRequest, res)
	// 		return
	// 	}

	// }

	// //Validate Alpha Numeric Struct Field
	// if err := validate.Struct(req); err != nil {
	// 	var jsonField bytes.Buffer
	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		// fmt.Println(err.StructField())
	// 		if err.ActualTag() == "alpha" || err.ActualTag() == "alphaunicode" || err.ActualTag() == "numeric" || err.ActualTag() == "min" || err.ActualTag() == "max" {
	// 			// res.ResponseCode = strconv.Itoa(http.StatusBadRequest) + constant.SERVICE_CODE_INTERBANK_TRANSFER + constant.CASE_CODE_EXTERNAL_SERVER_ERROR
	// 			// res.ResponseMessage = constant.MESSAGE_INVALID_FIELD_FORMAT + " " + strings.ToLower(string([]rune(err.StructField())[0])) + string([]rune(err.StructField())[1:])
	// 			// c.JSON(http.StatusBadRequest, res)

	// 			// return
	// 			arrayString := strings.Split(strings.Replace(err.StructNamespace(), "ReqBookingLoan.", "", -1), ".")
	// 			if len(arrayString) == 1 {
	// 				structRequired := strings.Replace(err.StructNamespace(), "ReqBookingLoan.", "", -1)
	// 				jsonField.WriteString(strings.ToLower(string([]rune(structRequired)[0])) + string([]rune(structRequired)[1:]))
	// 				jsonField.WriteString(" ")
	// 			} else {

	// 				for index, element := range arrayString {
	// 					if index == (len(arrayString) - 1) {
	// 						jsonField.WriteString(strings.ToLower(string([]rune(element)[0])) + string([]rune(element)[1:]))
	// 						jsonField.WriteString(" ")
	// 					} else {
	// 						jsonField.WriteString(strings.ToLower(string([]rune(element)[0])) + string([]rune(element)[1:]) + ".")
	// 					}
	// 				}

	// 			}
	// 		}

	// 	}

	// 	if jsonField.String() != "" {
	// 		res.ResponseMessage = constant.MESSAGE_INVALID_FIELD_FORMAT + " " + jsonField.String()
	// 		res.ResponseCode = strconv.Itoa(http.StatusBadRequest)
	// 		c.JSON(http.StatusBadRequest, res)
	// 		return
	// 	}

	// }

	base.InitLoanService(logName, traceID).BookingLoan(req, &res)
	response.JsonGen(c, res)

}
