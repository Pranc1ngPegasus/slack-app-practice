package configuration

import (
	"io"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		ListenPort string `required:"true" envconfig:"PORT" default:"3000"`

		Slack struct {
			SigningSecret string `required:"true" envconfig:"SLACK_SIGNING_SECRET" default:""`
		}
	}
)

var (
	globalConfig Config
)

func Usage(output io.Writer) {
	if err := envconfig.Usagef("", &globalConfig, output, envconfig.DefaultTableFormat); err != nil {
		panic(err.Error())
	}
}

func Load() {
	envconfig.MustProcess("", &globalConfig)
}

func Get() Config {
	return globalConfig
}
