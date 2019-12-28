package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type appConfig struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Host                 string
		Name                 string
		Port				 string
		AllowNativePasswords bool
		Params               struct {
			ParseTime 	string
			TimeZone 	string
		}
	}
	Server struct {
		Address	string
	}
	Debug bool
}

var C appConfig
func init(){
	env:="dev"
	if set,find:=os.LookupEnv("ENV");find{
		env=strings.ToLower(set)
	}
	configFile :=fmt.Sprintf("config_%s.yml",env)
	//viper.SetConfigFile(configFile)
	viper.SetConfigName(configFile)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	//viper.AutomaticEnv()

	if err:=viper.ReadInConfig();err!=nil{
		fmt.Println(err)
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
