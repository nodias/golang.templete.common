package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nodias/golang.templete.common/models"
	"github.com/nodias/golang.templete.common/shared/logger"
	"github.com/sirupsen/logrus"
)

var config models.TomlConfig
var log *logrus.Logger

func Init() {
	config = *models.GetConfig()
	log = logger.New(context.Background())
}

const (
	DatabaseUser     = "admin"
	DatabasePassword = "admin"
	DatabaseName     = "postgres"
)

type DataAccess interface {
	Get(id string) (*models.User, error)
}

func NewOpenDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		DatabaseUser,
		DatabasePassword,
		DatabaseName,
		config.Databases["postgres"].Server,
		config.Databases["postgres"].Port,
	)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
		panic("Invalid DB config")
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		panic("DB unreachable")
	}
	log.Debug("connected DB")
	return db
}
