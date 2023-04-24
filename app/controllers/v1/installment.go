package v1

import (
	"encoding/json"
	"process-loan/domain/base"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/env"
	"process-loan/lib/log"
	"process-loan/lib/response"

	"github.com/gin-gonic/gin"
)

type installmentController struct {
	Name    string
	BaseUrl string
}

func InitInstallmentController() *installmentController {
	baseUrl := env.String("MainSetup.ServerHost", "")
	return &installmentController{
		Name:    "InstallmentController - ",
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
func (ctrl *loanController) GetInstallment(c *gin.Context) {
	logName := ctrl.Name + " Get Installment "
	res := response.Response{
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/installment", nil, "GET", nil, &res)
	log.LogFmtTemp("Start " + logName)
	defer log.LogFmtTemp("End " + logName)

	base.InitInstallmentService(logName, traceID).GetInstallment(&res)
	response.JsonGen(c, res)

}

func (ctrl *loanController) InsertInstallment(c *gin.Context) {
	logName := ctrl.Name + " Insert Installment "
	res := response.Response{
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	var req models.ReqInstallment
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/installment", nil, "POST", &req, &res)
	log.LogFmtTemp("Start " + logName)
	defer log.LogFmtTemp("End " + logName)
	// res.Meta.TraceID = traceID
	// ini buat POST
	dataBody, err := c.GetRawData()

	if err != nil {
		log.LogData(traceID, logName, "GetRawData", constant.LEVEL_LOG_ERROR, err.Error())
		res.Meta.DebugParam = err.Error()
		response.JsonGen(c, res)
		return
	}
	if err = json.Unmarshal(dataBody, &req); err != nil {
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

	base.InitInstallmentService(logName, traceID).InsertInstallment(req, &res)
	response.JsonGen(c, res)

}