package main

import (
	"fastx-api/user-service/config"
	seeders2 "fastx-api/user-service/database/seeders"
	"fastx-api/user-service/src/pkg/database"
	"gorm.io/gorm"
	"log"
)

func runSeeders(db *gorm.DB, seeders []seeders2.Seeder) {
	for _, seeder := range seeders {
		log.Printf("Seeding %s...\n", seeder.TableName())
		if err := seeder.Seed(db); err != nil {
			log.Fatalf("Failed to seed %s: %v", seeder.TableName(), err)
		}
	}
}

func main() {
	config.LoadConfig()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	seeders := []seeders2.Seeder{
		&seeders2.UserSeeder{},
	}

	runSeeders(db, seeders)
}
