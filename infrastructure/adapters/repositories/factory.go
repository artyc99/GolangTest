package repositories

import (
	"EchoAPI/core/repositories"
	"context"
	"gorm.io/gorm"
)

type GormDB interface {
	WithContext(ctx context.Context) *gorm.DB
}

type factory struct {
	PostgresDB GormDB
}

func NewFactory(postgresDB GormDB) repositories.RepoFactoryI {
	return &factory{
		PostgresDB: postgresDB,
	}
}

func (f *factory) Users() repositories.UsersI {
	return Users(f.PostgresDB)
}
