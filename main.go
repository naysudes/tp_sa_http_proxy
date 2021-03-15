package main

import (
	"fmt"
	proxy "https://github.com/naysudes/tp_sa_http_proxy/proxy"
)

func main() {
	app, err := proxy.NewServer(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err.Error()))
	}
	defer app.Close()
	app.Run()
}
