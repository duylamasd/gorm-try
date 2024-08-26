package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/duylamasd/gorm-try/config"
	"github.com/duylamasd/gorm-try/controllers"
	"github.com/duylamasd/gorm-try/repositories"
	"github.com/duylamasd/gorm-try/routes"
	"github.com/duylamasd/gorm-try/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	container := InitContainer()
	engine := InitGinEngine(container)

	port := os.Getenv("PORT")
	server := &http.Server{Addr: ":" + port, Handler: engine}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signal.Notify(quit, signals...)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func InitContainer() *dig.Container {
	container := dig.New()

	err := config.InjectDatabase(container)
	if err != nil {
		log.Fatal("Error injecting database")
	}

	err = repositories.Inject(container)
	if err != nil {
		log.Fatal("Error injecting repositories")
	}

	err = services.Inject(container)
	if err != nil {
		log.Fatal("Error injecting services")
	}

	err = controllers.Inject(container)
	if err != nil {
		log.Fatal("Error injecting controllers")
	}

	return container
}

func InitGinEngine(container *dig.Container) *gin.Engine {
	engine := gin.New()

	err := routes.Init(engine, container)
	if err != nil {
		log.Fatal("Error initializing routes: ", err)
	}

	return engine
}
