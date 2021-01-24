package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/savsgio/atreugo/v10"

	"github.com/easedot/hb_vendor/config"
	"github.com/easedot/hb_vendor/devices/datastore"
	"github.com/easedot/hb_vendor/devices/router"
	"github.com/easedot/hb_vendor/registry"
)

// jwtCustomClaims are custom claims extending default ones.

func main() {
	config.InitConfig()
	db := datastore.NewDB()

	defer db.Close()

	r := registry.NewRegistry(db)

	//for echo server
	e := echo.New()
	router.NewERouter(e, r.NewAppController())
	go e.Start(config.C.Server.Address)

	//webview.Open("stock demo", "http://0.0.0.0:8080/assets/chat.html", 1024, 768, true)

	//only server
	log.Fatal(e.Start(config.C.Server.Address))

}
