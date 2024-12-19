package db

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewPostgresConnection(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	m, err := migrate.New("file://internal/app/db/migrations", dsn)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}
	return dbPool, nil
}

func NewMongoConnection(dsn string) (*mongo.Client, error) {
	client := options.Client()
	client.ApplyURI(dsn)
	mongoConn, err := mongo.Connect(context.Background(), client)
	if err != nil {
		return nil, err
	}
	return mongoConn, nil
}
