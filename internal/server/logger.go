package server

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// ZapLogger is an Echo middleware that logs incoming HTTP requests using Uber's Zap logger.
func ZapLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()

			// Process the request
			err := next(c)
			if err != nil {
				c.Error(err) // Ensure Echo handles the error properly
			}

			// Retrieve request and response data
			req := c.Request()
			res := c.Response()

			// Extract the request ID
			requestID := req.Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = res.Header().Get(echo.HeaderXRequestID)
			}

			// Log fields with structured context
			logFields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.Duration("latency", time.Since(startTime)),
				zap.String("request_id", requestID),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
				zap.Int64("content_length", req.ContentLength),
				zap.String("user_agent", req.UserAgent()),
			}

			// Log based on HTTP status code category
			switch {
			case res.Status >= 500:
				logger.Error("Server error", logFields...)
			case res.Status >= 400:
				logger.Warn("Client error", logFields...)
			case res.Status >= 300:
				logger.Info("Redirection", logFields...)
			default:
				logger.Info("Request processed successfully", logFields...)
			}

			return err
		}
	}
}
