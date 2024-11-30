

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Utility functions
func ToSnakeCase(input string) string {
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	snake := re.ReplaceAllString(input, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToPascalCase(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// Read JSON file and unmarshal data
func readJSONFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}
	return data, nil
}

// Insert data into the database for a specific table
func insertDataIntoTable(db *gorm.DB, tableName string, modelSlice interface{}) error {
	if err := db.Table(tableName).Create(modelSlice).Error; err != nil {
		return fmt.Errorf("failed to insert data into table %s: %w", tableName, err)
	}
	return nil
}

// Seed a specific table dynamically
func seedTable(db *gorm.DB, fileName string, pascalCaseName string, snakeCaseName string) error {
	// Read the JSON file
	data, err := readJSONFile(fileName)
	if err != nil {
		return err
	}

	// Find the corresponding model
	model, exists := ModelRegistry[pascalCaseName]
	if !exists {
		return fmt.Errorf("no model found for %s, skipping", pascalCaseName)
	}

	// Create a slice of the model type
	modelSlice := reflect.New(reflect.SliceOf(reflect.TypeOf(model).Elem())).Interface()

	// Unmarshal the JSON data into the slice
	if err := json.Unmarshal(data, modelSlice); err != nil {
		return fmt.Errorf("failed to unmarshal JSON for %s: %w", fileName, err)
	}

	// Insert the data into the table
	if err := insertDataIntoTable(db, snakeCaseName, modelSlice); err != nil {
		return err
	}

	log.Printf("Successfully seeded table: %s", snakeCaseName)
	return nil
}

// Seed the entire database
func SeedDatabase(db *gorm.DB) error {
	// Read the JSON files
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	prefixRegex := regexp.MustCompile(`^\d+_`)

	// Process each JSON file
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			// Extract names
			originalName := file.Name()
			cleanedName := prefixRegex.ReplaceAllString(originalName, "")
			camelCaseName := strings.TrimSuffix(cleanedName, ".json")
			snakeCaseName := ToSnakeCase(camelCaseName)
			pascalCaseName := ToPascalCase(snakeCaseName)

			// Seed the corresponding table
			if err := seedTable(db, file.Name(), pascalCaseName, snakeCaseName); err != nil {
				log.Printf("Error seeding table %s: %v", snakeCaseName, err)
			}
		}
	}
	return nil
}

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

	// Seed database
	if err := SeedDatabase(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	fmt.Println("Database seeding completed!")
}
