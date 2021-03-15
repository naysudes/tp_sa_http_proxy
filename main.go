package main

import (
	"fmt"
	"log"

	proxy "github.com/naysudes/tp_sa_http_proxy/proxy"
)

func main() {
	app, err := proxy.NewServer()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err.Error()))
	}
	app.Run()
}
