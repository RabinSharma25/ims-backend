package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	configFilePath := flag.String("config", "../../../conf", "Path to config file")
	flag.Parse()
	fmt.Println("Config file path:", *configFilePath)

	viper.SetConfigName("config")
	viper.AddConfigPath(*configFilePath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	logrus.Info("Config file loaded successfully")

	// PostgreSQL connection settings
	logrus.Info("db.GetDatabaseInstance")
	env := viper.GetString("env") + "."
	user := viper.GetString(env + "db.user")
	password := viper.GetString(env + "db.password")

	host := viper.GetString(env + "db.host")
	port := viper.GetString(env + "db.port")
	dbname := viper.GetString(env + "db.name")
	schema := viper.GetString(env + "db.schema")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, user, password, dbname, port, schema)

	// Connect to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Read the SQL schema from file
	schemaFile := "schema.sql"
	schemaContent, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	// Split SQL content into individual statements
	sqlStatements := strings.Split(string(schemaContent), ";")
	for _, statement := range sqlStatements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue // Skip empty statements
		}

		// Execute the statement using GORM
		err := db.Exec(statement).Error
		if err != nil {
			log.Fatalf("Failed to execute statement: %v\nSQL: %s", err, statement)
		}
		log.Printf("Executed: %s", statement)
	}

	log.Println("PostgreSQL schema migration completed successfully!")
}
