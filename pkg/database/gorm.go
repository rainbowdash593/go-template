package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	_maxIdleConns    = 10
	_maxOpenConns    = 100
	_connMaxLifetime = time.Hour
)

type Database struct {
	*gorm.DB
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
}

func New(dsn string, opts ...Option) (*Database, error) {
	database := &Database{
		maxIdleConns:    _maxIdleConns,
		maxOpenConns:    _maxOpenConns,
		connMaxLifetime: _connMaxLifetime,
	}

	for _, opt := range opts {
		opt(database)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(database.maxIdleConns)
	sqlDB.SetMaxOpenConns(database.maxOpenConns)
	sqlDB.SetConnMaxLifetime(database.connMaxLifetime)

	return &Database{
		DB: db,
	}, nil
}
