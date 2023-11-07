package eraspacelog

import (
	"context"

	"github.com/gin-gonic/gin"
)

const requestHeaderContextKey string = "github.com/erajayatech/go-eraspacelog/eraspacelog.RequestHeaderInfo"

type RequestHeaderInfo map[string]interface{}

func SetRequestHeaderInfoToContext(ginContext *gin.Context, requestHeaderInfo RequestHeaderInfo) {
	ginContext.Set(requestHeaderContextKey, requestHeaderInfo)
}

func GetRequestHeaderInfoFromContext(ctx context.Context) *RequestHeaderInfo {
	requestInfo, ok := ctx.Value(requestHeaderContextKey).(RequestHeaderInfo)
	if !ok {
		return nil
	}
	return &requestInfo
}

func (requestHeaderInfo *RequestHeaderInfo) ToString() string {
	if requestHeaderInfo == nil {
		return ""
	}

	return Dump(requestHeaderInfo)
}
