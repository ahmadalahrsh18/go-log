package logrus

import (
	log "github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLevelLogging(t *testing.T) {
	convey.Convey("TestLevelLogging", t, func() {
		log.SetLevel(log.InfoLevel)
		log.Trace("Something very low level.")
		log.Debug("Useful debugging information.")
		log.Info("Something noteworthy happened.")
		log.Warn("You should probably take a look at this.")
		log.Error("Something failed but I'm not quitting.")
		// Calls os.Exit(1) after logging
		log.Fatal("Bye.")
		// Calls panic() after logging
		log.Panic("I'm bailing.")
	})
}

func TestFields(t *testing.T) {
	convey.Convey("TestFields", t, func() {
		log.Infof("Failed to send event %s to topic %s with key %s", "event", "topic", "key")
		// 当然，上面那种日志输出方式也可以使用，但是 WithFields 迫使我们使用一种更规范的日志格式进行日志记录
		log.WithFields(log.Fields{
			"event": "event",
			"topic": "topic",
			"key":   "key",
		}).Info("Info to send event.")
		// Default Fields
		logger := log.WithFields(log.Fields{
			"event": "event",
			"topic": "topic",
			"key":   "key",
		})
		logger.Info("info active.")
		logger.Warn("warn active.")
	})
}

func TestFormatters(t *testing.T) {
	convey.Convey("TestFormatters", t, func() {
		log.SetFormatter(&log.TextFormatter{
			ForceColors:               true,
			DisableColors:             false,
			ForceQuote:                false,
			DisableQuote:              false,
			EnvironmentOverrideColors: false,
			DisableTimestamp:          true,
			FullTimestamp:             false,
			TimestampFormat:           "",
			DisableSorting:            false,
			SortingFunc:               nil,
			DisableLevelTruncation:    false,
			PadLevelText:              false,
			QuoteEmptyFields:          false,
			FieldMap:                  nil,
			CallerPrettyfier:          nil,
		})
		log.Info("info active.")
		log.Warn("warn active.")
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat:   "",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier:  nil,
			PrettyPrint:       false,
		})
		log.Info("info active.")
		log.SetReportCaller(true)
		log.Info("info active.")
	})
}

func TestFatalHandlers(t *testing.T) {
	convey.Convey("TestFatalHandlers", t, func() {
		handler := func() {
			log.Warn("gracefully shutdown.")
		}
		log.RegisterExitHandler(handler)
		log.Fatalln("log.Fatalln active.")
	})
}

func TestHook(t *testing.T) {
	convey.Convey("TestHook", t, func() {
		log.AddHook(new(doSomethingHook))
		log.Info("do something hook active.")
		log.AddHook(NewLfsHook())
		log.Info("lfs hook active.")
	})
}
