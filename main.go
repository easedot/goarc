package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	router.NewRouter(e, r.NewAppController())
	router.NewChatServer(e, rds)
	//router.NewWebSocket(e)
	//router.NewSocketIO(e)

	//for standlone server
	//log.Fatal(e.Start(config.C.Server.Address))

	go e.Start(config.C.Server.Address)
	webview.Open("stock demo", "http://localhost:8080/chat.html", 1024, 768, true)

}
