package version

import (
	"discheck/constants"
	http "discheck/networking"
	"strconv"
	"strings"
)

func IsOutdated() bool {
	version := http.GetVersion()
	version_float, _ := strconv.ParseFloat(strings.Trim(version, "\n"), 64)
	return version_float != constants.SOFTWARE_VERSION
}
