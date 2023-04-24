package collection

import (
	v1 "process-loan/app/controllers/v1"
	"process-loan/config/middleware"
	"process-loan/lib/env"

	"github.com/gin-gonic/gin"
)

func Loan(main *gin.RouterGroup) {
	ctrl := v1.InitLoanController()
	group := main.Group(env.String("InternalRouting.V1.Prefix", ""))
	group.Use(middleware.ValidateSymetricSignature())
	{
		group.GET(env.String("InternalRouting.V1.Installment.Send", ""), ctrl.GetInstallment)
		group.POST(env.String("InternalRouting.V1.Installment.Send", ""), ctrl.InsertInstallment)
		group.POST(env.String("InternalRouting.V1.BookingLoan.Send", ""), ctrl.BookingLoan)
	}

}
