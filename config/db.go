package config

import (
	"log"
	"os"

	"github.com/duylamasd/gorm-try/interfaces"
	"go.uber.org/dig"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d *Database) GetInstance() *gorm.DB {
	return d.db
}

func NewDatabase() interfaces.Database {
	dsn := os.Getenv("DB_DSN")
	dialector := postgres.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return &Database{db: db}
}

func InjectDatabase(container *dig.Container) error {
	_ = container.Provide(NewDatabase)

	return nil
}
