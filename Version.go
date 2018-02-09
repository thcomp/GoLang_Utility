package utility

import (
	"strconv"
)

const versionMajorCode = 1
const versionMinorCode = 2
const versionRevisionCode = 0

var version string

func Version() string {
	if len(version) == 0 {
		version = `v` + strconv.Itoa(versionMajorCode) + `.` + strconv.Itoa(versionMinorCode) + `.` + strconv.Itoa(versionRevisionCode)
	}

	return version
}

func VersionMajorCode() int {
	return versionMajorCode
}

func VersionMinorCode() int {
	return versionMinorCode
}

func VersionRevisionCode() int {
	return versionRevisionCode
}
