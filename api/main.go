package main

import (
	"bytes"
	"log"
	"os"

	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/db/mock"
	"github.com/dgkg/goworkshop/api/db/postgres"
	"github.com/dgkg/goworkshop/api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	Production = "prod"
	Test       = "test"
)

var yamlExample = []byte(`
postgres:
  port: 5432
  host:     localhost
  dbname:   gorm
  user:     gorm
  password: mypassword
`)

func init() {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	env := os.Getenv("APP_VOTE_ENV")
	log.Println("run app in", env, "env")
	var db db.DB

	if env == Production {
		c := &postgres.DBCredential{
			Port:     viper.GetString("postgres.port"),
			Host:     viper.GetString("postgres.host"),
			DBname:   viper.GetString("postgres.dbname"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
		}
		log.Println("use postgres connection for DB")
		db = postgres.New(c)
	} else {
		log.Println("use mock connection for DB")
		db = mock.New()
	}
	service.New(r, db)
	r.Run(":8080")
}
