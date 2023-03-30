package example

import (
	"fmt"

	"github.com/erajayatech/go-eraspacelog"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// inject to middleware
func TraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, _ := uuid.NewV4()
		traceID := uuid.String()

		// set trace id to header and to the context gin
		ctx.Header("x-trace-id", traceID)
		ctx.Set("x-trace-id", traceID)

		ctx.Next()
	}
}

// sample log
func Hello(ctx *gin.Context) {
	// log print console
	eraspacelog.New().Print("INFO", map[string]interface{}{
		"function": "Hello",
		"message":  "this message will show in terminal",
	})

	// log message
	eraspacelog.New().Log("INFO", map[string]interface{}{
		"function": "Hello",
		"message":  "just a message",
	})

	// log with trace id
	eraspacelog.New().LogWithTrace("INFO",
		fmt.Sprintf("%v", ctx.Value("x-trace-id")),
		map[string]interface{}{
			"function": "Hello",
			"message":  "hello bro",
		})

	ctx.JSON(200, gin.H{
		"status":   "OK",
		"message":  "Hello bro!",
		"trace_id": fmt.Sprintf("%v", ctx.Value("x-trace-id")),
	})
}

// sample error with trace id
func Error(ctx *gin.Context) {
	errorMessage := "something went wrong"

	eraspacelog.New().LogWithTrace("ERROR",
		fmt.Sprintf("%v", ctx.Value("x-trace-id")),
		map[string]interface{}{
			"function": "Error",
			"message":  errorMessage,
		})

	ctx.JSON(200, gin.H{
		"status":   "ERROR",
		"message":  errorMessage,
		"trace_id": fmt.Sprintf("%v", ctx.Value("x-trace-id")),
	})
}

func main() {
	app := gin.Default()

	// route
	apiv1 := app.Group("/v1", TraceMiddleware())
	{
		apiv1.GET("/hello", Hello)
		apiv1.GET("/error", Error)
	}
}
