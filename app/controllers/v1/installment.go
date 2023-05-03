package v1

import (
	"process-loan/domain/base"
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
func (ctrl *installmentController) GetInstallment(c *gin.Context) {
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
