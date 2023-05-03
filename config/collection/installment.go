package collection

import (
	v1 "process-loan/app/controllers/v1"
	"process-loan/lib/env"

	"github.com/gin-gonic/gin"
)

func Installment(main *gin.RouterGroup) {
	ctrl := v1.InitInstallmentController()
	group := main.Group(env.String("InternalRouting.V1.Prefix", ""))
	{
		group.GET(env.String("InternalRouting.V1.Installment.Send", ""), ctrl.GetInstallment)
	}
}
