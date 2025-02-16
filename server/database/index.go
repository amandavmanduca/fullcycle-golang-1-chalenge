package database

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	Connect(ctx context.Context) (*gorm.DB, error)
	GetConnection(ctx context.Context) (*gorm.DB, error)
	AutoMigrate(ctx context.Context, dst ...interface{}) error
}

type SqliteImpl struct {
	Dsn        string
	connection *gorm.DB
}

func NewSqlite(config SqliteImpl) Database {
	return &config
}

func (s *SqliteImpl) Connect(ctx context.Context) (*gorm.DB, error) {
	if s.connection != nil {
		return s.connection, nil
	}
	db, err := gorm.Open(sqlite.Open(s.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	s.connection = db
	return db.WithContext(ctx), nil
}

func (s SqliteImpl) GetConnection(ctx context.Context) (*gorm.DB, error) {
	if s.connection == nil {
		return s.Connect(ctx)
	}
	return s.connection, nil
}

func (s *SqliteImpl) AutoMigrate(ctx context.Context, dst ...interface{}) error {
	s.connection.AutoMigrate(dst...)

	return nil
}
