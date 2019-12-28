package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"github.com/easedot/goarc/config"
	"github.com/easedot/goarc/registry"
	"github.com/easedot/goarc/tools/datastore"
	"github.com/easedot/goarc/tools/router"
)

func main() {

	db:=datastore.NewDB()
	defer db.Close()

	r:=registry.NewRegistry(db)

	e := echo.New()
	router.NewRouter(e,r.NewAppController())

	log.Fatal(e.Start(config.C.Server.Address))

}
