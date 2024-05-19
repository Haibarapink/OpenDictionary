package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Haibarapink/OpenDictionary/handler"
	"github.com/gin-gonic/gin"
)

func handle_account(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"asd": "Asd",
	})
}

func main() {
	log.Println("Server starting...")
	router := gin.Default()
	handler.NewHandler(&handler.Config{
		R: router,
	})
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Printf("Listen to %s", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
