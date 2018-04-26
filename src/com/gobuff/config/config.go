package config

type Config struct {
	Protocol string
	ServerIp string
	Port     int
}

func DefaultConfig() Config {
	conn := Config{"tcp", "127.0.0.1", 8545}

	return conn
}
