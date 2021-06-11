package database

import "go-template/internal/pkg/database"

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Passwrod string `mapstructure:"password"`
}

func (pc *PostgresConfig) Postgres() *database.Postgres {
	return database.NewPostgres(
		database.PostgresHostOption(pc.Host),
		database.PostgresPortOption(pc.Port),
		database.PostgresDatabaseOption(pc.Database),
		database.PostgresUserOption(pc.User),
		database.PostgresPasswordOption(pc.Passwrod),
	)
}
