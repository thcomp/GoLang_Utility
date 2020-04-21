package utility

import (
	"strconv"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	var version = `v1.9.0` // TODO 変更する度に更新すること
	var versionSep []string = strings.Split(version, `.`)
	versionMajorCode, _ := strconv.Atoi(strings.Replace(versionSep[0], `v`, ``, -1))
	versionMinorCode, _ := strconv.Atoi(versionSep[1])
	versionRevisionCode, _ := strconv.Atoi(versionSep[2])

	if Version() != version {
		t.Error(`version is not matched`)
	}

	if VersionMajorCode() != versionMajorCode {
		t.Error(`major code is not matched`)
	}

	if VersionMinorCode() != versionMinorCode {
		t.Error(`minor code is not matched`)
	}

	if VersionRevisionCode() != versionRevisionCode {
		t.Error(`revision code is not matched`)
	}
}
