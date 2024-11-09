package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/md-muzzammil-rashid/traffic-flow/config"
)

func main() {
	cfg, err:= config.InitConfig()
	if err!= nil {
        fmt.Println("Error initializing config:", err)
        return
    }

	router := mux.NewRouter()

	server := http.Server{
		Handler: router,	
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func ()  {
		slog.Info("Server is running on PORT: ",cfg.Port)
		err = server.ListenAndServe(); if err != nil {
			log.Fatalf("Error: failed to start the server : %s", err.Error())
		}
	}()

	<- done

	log.Println("Shutting down the server...")
	
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*30)

	defer cancle()

	err = server.Shutdown(ctx); if err != nil {
		log.Fatalf("Error: failed to gracefully shutdown the server : %s", err.Error())
	}

	slog.Info("Server shutdown completed")




}