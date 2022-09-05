package configuration

type Configuration struct {
	AppPort string
}

func GetConfigs() (*Configuration, error) {
	return &Configuration{
		AppPort: "8083",
	}, nil
}
