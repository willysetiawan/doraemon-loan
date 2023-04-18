package config

import (
	"net/http"
	"process-loan/lib/env"

	"github.com/gin-gonic/gin"
)

func versioning(c *gin.Context) {
	res := Version{
		VersionRelease: env.String("Version.ReleaseVersion", ""),
		VersionType:    env.String("Version.VersionType", ""),
		ServiceName:    env.String("MainSetup.ServiceName", ""),
		ServiceType:    env.String("MainSetup.ServiceType", ""),
		Notes:          env.String("Version.Notes", ""),
		GitCommit:      GitCommit,
		BranchName:     env.String("Version.BranchName", ""),
		DateTime:       DateTime,
	}
	c.JSON(http.StatusOK, res)
}
