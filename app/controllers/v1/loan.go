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
		ResponseCode: constant.CODE_FAILED,
	}
	traceID := c.GetHeader("Trace-ID")
	// authorization := strings.Split(c.Request.Header.Get("Authorization"), " ")
	var req models.ReqBookingLoan
	defer log.LogSendRequest(traceID, logName, env.String("MainSetup.ServerHost", "")+"/v1/booking-loan", nil, "POST", &req, &res)
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

	base.InitLoanService(logName, traceID).BookingLoan(req, &res)
	response.JsonGen(c, res)

}
