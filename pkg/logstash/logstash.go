package logstash

import (
	"cmd/main.go/config"
	mylogger "cmd/main.go/pkg/logger"
	logstash_logger "github.com/KaranJagtiani/go-logstash"
)

type LogstashLoger struct {
	logstash *logstash_logger.Logstash
}

var GlobalLogstash LogstashLoger

func NewLogstashLogger(cfg config.Config) LogstashLoger {
	logstash := logstash_logger.Init(cfg.HttpServer.ElkDomain, cfg.HttpServer.ELKPort, "tcp", 5)
	GlobalLogstash = LogstashLoger{logstash: logstash}
	return GlobalLogstash
}

func (l LogstashLoger) LogstashError(message string, err error) {
	mylogger.GlobalLogger.Error(message + err.Error())
	l.logstash.Error(map[string]interface{}{"message": message, "error": err})
}

func (l LogstashLoger) LogstashInfo(message string) {
	mylogger.GlobalLogger.Info(message)
	l.logstash.Info(map[string]interface{}{"message": message})
}

func (l LogstashLoger) LogstashDebug(message string) {
	mylogger.GlobalLogger.Debug(message)
	l.logstash.Debug(map[string]interface{}{"message": message})
}

func (l LogstashLoger) LogstashWarn(message string) {
	mylogger.GlobalLogger.Warn(message)
	l.logstash.Warn(map[string]interface{}{"message": message})
}
