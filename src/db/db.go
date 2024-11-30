package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dba *gorm.DB

func GetDatabaseInstance() (db *gorm.DB) {
	logrus.Info("db.GetDatabaseInstance")
	env := viper.GetString("env") + "."
	user := viper.GetString(env + "db.user")
	password := viper.GetString(env + "db.password")

	host := viper.GetString(env + "db.host")
	port := viper.GetString(env + "db.port")
	dbname := viper.GetString(env + "db.name")
	schema := viper.GetString(env + "db.schema")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, user, password, dbname, port, schema)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	dba = db
	if err != nil {
		logrus.Panic(fmt.Sprintf("Error connecting to the database at %s:%s/%s", host, port, dbname))
	}
	sqlDB, err := dba.DB()
	if err != nil {
		logrus.Panic("Error getting GORM DB definition")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	logrus.Info(fmt.Sprintf("Successfully established connection to %s:%s/%s", host, port, dbname))

	return dba
}
