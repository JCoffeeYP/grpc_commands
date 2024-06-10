package config

type Config struct {
	RedisHost string `env:"IGNITE_HOST" env-default:"localhost"`
	RedisPort int    `env:"IGNITE_PORT" env-default:"6379"`
	LogLevel  string `env:"LOG_LEVEL" env-default:"INFO"`
}
