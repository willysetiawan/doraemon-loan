package collection

import (
	v1 "process-loan/app/controllers/v1"
	"process-loan/lib/env"

	"github.com/gin-gonic/gin"
)

func Whitelist(main *gin.RouterGroup) {
	ctrl := v1.InitWhitelistController()
	group := main.Group(env.String("InternalRouting.V1.Prefix", ""))
	// group.Use(middleware.ValidateAccessToken())
	{
		group.POST(env.String("InternalRouting.V1.CheckWhitelist.Send", ""), ctrl.CheckWhitelist)
	}

}
