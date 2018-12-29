package config

type Config struct {
	AppName string
	HttpPort int
}

func InitConfig() Config  {
	return Config{
		AppName: "Soctra Go",
		HttpPort: 8080,
	}
}