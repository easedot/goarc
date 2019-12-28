package datastore

import (
	"database/sql"
	"net/url"
	"fmt"
	"log"
	"github.com/easedot/godbs"

	"github.com/easedot/goarc/config"
)

func NewDB() *godbs.DbHelper{

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.C.Database.User,
		config.C.Database.Password,
		config.C.Database.Host, config.C.Database.Port,
		config.C.Database.Name,
	)
	val := url.Values{}
	val.Add("parseTime", config.C.Database.Params.ParseTime)
	val.Add("loc", config.C.Database.Params.TimeZone)
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err:= dbConn.Ping();err != nil {
		log.Fatal(err)
	}

	db:=godbs.NewHelper(dbConn,config.C.Debug)
	return &db
}
