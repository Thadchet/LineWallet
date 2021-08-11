package main

import (
	"flag"
	"fmt"
	"line-wallet/config"
	"line-wallet/routers"
	"os"

	"github.com/labstack/gommon/log"
)

func main() {

	environment := flag.String("env", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	cv := config.Config{}
	if err := cv.InitAllConfiguration(*environment); err != nil {
		log.Errorf("init all configuration %v error", err)
		return
	}

	routers.Init(cv)
}
