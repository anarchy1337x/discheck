package http

import (
	"discheck/constants"
	"discheck/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CheckAccount(token string) bool {
	client := &http.Client{}
	data, err := http.NewRequest("GET", fmt.Sprintf("https://discord.com/api/v%d/users/@me", constants.API_VERSION), nil)
	data.Header = http.Header{
		"Authorization": []string{token},
	}

	req, err := client.Do(data)

	if err != nil {
		return false
	}

	defer req.Body.Close()

	raw, err := io.ReadAll(req.Body)
	if err != nil {
		return false
	}

	if req.StatusCode == 200 {
		var discord_response structs.DiscordSuccessResponse
		err = json.Unmarshal(raw, &discord_response)
		if err != nil {
			return false
		}
		return true
	} else if req.StatusCode == 401 {
		return false
	}
	return false
}

func GetVersion() string {
	client := &http.Client{}
	data, err := http.NewRequest("GET", "https://raw.githubusercontent.com/anarchy1337x/Discheck/master/version.txt", nil)

	req, err := client.Do(data)

	if err != nil {
		return ""
	}

	defer req.Body.Close()

	raw, err := io.ReadAll(req.Body)
	if err != nil {
		return ""
	}

	return string(raw)
}
