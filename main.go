package main

import (
	"log"

	"freddy.facuellarg.com/adapters/router"
	"freddy.facuellarg.com/config"
	"freddy.facuellarg.com/domain/connection"
	"freddy.facuellarg.com/domain/usecase"
	"freddy.facuellarg.com/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var databaseFields = connection.DataBaseConnection{
	"root",
	"root1234",
	"localhost",
	3306,
	"vim",
}

func main() {
	if err := config.ReadConf(); err != nil {
		log.Fatal(err)
	}
	if err := mapstructure.Decode(
		viper.GetStringMap("database"),
		&databaseFields,
	); err != nil {
		log.Fatal(err)
	}
	//get the db connection
	db, err := connection.DBConnection(databaseFields)
	if err != nil {
		log.Fatalf("in Connection %s", err)
	}

	//database migrations, if debug, delete and create the database
	if utils.Debug() {
		err = MigrateRestart(db)
	} else {
		err = MigrateUp(db)
	}
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("in migration ", err)
	}

	// repositories
	gormDB, err := connection.GormDB(db)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := usecase.NewUserRepository(gormDB)

	server := echo.New()

	//router
	router := router.NewRouter(server, userRepo)
	if err := router.RouteUsers(); err != nil {
		log.Fatal(err)
	}

	server.Use(middleware.Logger())
	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

//Start start a applicatopm
func Start() {

	return
}
