package bootstrap

import (
	"log"

	"github.com/41x3n/TeleUtil/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) *gorm.DB {
	dsn := env.DSN

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connection to PostgreSQL established.")

	return client
}

func AutoMigrate(client *gorm.DB) {
	err := client.AutoMigrate(&domain.User{}, &domain.Photo{})
	if err != nil {
		panic("Failed to auto migrate: " + err.Error())
	}

	log.Println("Auto migration completed.")
}

func ClosePostgresDBConnection(client *gorm.DB) {
	sqlDB, err := client.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB: ", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Failed to close database connection: ", err)
	}

	log.Println("Connection to PostgreSQL closed.")
}
