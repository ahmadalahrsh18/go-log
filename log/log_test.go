package log

import (
	_ "github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	convey.Convey("TestPrint", t, func() {
		// Println 写到标准日志记录器
		log.Print("log.Print active.")
		log.Printf("%s active.", "log.Printf")
		log.Println("log.Println active.")
	})
}

func TestPanic(t *testing.T) {
	convey.Convey("TestPanic", t, func() {
		defer func() {
			r := recover()
			convey.So(r, convey.ShouldNotBeNil)
		}()
		// Panicln 在调用 Println() 之后会接着调用 panic()
		log.Panicln("log.Panicln active.")
	})
}

func TestFatal(t *testing.T) {
	convey.Convey("TestFatal", t, func() {
		// Fatalln 在调用 Println() 之后会接着调用 os.Exit(1)
		log.Fatalln("log.Fatalln active.")
	})
}

func TestLogPrefixAndFlags(t *testing.T) {
	convey.Convey("TestLogPrefixAndFlags", t, func() {
		log.Println("log.Println active.")
		// set prefix
		log.SetPrefix("TRACE: ")
		// set flags
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
		log.Println("log.Println active.")
	})
}

var (
	Trace *log.Logger // 记录所有日志
	Info  *log.Logger // 重要的信息
	Warn  *log.Logger // 需要注意的信息
	Error *log.Logger // 非常严重的问题
)

func TestLoggers(t *testing.T) {
	convey.Convey("TestLoggers", t, func() {
		// open file
		file, err := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open error log file:", err)
		}

		// init loggers
		Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
		Error = log.New(io.MultiWriter(os.Stderr, file), "ERROR: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

		// use
		Trace.Println("trace active.")
		Info.Println("info active.")
		Warn.Println("warn active.")
		Error.Println("error active.")
	})
}
