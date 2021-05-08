package datastore

import (
	"log"

	"vote-app-api/domain/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

type MysqlConfig struct {
	User     string `envconfig:"MYSQL_USER"`
	Password string `envconfig:"MYSQL_PASSWORD"`
	Host     string `envconfig:"MYSQL_HOST"`
	DB       string `envconfig:"MYSQL_DATABASE"`
}

func NewMysqlDB() *gorm.DB {
	var mysqlConfig MysqlConfig
	envconfig.Process("MYSQL", &mysqlConfig)
	dbInfo := mysqlConfig.User + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Host + ")/"
	dbInfoWithDatabase := dbInfo + mysqlConfig.DB + "?charset=utf8mb4&parseTime=True&loc=Local"

	conn, err := gorm.Open("mysql", dbInfoWithDatabase)
	// This is settings for production.
	// Create database if not exists.
	if err != nil {
		log.Printf("database connection error: %+v\n", err.Error())
		conn, _ := gorm.Open("mysql", dbInfo)
		_ = conn.Exec("CREATE DATABASE IF NOT EXISTS " + mysqlConfig.DB + " DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
		if conn.Error != nil {
			log.Printf("database creation error: %+v\n", err.Error())
			conn, err = gorm.Open("mysql", dbInfo)
			if err != nil {
				log.Printf("database connection error: %+v\n", err.Error())
				panic(err)
			}
		}
	}

	err = conn.DB().Ping()
	if nil != err {
		log.Printf("ping error is %+v\n", err)
		// Reconnect to database.
		conn, err = gorm.Open("mysql", dbInfoWithDatabase)
		if err != nil {
			panic(err)
		}
	}
	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	conn.AutoMigrate(
		&model.MasterMovie{},
		&model.MasterPerson{},
		&model.Vote{},
	)

	conn.DB().SetMaxOpenConns(20)
	return conn
}
