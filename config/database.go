package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUser  = "uniseruapp"
	DBPass  = "uniseruapp~!@#"
	DBHost  = "13.229.47.238"
	DBName  = "postgres"
	DBPort  = "5432"
	SSLMode = "disable"
)

func ConnectPGLocal() *gorm.DB {
	// dsn
	dsn := fmt.Sprintf(`
		host=%s user=%s password=%s dbname=%s port=%s sslmode=%s`,
		DBHost,
		DBUser,
		DBPass,
		DBName,
		DBPort,
		SSLMode,
	)

	log.Print("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Warn("Connected to database local Failed:", err)
	}
	log.Warn("Connected to database local")

	return db
}
