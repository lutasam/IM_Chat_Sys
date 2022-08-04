package utils

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//	"os"
//)
//
//type Logger struct {
//	tempLogger *logrus.Logger
//	ctx        *gin.Context
//}
//
//func GetCtxLogger(c *gin.Context) Logger {
//	logger := logrus.New()
//	logger.SetFormatter(&logrus.JSONFormatter{
//		TimestampFormat: "2006-01-02 15:04:05",
//	})
//	logger.SetLevel(logrus.InfoLevel) // 设置日志级别
//	logger.SetReportCaller(false)     // 设置在输出日志中添加文件名和方法信息
//	logfile, _ := os.OpenFile("./bytedance_begin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
//	logger.Out = logfile
//	return Logger{tempLogger: logger, ctx: c}
//}
//
//func (l *Logger) DoInfo(info string) {
//	l.tempLogger.WithFields(logrus.Fields{
//		"tracer_id": GetTracerId(l.ctx),
//	}).Info(info)
//}
//
//func (l *Logger) DoError(err string) {
//	l.tempLogger.WithFields(logrus.Fields{
//		"tracer_id": GetTracerId(l.ctx),
//	}).Error(err)
//}
//
//func GetTracerId(c *gin.Context) string {
//	trace, _ := c.Get("trace")
//	traceContext, _ := trace.(*TraceContext)
//	traceId := ""
//	if traceContext != nil {
//		traceId = traceContext.TraceId
//	}
//	return traceId
//}
