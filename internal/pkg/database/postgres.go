package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOption func(*Postgres)

func PostgresHostOption(host string) PostgresOption {
	return func(p *Postgres) {
		if host == "" {
			return
		}
		p.Host = host
	}
}

func PostgresPortOption(port int) PostgresOption {
	return func(p *Postgres) {
		if port == 0 {
			return
		}
		p.Port = port
	}
}

func PostgresDatabaseOption(db string) PostgresOption {
	return func(p *Postgres) {
		if db == "" {
			return
		}
		p.Database = db
	}
}

func PostgresUserOption(user string) PostgresOption {
	return func(p *Postgres) {
		if user == "" {
			return
		}
		p.User = user
	}
}

func PostgresPasswordOption(pwd string) PostgresOption {
	return func(p *Postgres) {
		if pwd == "" {
			return
		}
		p.Passwrod = pwd
	}
}

type Postgres struct {
	Host     string
	Port     int
	Database string
	User     string
	Passwrod string
}

func NewPostgres(opts ...PostgresOption) *Postgres {
	p := &Postgres{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Passwrod: "postgres",
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p Postgres) dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		p.Host,
		p.User,
		p.Passwrod,
		p.Database,
		p.Port,
	)
}

func (p Postgres) Dialector() gorm.Dialector {
	return postgres.Open(p.dsn())
}
