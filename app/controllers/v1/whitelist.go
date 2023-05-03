package v1

import (
	"encoding/json"
	"process-loan/domain/base"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/env"
	"process-loan/lib/log"
	"process-loan/lib/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type whitelistController struct {
	Name    string
	BaseUrl string
}

func InitWhitelistController() *whitelistController {
	baseUrl := env.String("MainSetup.ServerHost", "")
	return &whitelistController{
		Name:    "WhitelistController - ",
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
func (ctrl *whitelistController) CheckWhitelist(c *gin.Context) {
	logName := ctrl.Name + " Check Whitelist "
	res := response.Response{
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	// authorization := strings.Split(c.Request.Header.Get("Authorization"), " ")
	var req models.ReqCheckWhitelist
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/whitelist", nil, "POST", &req, &res)
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

	base.InitWhitelistService(logName, traceID).CheckWhitelist(req, &res)
	response.JsonGen(c, res)

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
func (ctrl *whitelistController) GetWhitelist(c *gin.Context) {
	logName := ctrl.Name + " Get Whitelist "
	res := response.Response{
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/whitelist", nil, "GET", nil, &res)
	log.LogFmtTemp("Start " + logName)
	defer log.LogFmtTemp("End " + logName)

	base.InitWhitelistService(logName, traceID).GetWhitelist(&res)
	response.JsonGen(c, res)

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
func (ctrl *whitelistController) InsertWhitelist(c *gin.Context) {
	logName := ctrl.Name + " Insert Whitelist "
	res := response.Response{
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	// authorization := strings.Split(c.Request.Header.Get("Authorization"), " ")
	var req models.ReqWhitelist
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/whitelist", nil, "POST", &req, &res)
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

	base.InitWhitelistService(logName, traceID).AddWhitelist(req, &res)
	response.JsonGen(c, res)

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
func (ctrl *whitelistController) UploadWhitelist(c *gin.Context) {
	logName := ctrl.Name + " Upload Whitelist "
	res := response.Response{
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	// authorization := strings.Split(c.Request.Header.Get("Authorization"), " ")
	var req []models.ReqWhitelist
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/upload-whitelist", nil, "POST", &req, &res)
	log.LogFmtTemp("Start " + logName)
	defer log.LogFmtTemp("End " + logName)
	// res.Meta.TraceID = traceID

	reqPartnerId := c.Request.FormValue("partnerId")
	if reqPartnerId == "" {
		res.ResponseMessage = "Partner Id required"
		response.JsonGen(c, res)
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		res.ResponseMessage = constant.MESSAGE_SOMETHING_WENT_WRONG + " File whitelist"
		res.Meta.DebugParam = err.Error()
		response.JsonGen(c, res)
		return
	}

	f, err := excelize.OpenReader(file)
	//Remove First Row
	err = f.RemoveRow("Sheet1", 1)
	//Read All Rows
	rows, err := f.GetRows("Sheet1")
	//Looping All Rows Except First Row
	for index, row := range rows {
		parseSalary, _ := strconv.ParseFloat(row[7], 64)
		parseMaxLoanAmount, _ := strconv.ParseFloat(row[10], 64)
		if index >= 0 {
			req = append(req, models.ReqWhitelist{
				PartnerId:             reqPartnerId,
				CIF:                   row[1],
				EmployeeName:          row[2],
				EmployeeId:            row[3],
				EmployeeIdentityNo:    row[4],
				EmployeeMobilePhoneNo: row[5],
				EmployeeEmail:         row[6],
				EmployeeSalary:        parseSalary,
				EmployeeStatus:        row[8],
				PayrollMonth:          row[9],
				MaxLoan:               parseMaxLoanAmount,
			})
		}
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

	base.InitWhitelistService(logName, traceID).UploadWhitelist(req, &res)
	response.JsonGen(c, res)

}
