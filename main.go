package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/easedot/godbs"
	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)
var dsn string
var db godbs.DbHelper

func init() {
	env:="dev"
	if set,find:=os.LookupEnv("ENV");find{
		env=strings.ToLower(set)
	}
	config_file:=fmt.Sprintf("config_%s.json",env)
	viper.SetConfigFile(config_file)
	if err:=viper.ReadInConfig();err!=nil{
		panic(err)
	}
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	parsTime:= viper.GetString(`parse_time`)
	timeZone:= viper.GetString(`time_zone`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", parsTime)
	val.Add("loc", timeZone)
	dsn = fmt.Sprintf("%s?%s", connection, val.Encode())

}

func main() {
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//db:=godbs.NewHelper(dbConn,true)
	//
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/articles", func(c echo.Context) error {
		query:="Select * from article "
		list,err:=db.SqlMap(query)
		if err!=nil{
			log.Println(err)
		}
		return c.JSON(http.StatusOK,list)
	})

	log.Fatal(e.Start(viper.GetString("server.address")))

}
