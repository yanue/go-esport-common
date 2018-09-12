package middleware

import (
	"bytes"
	"github.com/yanue/go-esport-common/logger"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type bodyReader struct {
	io.ReadCloser
	body *bytes.Buffer
}

func (r *bodyReader) Read(p []byte) (n int, err error) {
	r.body.Read(p)
	return r.ReadCloser.Read(p)
}

// Gin日志中间件
// 不记录exclude中的url路径
//
// @param config *Config
// @param exclude ...string
//
// @return middleware.HandlerFunc
//
func GinLoggerMiddleware(config *logger.Config, exclude ...string) gin.HandlerFunc {
	log := logger.New(config)
	// 修改默认的logger
	logger.SetSugarLogger(log.Sugar())

	return func(c *gin.Context) {
		var logTemp *zap.Logger = log
		for _, item := range exclude {
			if strings.Contains(c.Request.RequestURI, item) {
				c.Next()
				return
			}
		}

		// Start timer
		start := time.Now()

		if id := c.GetString("RequestId"); id != "" {
			logTemp = logTemp.With(zap.String("requestId", id))
		}

		request := logTemp.Named("request").
			With(zap.String("client", c.ClientIP())).
			With(zap.String("method", c.Request.Method)).
			With(zap.String("url", c.Request.RequestURI))

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			bodyCopy := new(bytes.Buffer)
			_, err := io.Copy(bodyCopy, c.Request.Body)
			if err == nil {
				request = request.With(zap.String("body", bodyCopy.String()))
				c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyCopy.Bytes()))
			}
		}

		request.Info("")

		w := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		response := logTemp.Named("response").
			With(zap.Duration("latency", latency)).
			With(zap.Int("statusCode", c.Writer.Status())).
			With(zap.String("body", w.body.String()))

		response.Info("")
	}
}
