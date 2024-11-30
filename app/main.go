// this is our entry point
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/middlewares"
	"github.com/rabinsharma25/ims-backend/src/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Read config file
func loadConfig() {
	configFilePath := flag.String("config", "../conf", "Path to config file")
	flag.Parse()
	fmt.Println("Config file path:", *configFilePath)

	viper.SetConfigName("config")
	viper.AddConfigPath(*configFilePath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	logrus.Info("Config file loaded successfully")
}

// Start Gin server
func startServer(engine *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	go func() {
		logrus.Info("Starting server on port ", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("Shutting down server...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	logrus.Info("Server exiting")
}

// init is invoked before main()
func init() {
	loadConfig()
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())

	// Environment Setup
	env := viper.GetString("env") // dev, qa, prod
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Server Context Path
	logrus.Info("Server context path ", viper.GetString(env+".server.contextPath"))

	// Apply Global Middlewares
	middlewares.ApplyMiddlewares(engine)
	// Middlewares can be applied to specific routes as well use the routes package to do so

	// Setup Routes
	routes.SetupRoutes(engine)

	startServer(engine, viper.GetString(env+".server.port"))
}

/*

taking a copy of the database form the authnull_dev user to my local postgres user

// to backup
pg_dump -h 65.109.82.187 -p 2514 -U kloudone -d authnull_dev -F t -f authnull_dev_backup.tar

// to restore
pg_restore -h localhost -p 5432 -U postgres -d authnull_dev -W -F t authnull_prod_backup.tar
*/
