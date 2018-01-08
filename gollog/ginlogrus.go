package gollog

import (
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	GinLogger     = logrus.New()
	GinTimeFormat = time.RFC3339
	GinEnableUTC  = false
)

func ginLogrus(ctx *gin.Context) {
	startCtx := time.Now()
	path := ctx.Request.URL.Path

	ctx.Next()

	stopCtx := time.Since(startCtx)
	msUsed := int(math.Ceil(float64(stopCtx.Nanoseconds()) / 1000.0))

	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()
	clientUserAgent := ctx.Request.UserAgent()
	referer := ctx.Request.Referer()
	dataLength := ctx.Writer.Size()
	httpMethod := ctx.Request.Method
	if dataLength < 0 {
		dataLength = 0
	}

	if GinEnableUTC {
		startCtx = startCtx.UTC()
	}
	startTime := startCtx.Format(GinTimeFormat)

	entry := logrus.NewEntry(GinLogger).WithFields(logrus.Fields{
		"microseconds-used": msUsed,
		"client-ip":         clientIP,
		"user-agent":        clientUserAgent,
		"http-status-code":  statusCode,
		"http-method":       httpMethod,
		"http-path":         path,
		"http-referer":      referer,
		"data-length":       dataLength,
		"requested-at":      startTime,
	})

	if len(ctx.Errors) > 0 {
		entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		if statusCode > 499 {
			entry.Error()
		} else if statusCode > 399 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

func GinLogrus() gin.HandlerFunc {
	logrus.SetOutput(os.Stdout)
	GinLogger.Formatter = &logrus.JSONFormatter{}
	GinLogger.Level = logrusLevel()
	return ginLogrus
}
