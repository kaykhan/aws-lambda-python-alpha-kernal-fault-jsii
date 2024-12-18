package Config

type ConfigOptions struct {
	ENV           string
}

var Env ConfigOptions

func Init() {
	configuration := ConfigOptions{
		ENV: "dev",
	}

	Env = configuration
}
