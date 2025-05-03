package middleware

import (
	"bytes"
	"strings"

	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func buildRequestMessage(c *gin.Context) string {
	var result strings.Builder

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	result.WriteString(c.Request.RemoteAddr)
	result.WriteString(" - ")
	result.WriteString(c.Request.Method)
	result.WriteString(" ")
	result.WriteString(c.Request.RequestURI)
	result.WriteString(" - ")
	result.WriteString(blw.body.String())

	c.Next()

	return result.String()
}

func Logger(l logger.Interface) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		l.Info(buildRequestMessage(c))
	}
}
