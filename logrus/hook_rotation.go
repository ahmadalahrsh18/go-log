package logrus

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	_ "github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	logName = "hook_rotation.log"
)

func NewLfsHook() log.Hook {
	// new writer
	writer, err := rotatelogs.New(
		// 文件名
		logName+".%Y%m%d%H",
		// 为最新的日志创建软连接
		rotatelogs.WithLinkName(logName),
		// 设置日志分割时间，这里设置 1 小时
		rotatelogs.WithRotationTime(time.Hour),
		// 设置日志文件最长的保存时间或最大的保存数量（2 选 1）
		rotatelogs.WithMaxAge(time.Hour*24),
		//rotatelogs.WithRotationCount(24),
	)

	if err != nil {
		log.Fatalf("config local file system for logger error: %v", err)
	}

	// new hook
	hook := lfshook.NewHook(lfshook.WriterMap{
		log.TraceLevel: writer,
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return hook
}
