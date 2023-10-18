# go-eraspacelog

## Getting Started

eraspacelog is a structured logger for Go (golang), completely API compatible with the standard library logger.

## Dependency
* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [Logrus](https://github.com/sirupsen/logrus)
* [Open Telemetry](https://pkg.go.dev/go.opentelemetry.io/otel)
* [GoDotEnv by joho](https://github.com/joho/godotenv)

### Installation
Go Version 1.16+
```shell
go get github.com/erajayatech/go-eraspacelog
```
## Setup Environment
- Set the following environment variables:

* `MODE=<your_application_mode>`
  * Example : `prod`

## How To Use
- First you need to import go-eraspacelog package for using eraspacelog, one simplest example likes the follow

```go
package main

import (
  "github.com/erajayatech/go-eraspacelog"
)

func main() {
 ...
 eraspacelog.SetupLogger(helper.GetEnv("MODE"))

// your code goes here
}
```
`SetupLogger()` need paramater `mode` for determine formater to print into your terminal.
#### Example local mode
with local mode you'll see nicely color-coded
![Colored](https://i.ibb.co/HCTj2tz/log.png)

#### Example development mode
with development mode you'll see json
![raw](https://i.ibb.co/0J1FJCg/log-raw.png)

## Set auth-header and request-header
- Before implementation logger to every single function on your application, you must set the auth-header and request-header into a middleware.

#### Example set auth-header
to set auth-header, just call `SetAuthHeaderInfoToContext()`

```go
package middleware

import (
  "github.com/erajayatech/go-eraspacelog"
)

func (middleware *Middleware) HeaderValidatorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		source := context.Request.Header.Get("Source")
		
		eraspacelog.SetAuthHeaderInfoToContext(context, eraspacelog.AuthHeaderInfo{
			"source":        source,
		})
	}
}

// your code goes here
```

#### Example set request-header
to set request-header, just call `SetAuthHeaderInfoToContext()`

```go
package middleware

import (
	"github.com/erajayatech/go-eraspacelog"
	"github.com/gin-gonic/gin"
)

func (middleware *Middleware) TraceMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		var traceID = helper.GenerateUUID()

		context.Header("X-Trace-Id", traceID)
		context.Set("X-Trace-Id", traceID)
		context.Set("traceID", traceID)

		eraspacelog.SetRequestHeaderInfoToContext(context, eraspacelog.RequestHeaderInfo{
			"request_id": traceID,
			"path":       fmt.Sprintf("%s %s", context.Request.Method, context.Request.URL.Path),
		})

		context.Next()
	}
}

// your code goes here
```