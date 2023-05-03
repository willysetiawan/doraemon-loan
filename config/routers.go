package config

import (
	"process-loan/config/collection"
	"process-loan/docs"
	"process-loan/lib/env"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	Version struct {
		VersionRelease string `json:"version_release"`
		ServiceName    string
		ServiceType    string
		GitCommit      string
		BranchName     string
		DateTime       string
		VersionType    string
		Notes          string
	}
)

var (
	Routers   = gin.Default()
	GitCommit string
	DateTime  string
)

func init() {
	corsConfig(Routers)
	Routers.Static("/assets", "./assets")

	docs.SwaggerInfo.Title = env.String("MainSetup.ServiceName", "")
	docs.SwaggerInfo.Description = "This is a list of Loan  api for " + env.String("MainSetup.ServiceType", "") + env.String("MainSetup.ServiceName", "")
	docs.SwaggerInfo.Version = env.String("Version.ReleaseVersion", "")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = env.String("MainSetup.ServerHost", "")

	Routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Routers.GET("/version", versioning)

	v1 := Routers.Group(env.String("InternalRouting.Base", ""))
	// collection.Example(v1)
	collection.Loan(v1)
	collection.Installment(v1)
	collection.Whitelist(v1)
}
