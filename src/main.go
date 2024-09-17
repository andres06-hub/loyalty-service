package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/andres06-hub/loyalty-service/src/docs"
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	"github.com/andres06-hub/loyalty-service/src/internal/config/db"
	"github.com/andres06-hub/loyalty-service/src/internal/handler"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/definition.yaml", "the config file")

// @title			Loyalty Service API
// @version		1.0
// @description	This is a loyalty service server.
// @host			localhost:8080
// @BasePath		/api
func main() {
	flag.Parse()

	if true {
		db.RunMigrations()
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

	go func() {
		fmt.Printf("Starting server at %s:%d\n", c.Host, c.Port)
		server.Start()
	}()

	time.Sleep(2 * time.Second)

	select {}
}
