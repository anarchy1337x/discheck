package http

import (
	"discheck/constants"
	"fmt"
	"net/http"
)

func CheckAccount(token string) bool {
	_, err := http.Get(fmt.Sprintf("https://discord.com/api/v%d/auth/login", constants.API_VERSION))

	if err != nil {
		return false
	}

	return true
}
