package structs

type Config struct {
	Proxies struct {
		Enabled bool   `yaml:"enabled"`
		Type    string `yaml:"type"`
		Path    string `yaml:"path"`
		Timeout int    `yaml:"timeout"`
		Order   string `yaml:"order"`
	} `yaml:"proxies"`

	Tokens struct {
		Regex string `yaml:"regex"`
		Path  string `yaml:"path"`
	} `yaml:"tokens"`

	Settings struct {
		Threads int `yaml:"threads"`
		Timeout int `yaml:"timeout"`
		Delay   int `yaml:"delay"`
	} `yaml:"settings"`
}

type DiscordSuccessResponse struct {
	Code   int `json:"code"`
	Errors struct {
		UserID struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"user_id"`
	} `json:"errors"`
	Message string `json:"message"`
}
