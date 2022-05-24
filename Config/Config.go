package Config

type Config struct {
	SlackWebhook string `yaml:"slackWebhook" env:"slackhook"`
	Timezone     string `yaml:"timezone"`
	Hostname     string `yaml:"hostname" env:"apexhost"`
	TestsADay    int8   `yaml:"tests"`
	Limits       struct {
		Temp struct {
			High float64 `yaml:"high" env:"tempHigh"`
			Low  float64 `yaml:"low" env:"tempLow"`
		} `yaml:"temp"`
		Ph struct {
			High float64 `yaml:"high" env:"phHigh"`
			Low  float64 `yaml:"low" env:"phLow"`
		} `yaml:"ph"`
	} `yaml:"limits"`
}
