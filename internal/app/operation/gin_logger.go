package operation

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// LoggerConfig defines the config for Logger middleware.
type LoggerConfig struct {
	// SkipPaths is an url path array which logs are not written.
	// Optional.
	SkipPaths []string

	// LogLevel for everything else than errors
	// Optional. Default value is zerolog.DebugLevel
	DefaultLogLevel zerolog.Level
}

func GinLogger() gin.HandlerFunc {
	return GinLoggerWithConfig(LoggerConfig{})
}

func GinLoggerWithConfig(conf LoggerConfig) gin.HandlerFunc {
	notlogged := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		// Process request
		ctx.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: ctx.Request,
				Keys:    ctx.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = ctx.ClientIP()
			param.Method = ctx.Request.Method
			param.StatusCode = ctx.Writer.Status()
			param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()

			param.BodySize = ctx.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			logParams(ctx, param, conf)
		}
	}
}

func logParams(ctx *gin.Context, param gin.LogFormatterParams, conf LoggerConfig) {
	logEvent := createLogEvent(ctx, conf)
	logEvent.
		Str("client_id", param.ClientIP).
		Str("method", param.Method).
		Int("status_code", param.StatusCode).
		Int("body_size", param.BodySize).
		Str("path", param.Path).
		Str("latency", param.Latency.String()).
		Msg(param.ErrorMessage)
}

func createLogEvent(c *gin.Context, conf LoggerConfig) *zerolog.Event {
	var logEvent *zerolog.Event

	if c.Writer.Status() >= 500 {
		logEvent = Logger.Error()
	} else {
		logEvent = Logger.WithLevel(conf.DefaultLogLevel)
	}
	return logEvent
}
