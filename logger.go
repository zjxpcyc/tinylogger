package tinylogger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	// 日志 level 级别
	// 模拟 https://golang.google.cn/pkg/log/syslog/#Priority
	// 重新定义, 是为了防止 windows 系统报错, 因为 syslog 不支持 windows 系统
	LOG_EMERG = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

// LogService 日志服务
// 日志 level 级别 https://golang.google.cn/pkg/log/syslog/#Priority
type LogService interface {
	// Error 级别日志
	Error(...interface{})

	// Info 级别错误
	Info(...interface{})
}

// Logger 是 LogService 接口的默认实现
// 默认是 输出到 os.Stdout 或者 os.Stderr
type Logger struct {
	w io.Writer
}

// NewLogger returns a new logger.
// If w is nil, os.StdOut will be used in commone log, else os.Stderr will be used
func NewLogger(w ...io.Writer) *Logger {
	if w == nil || len(w) == 0 {
		return new(Logger)
	}

	return &Logger{
		w: w[0],
	}
}

// Info 级别日志
func (t *Logger) Info(v ...interface{}) {
	t.write(LOG_INFO, v...)
}

// Error 级别日志
func (t *Logger) Error(v ...interface{}) {
	t.write(LOG_ERR, v...)
}

// write 日志内容到 target.
func (t *Logger) write(level int, v ...interface{}) {
	var prefix string
	switch level {
	case LOG_EMERG:
		prefix = "EMERG"
	case LOG_ALERT:
		prefix = "ALERT"
	case LOG_CRIT:
		prefix = "CRIT"
	case LOG_ERR:
		prefix = "ERR"
	case LOG_WARNING:
		prefix = "WARNING"
	case LOG_NOTICE:
		prefix = "NOTICE"
	case LOG_INFO:
		prefix = "INFO"
	case LOG_DEBUG:
		prefix = "DEBUG"
	}

	now := time.Now().Local().Format("2006-01-02 15:04:05")

	_, file, line, _ := runtime.Caller(2)
	_, f := path.Split(file)

	msg := strings.TrimSuffix(strings.TrimPrefix(fmt.Sprintf("%v", v), "["), "]")
	content := fmt.Sprintf("[%s] [%s] [%s:%d] %s \n", now, prefix, f, line, msg)

	if t.w != nil {
		t.w.Write([]byte(content))
		return
	}

	if level >= LOG_WARNING {
		os.Stdout.WriteString(content)
	} else {
		os.Stderr.WriteString(content)
	}
}

var logger LogService

// GetLogger 获取日志记录器
func GetLogger() LogService {
	return logger
}

func init() {
	logger = new(Logger)
}
