package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var DB *sql.DB

func InitDB() {
	// Railway sudah kasih DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("❌ Failed to open DB connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	log.Println("✅ Database connected successfully")

	// Jalankan migration
	RunMigration()
}

func RunMigration() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations", 
	}

	n, err := migrate.Exec(DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Printf("✅ Applied %d migrations!\n", n)
}
