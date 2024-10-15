package database

import "log"

func MigrateDB() {
	err := DB.AutoMigrate(
		&TotalTime{},
		&WorkTime{},
		&BreakTime{},
		&Project{},
		&Task{},
		&Cost{},
		&Brb{},
		&WorkTimeProject{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	log.Println("Database migration completed successfully.")
}
