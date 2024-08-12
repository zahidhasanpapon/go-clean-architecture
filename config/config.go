package config

type (
	Config struct {
		App `yaml:"app"`
	}

	App struct {
		Name string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	}
)
