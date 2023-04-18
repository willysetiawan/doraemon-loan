package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthData struct {
	ID        string
	Email     string
	RoleID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAuthUserCtx(ctx *gin.Context) (*AuthData, error) {
	authDataStr := ctx.GetString("auth")

	var authData AuthData
	err := json.Unmarshal([]byte(authDataStr), &authData)
	if err != nil {
		return nil, err
	}

	return &authData, nil
}
