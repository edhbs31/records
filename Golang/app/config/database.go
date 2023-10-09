package config

import (
	"gorm.io/gorm"
)

type Connection interface {
	GetConnections() *gorm.DB
	SetDsn(dsn string)
}

type ConnectionDB interface {
	Build() Connection
}
