package version

import (
	"discheck/constants"
)

func IsOutdated() bool {
	return constants.SOFTWARE_VERSION < 5.01
}
