package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/andres06-hub/loyalty-service/src/internal/config"
	"github.com/andres06-hub/loyalty-service/src/internal/config/db"
	"github.com/andres06-hub/loyalty-service/src/internal/handler"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/definition.yaml", "the config file")
var migrationsPath = flag.String("migrations", "../migrations", "the migrations folder")

func runMigrations() {
	dbUri := os.Getenv("DATABASE_URL")

	absMigrationsPath, err := filepath.Abs(*migrationsPath)
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

func main() {
	flag.Parse()

	if true {
		runMigrations()
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	conn, err := db.GetConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	ctx := svc.NewServiceContext(c, conn)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d\n", c.Host, c.Port)
	server.Start()
}
