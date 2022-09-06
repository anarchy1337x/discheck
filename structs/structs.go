package structs

type Config struct {
	Proxies struct {
		Enabled  bool   `yaml:"enabled"`
		Type     string `yaml:"type"`
		Path     string `yaml:"path"`
		Timeout  int    `yaml:"timeout"`
		Rotating bool   `yaml:"rotating"`
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
	ID                string      `json:"id"`
	Username          string      `json:"username"`
	Avatar            string      `json:"avatar"`
	AvatarDecoration  interface{} `json:"avatar_decoration"`
	Discriminator     string      `json:"discriminator"`
	PublicFlags       int         `json:"public_flags"`
	Flags             int         `json:"flags"`
	PurchasedFlags    int         `json:"purchased_flags"`
	PremiumUsageFlags int         `json:"premium_usage_flags"`
	Banner            interface{} `json:"banner"`
	BannerColor       string      `json:"banner_color"`
	AccentColor       int         `json:"accent_color"`
	Bio               string      `json:"bio"`
	Locale            string      `json:"locale"`
	NsfwAllowed       bool        `json:"nsfw_allowed"`
	MfaEnabled        bool        `json:"mfa_enabled"`
	PremiumType       int         `json:"premium_type"`
	Email             string      `json:"email"`
	Verified          bool        `json:"verified"`
	Phone             string      `json:"phone"`
}
