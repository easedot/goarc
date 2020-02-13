package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/savsgio/atreugo/v10"
	"github.com/zserge/webview"

	"github.com/easedot/goarc/config"
	"github.com/easedot/goarc/devices/datastore"
	"github.com/easedot/goarc/devices/router"
	"github.com/easedot/goarc/registry"
)

func main() {

	db := datastore.NewDB()
	defer db.Close()

	rds, _ := datastore.NewRedis()
	defer rds.Close()

	r := registry.NewRegistry(db)

	//for fasthttp
	//config := atreugo.Config{
	//	Addr: config.C.Server.Address,
	//}
	//e := atreugo.New(&config)
	//router.NewFRouter(e, r.NewAppController(),rds)
	//go e.ListenAndServe()

	//for echo server
	e := echo.New()
	router.NewERouter(e, r.NewAppController(), rds)
	go e.Start(config.C.Server.Address)

	webview.Open("stock demo", "http://0.0.0.0:8080/assets/chat.html", 1024, 768, true)

	//only server
	//log.Fatal(e.Start(config.C.Server.Address))

}
