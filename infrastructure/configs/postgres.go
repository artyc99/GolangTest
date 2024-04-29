package configs

type PostgresConfig struct {
	ConnectionString string `env:"DB_CONNECTION,required" validate:"required"`
}
