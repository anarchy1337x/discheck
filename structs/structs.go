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
