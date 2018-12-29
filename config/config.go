package config

type Config struct {
	AppName string
	HttpPort int
}

var AppConfig = Config{
	AppName: "SoctraGo",
	HttpPort: 8080,
}
