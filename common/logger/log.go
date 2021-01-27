package logger

import (
	"fmt"
	"github.com/flowers-bloom/gin_web_scaffold/util/sys"
	"github.com/flowers-bloom/gin_web_scaffold/util/time"
	"github.com/spf13/viper"
)

const (
	Debug = iota
	Info
	Warn
	Error

	DefaultCallDepth  = 3
	DefaultFileMaxLength = 100
)

type LogWriter interface {
	Write(data interface{})
}

type Logger struct {
	level int
	writerList []LogWriter
}

var logger *Logger

func Log(level string, data string) {
	prefix := sys.GetCurFuncAndLine(DefaultCallDepth)
	data = fmt.Sprintf("%s %s %s %s", time.GetCurrentTime(), level, prefix, data)

	for _, writer := range logger.writerList {
		writer.Write(data)
	}
}

func GetLogger() *Logger {
	if logger == nil {
		Init(viper.GetString("logger.level"))
		writerList := viper.GetStringSlice("logger.writers")

		for _, writer := range writerList {
			SetWriters(writer)
		}
	}

	return logger
}

func Init(level string) *Logger {
	intLevel := Info

	switch level {
	case "panic":
		intLevel = Error
	case "warn":
		intLevel = Warn
	case "info":
		intLevel = Info
	case "debug":
		intLevel = Debug
	}

	logger = &Logger{
		level: intLevel,
		writerList: []LogWriter{},
	}

	return logger
}

func SetWriters(writer string) {
	if logger != nil {
		switch writer {
		case "std":
			logger.writerList = append(logger.writerList, &StdWriter{})
		case "file":
			var maxLength int64 = DefaultFileMaxLength
			maxLength = viper.GetInt64("logger.file.max_length")

			fw := &FileWriter{
				maxLength:maxLength,
			}
			fw.setFile(time.GetCurrentTimeF("2006-01-02") + ".log")
			logger.writerList = append(logger.writerList, fw)
		default:
			fmt.Printf("unknown writer type: %s\n", writer)
		}
	}else {
		fmt.Println("can not set writer, logger is nil")
	}
}

func (log *Logger) ErrorF(format string, v ...interface{}) {
	if Error < log.level {
		return
	}

	data := fmt.Sprintf(format, v...)
	Log("[Error]", data)
}

func (log *Logger) ErrorL(v ...interface{}) {
	if Error < log.level {
		return
	}

	data := fmt.Sprintf("%v", v)
	Log("[Error]", data)
}

func (log *Logger) WarnF(format string, v ...interface{}) {
	if Warn < log.level {
		return
	}

	data := fmt.Sprintf(format, v...)
	Log("[Warn]", data)
}

func (log *Logger) WarnL(v ...interface{}) {
	if Warn < log.level {
		return
	}

	data := fmt.Sprintf("%v", v)
	Log("[Warn]", data)
}

func (log *Logger) InfoF(format string, v ...interface{}) {
	if Info < log.level {
		return
	}

	data := fmt.Sprintf(format, v...)
	Log("[Info]", data)
}

func (log *Logger) InfoL(v ...interface{}) {
	if Info < log.level {
		return
	}

	data := fmt.Sprintf("%v", v)
	Log("[Info]", data)
}

func (log *Logger) DebugF(format string, v ...interface{}) {
	if Debug < log.level {
		return
	}

	data := fmt.Sprintf(format, v...)
	Log("[Debug]", data)
}

func (log *Logger) DebugL(v ...interface{}) {
	if Debug < log.level {
		return
	}

	data := fmt.Sprintf("%v", v)
	Log("[Debug]", data)
}