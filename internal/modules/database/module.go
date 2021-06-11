package database

import (
	"errors"
	"go-template/internal/modules/config"
	"go-template/internal/pkg/database"

	"gorm.io/gorm"
)

type DatabaseModule struct {
	gd *database.GormDatabase
}

var dm *DatabaseModule

func DM() *DatabaseModule {
	if dm == nil {
		panic(errors.New("database module is not inited"))
	}
	return dm
}

func init() {
	cm := config.CM()

	// logger
	logger := database.NewLogger()

	// postgres
	pg := database.NewPostgres(
		database.PostgresHostOption(cm.Config().Postgres.Host),
		database.PostgresPortOption(cm.Config().Postgres.Port),
		database.PostgresDatabaseOption(cm.Config().Postgres.Database),
		database.PostgresUserOption(cm.Config().Postgres.User),
		database.PostgresPasswordOption(cm.Config().Postgres.Passwrod),
	)

	gd := database.NewGorm(
		database.GormDatabaseDialector(pg.Dialector()),
		database.GormDatabaseMigration(cm.Config().Database.Migration),
		database.GormDatabaseLogger(logger.GormLogger()),
	)

	// database module
	dm = &DatabaseModule{gd: gd}
}

func (dm *DatabaseModule) Gorm() *gorm.DB {
	return dm.gd.Gorm()
}

func (dm *DatabaseModule) Migrate(entities ...interface{}) {
	dm.gd.Migrate(entities...)
}
