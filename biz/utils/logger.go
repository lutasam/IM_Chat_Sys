package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	tempLogger *logrus.Logger
	ctx        *gin.Context
}

func GetCtxLogger(c *gin.Context) Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logfile, err := os.OpenFile(GetConfigResolve().GetConfigString("log.filepath"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	logger.Out = logfile
	return Logger{tempLogger: logger, ctx: c}
}

func (l *Logger) DoInfo(info string) {
	l.tempLogger.WithFields(logrus.Fields{
		"trace_id": GetTracerId(l.ctx),
	}).Info(info)
}

func (l *Logger) DoWarning(warning string) {
	l.tempLogger.WithFields(logrus.Fields{
		"trace_id": GetTracerId(l.ctx),
	}).Warning(warning)
}

func (l *Logger) DoError(err string) {
	l.tempLogger.WithFields(logrus.Fields{
		"trace_id": GetTracerId(l.ctx),
	}).Error(err)
}

func GetTracerId(c *gin.Context) string {
	trace, _ := c.Get("trace_id")
	traceContext, _ := trace.(*TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	return traceId
}
