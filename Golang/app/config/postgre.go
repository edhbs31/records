package config

import (
	"log"
	"os"
	"time"

	"record/app/domains"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ScoopConnection struct {
	Dsn string `json:"dsn"`
}

func (s *ScoopConnection) SetDsn(dsn string) {
	s.Dsn = dsn
}

func (s ScoopConnection) GetConnections() *gorm.DB {
	db, err := gorm.Open(postgres.Open(s.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		log.Fatalln("error Connection DB", os.Getenv("DB_HOST"))
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	e, _ := db.DB()
	e.SetMaxIdleConns(domains.MaxIddleConnection)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	e.SetMaxOpenConns(domains.MaxOpenConnection)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	e.SetConnMaxLifetime(time.Hour)
	return db
}

type Db struct {
}

func (d *Db) Build() Connection {
	return &ScoopConnection{}
}
