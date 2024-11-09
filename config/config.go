package config

type Config struct {
	Port int `env:"PORT"`
	Host string `env:"HOST"`
	TotalServers int `env:"TOTAL_SERVERS"`
	
}

