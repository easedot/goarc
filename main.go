package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"github.com/easedot/goarc/config"
	"github.com/easedot/goarc/drivers/websocket"
	"github.com/easedot/goarc/registry"
	"github.com/easedot/goarc/drivers/datastore"
	"github.com/easedot/goarc/drivers/router"
)

func main() {

	db:=datastore.NewDB()
	defer db.Close()

	r:=registry.NewRegistry(db)

	e := echo.New()
	router.NewRouter(e,r.NewAppController())
	websocket.NewWebSocket(e)

	log.Fatal(e.Start(config.C.Server.Address))

}
