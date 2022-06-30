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

	"ebay_dapp_golang/api/routers"
	"ebay_dapp_golang/config"

	"github.com/gin-gonic/gin"
)

func init() {
	cfg := flag.String("c", "cfg.yml", "configuration file")
	flag.Parse()
	handleConfig(*cfg)
}

//handleConfig
func handleConfig(configFile string) {
	if err := config.Parse(configFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

//Serve
func Serve() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()
	g = routers.Load(g)

	srv := &http.Server{
		Addr:    config.G.Serve.Listen,
		Handler: g,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func main() {
	Serve()
}
