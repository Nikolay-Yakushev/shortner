package pkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mango/data"
	"time"
)


var dbConn *gorm.DB

func CreateDBConnection(config *data.Config) error {
	//dsn := "host=localhost user=admin password=docker dbname=test_db port=54320 sslmode=disable"
	dsn := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	config.Postgres.PostgresHost,
	config.Postgres.PostgresUser,
	config.Postgres.PostgresUserPassword,
	config.Postgres.PostgresDatabaseName,
	config.Postgres.PostgresPort,
	)
	var logLevel logger.LogLevel = logger.Info
	if config.AppMode=="prod"{
		logLevel = logger.Silent
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	db.AutoMigrate(&data.HashedData{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error occurred while connecting with the database %s", err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db
	return err
}
func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}
