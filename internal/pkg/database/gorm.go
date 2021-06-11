package database

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormDatabaseOption func(*GormDatabase)

func GormDatabaseMigration(b bool) GormDatabaseOption {
	return func(gd *GormDatabase) {
		gd.Migration = true
	}
}

func GormDatabaseDialector(d gorm.Dialector) GormDatabaseOption {
	return func(gd *GormDatabase) {
		gd.Dialector = d
	}
}

func GormDatabaseLogger(logger logger.Interface) GormDatabaseOption {
	return func(gd *GormDatabase) {
		gd.Logger = logger
	}
}

type GormDatabase struct {
	Migration      bool
	DB             *gorm.DB
	NamingStrategy schema.Namer
	Dialector      gorm.Dialector
	Logger         logger.Interface
}

func NewGorm(opts ...GormDatabaseOption) *GormDatabase {
	g := &GormDatabase{
		Migration: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "ls_",
		},
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func (gd *GormDatabase) open() {
	if gd.DB != nil {
		return
	}

	if gd.Dialector == nil {
		panic(errors.New("no dialector"))
	}

	db, err := gorm.Open(gd.Dialector, &gorm.Config{
		Logger:                                   gd.Logger,
		NamingStrategy:                           gd.NamingStrategy,
		DisableNestedTransaction:                 true,
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
	})

	if err != nil {
		panic(err)
	}

	gd.DB = db
}

func (gd *GormDatabase) Gorm() *gorm.DB {
	if gd.DB == nil {
		gd.open()
	}

	return gd.DB
}

func (gd *GormDatabase) Migrate(entities ...interface{}) {
	if gd.Migration == true {
		gd.Gorm().AutoMigrate(entities...)
	}
}
