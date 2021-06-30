// LogUtil.go
package utility

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type LoggerIF interface {
	logInit()
	GetLogLevel() int
	ChangeLogLevel(logLevel int)
	LogV(content string)
	LogD(content string)
	LogI(content string)
	LogW(content string)
	LogE(content string)
}

type Logger struct {
	LogLevel    int
	initialized bool
	outputFile  string
	useStdOut   bool
}

const LogLevelE = 1
const LogLevelW = 3
const LogLevelI = 7
const LogLevelD = 15
const LogLevelV = 31

const logTypeV = `V: `
const logTypeD = `D: `
const logTypeI = `I: `
const logTypeW = `W: `
const logTypeE = `E: `

var gLogger Logger

const gDefaultLogLevel = LogLevelW | LogLevelE

func LogInit() {
	if !gLogger.initialized {
		gLogger.LogLevel = gDefaultLogLevel
		gLogger.initialized = true
	}
}

func GetLogLevel() int {
	LogInit()
	return gLogger.LogLevel
}

func ChangeLogLevel(logLevel int) {
	gLogger.LogLevel = logLevel
	gLogger.initialized = true
}

func ChangeLogLevelByText(logLevelText string) {
	if logLevel, getRet := getLogLevel(logLevelText); getRet {
		gLogger.LogLevel = logLevel
		gLogger.initialized = true
	}
}

func ChangeOutput(outputFile string) {
	gLogger.outputFile = outputFile
}

func UseStdout(useStdout bool) {
	gLogger.UseStdout(useStdout)
}

func LogV(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelV) == LogLevelV {
		gLogger.LogV(content)
	}
}

func LogD(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelD) == LogLevelD {
		gLogger.LogD(content)
	}
}

func LogI(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelI) == LogLevelI {
		gLogger.LogI(content)
	}
}

func LogW(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelW) == LogLevelW {
		gLogger.LogW(content)
	}
}

func LogE(content string) {
	LogInit()

	if gLogger.LogLevel&LogLevelE == LogLevelE {
		gLogger.LogE(content)
	}
}

func LogfV(format string, args ...interface{}) {
	LogInit()

	if (gLogger.LogLevel & LogLevelV) == LogLevelV {
		gLogger.LogfV(format, args...)
	}
}

func LogfD(format string, args ...interface{}) {
	LogInit()

	if (gLogger.LogLevel & LogLevelD) == LogLevelD {
		gLogger.LogfD(format, args...)
	}
}

func LogfI(format string, args ...interface{}) {
	LogInit()

	if (gLogger.LogLevel & LogLevelI) == LogLevelI {
		gLogger.LogfI(format, args...)
	}
}

func LogfW(format string, args ...interface{}) {
	LogInit()

	if (gLogger.LogLevel & LogLevelW) == LogLevelW {
		gLogger.LogfW(format, args...)
	}
}

func LogfE(format string, args ...interface{}) {
	LogInit()

	if gLogger.LogLevel&LogLevelE == LogLevelE {
		gLogger.LogfE(format, args...)
	}
}

func (this *Logger) logInit() {
	if !this.initialized {
		this.LogLevel = gDefaultLogLevel
		this.initialized = true
	}
}

func (this *Logger) GetLogLevel() int {
	return this.LogLevel
}

func (this *Logger) ChangeLogLevel(logLevel int) {
	this.LogLevel = logLevel
	this.initialized = true
}

func (this *Logger) ChangeLogLevelByText(logLevelText string) {
	if logLevel, getRet := getLogLevel(logLevelText); getRet {
		this.LogLevel = logLevel
		this.initialized = true
	}
}

func getLogLevel(logLevelText string) (int, bool) {
	var logLevelLow string = strings.ToLower(logLevelText)
	var logLevelInt int = 0
	var ret bool = true

	if logLevelLow == "debug" || logLevelLow == "d" {
		logLevelInt = LogLevelD
	} else if logLevelLow == "verbose" || logLevelLow == "v" {
		logLevelInt = LogLevelV
	} else if logLevelLow == "info" || logLevelLow == "i" {
		logLevelInt = LogLevelI
	} else if logLevelLow == "warn" || logLevelLow == "w" {
		logLevelInt = LogLevelW
	} else {
		ret = false
	}

	return logLevelInt, ret
}

func (this *Logger) LogV(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelV) == LogLevelV {
		this.output(logTypeV + content + "\n")
	}
}

func (this *Logger) LogD(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelD) == LogLevelD {
		this.output(logTypeD + content + "\n")
	}
}

func (this *Logger) LogI(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelI) == LogLevelI {
		this.output(logTypeI + content + "\n")
	}
}

func (this *Logger) LogW(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelW) == LogLevelW {
		this.output(logTypeW + content + "\n")
	}
}

func (this *Logger) LogE(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelE) == LogLevelE {
		this.output(logTypeE + content + "\n")
	}
}

func (this *Logger) LogfV(format string, args ...interface{}) {
	this.logInit()

	if (this.LogLevel & LogLevelV) == LogLevelV {
		this.output(logTypeV+format+"\n", args...)
	}
}

func (this *Logger) LogfD(format string, args ...interface{}) {
	this.logInit()

	if (this.LogLevel & LogLevelD) == LogLevelD {
		this.output(logTypeD+format+"\n", args...)
	}
}

func (this *Logger) LogfI(format string, args ...interface{}) {
	this.logInit()

	if (this.LogLevel & LogLevelI) == LogLevelI {
		this.output(logTypeI+format+"\n", args...)
	}
}

func (this *Logger) LogfW(format string, args ...interface{}) {
	this.logInit()

	if (this.LogLevel & LogLevelW) == LogLevelW {
		this.output(logTypeW+format+"\n", args...)
	}
}

func (this *Logger) LogfE(format string, args ...interface{}) {
	this.logInit()

	if (this.LogLevel & LogLevelE) == LogLevelE {
		this.output(logTypeE+format+"\n", args...)
	}
}

func (this *Logger) output(format string, args ...interface{}) {
	if len(this.outputFile) == 0 {
		if this.useStdOut {
			fmt.Printf(format, args...)
		} else {
			log.Printf(format, args...)
		}
	} else {
		if desc, openErr := os.OpenFile(this.outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); openErr == nil {
			defer desc.Close()
			desc.Write([]byte(fmt.Sprintf(format, args...)))
		}
	}
}

func (this *Logger) ChangeOutput(outputFile string) {
	this.logInit()
	this.outputFile = outputFile
}

func (this *Logger) UseStdout(useStdout bool) {
	this.logInit()
	this.useStdOut = useStdout
}
