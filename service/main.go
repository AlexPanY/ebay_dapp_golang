package main

import (
	"fmt"

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

func main() {
	api.Serve()
}
