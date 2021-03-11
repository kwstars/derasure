package main

import (
	"flag"
	"log"
)

var configFile = flag.String("f", "", "set config file which viper will loading.")

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	flag.Parse()

	app, closeFunc, err := CreateApp(*configFile)
	if err != nil {
		if closeFunc != nil {
			closeFunc()
		}
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
