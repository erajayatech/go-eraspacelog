package eraspacelog

import (
	"context"

	"github.com/gin-gonic/gin"
)

// use module path to make it unique
const authHeaderContextKey string = "github.com/erajayatech/go-eraspacelog/eraspacelog.AuthHeaderInfo"

type AuthHeaderInfo map[string]interface{}

func SetAuthHeaderInfoToContext(ginContext *gin.Context, authHeaderInfo AuthHeaderInfo) {
	ginContext.Set(authHeaderContextKey, authHeaderInfo)
}

func GetAuthHeaderInfoFromContext(context context.Context) *AuthHeaderInfo {
	headerInfo, ok := context.Value(authHeaderContextKey).(AuthHeaderInfo)
	if !ok {
		return nil
	}

	return &headerInfo
}
