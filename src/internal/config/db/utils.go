package db

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var migrationsPath = flag.String("m", "../migrations", "the migrations folder")
var seederShPath = flag.String("seeders", "../run-seed.sh", "the seeders script")

func RunMigrations() {
	dbUri := os.Getenv("DATABASE_URL")

	absMigrationsPath, err := filepath.Abs(*migrationsPath)
	fmt.Println(absMigrationsPath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s", absMigrationsPath),
		dbUri,
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration successful")
}

func RunSeederScript() {
	absSeedPath, err := filepath.Abs(*seederShPath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	cmd := exec.Command("/bin/sh", absSeedPath, "dev")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Printf("Error running seeder script: %v", err)
	} else {
		log.Println("Seeder script completed successfully")
	}
}
