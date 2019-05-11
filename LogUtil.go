// LogUtil.go
package utility

import (
	"log"
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

func LogV(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelV) == LogLevelV {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeV).Append(content)
		log.Println(logBuffer.String())
	}
}

func LogD(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelD) == LogLevelD {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeD).Append(content)
		log.Println(logBuffer.String())
	}
}

func LogI(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelI) == LogLevelI {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeI).Append(content)
		log.Println(logBuffer.String())
	}
}

func LogW(content string) {
	LogInit()

	if (gLogger.LogLevel & LogLevelW) == LogLevelW {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeW).Append(content)
		log.Println(logBuffer.String())
	}
}

func LogE(content string) {
	LogInit()

	if gLogger.LogLevel&LogLevelE == LogLevelE {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeE).Append(content)
		log.Println(logBuffer.String())
	}
}

func (this Logger) logInit() {
	if !this.initialized {
		this.LogLevel = gDefaultLogLevel
		this.initialized = true
	}
}

func (this Logger) GetLogLevel() int {
	return this.LogLevel
}

func (this Logger) ChangeLogLevel(logLevel int) {
	this.LogLevel = logLevel
	this.initialized = true
}

func (this Logger) ChangeLogLevelByText(logLevelText string) {
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

func (this Logger) LogV(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelV) == LogLevelV {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeV).Append(content)
		log.Println(logBuffer.String())
	}
}

func (this Logger) LogD(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelD) == LogLevelD {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeD).Append(content)
		log.Println(logBuffer.String())
	}
}

func (this Logger) LogI(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelI) == LogLevelI {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeI).Append(content)
		log.Println(logBuffer.String())
	}
}

func (this Logger) LogW(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelW) == LogLevelW {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeW).Append(content)
		log.Println(logBuffer.String())
	}
}

func (this Logger) LogE(content string) {
	this.logInit()

	if (this.LogLevel & LogLevelE) == LogLevelE {
		var logBuffer StringBuilder
		logBuffer.Append(logTypeE).Append(content)
		log.Println(logBuffer.String())
	}
}
