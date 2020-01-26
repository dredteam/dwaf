package main

import (
	"gitea.dredteam.com/dwaf/dwaf/internal/dwafconfig"
	"gitea.dredteam.com/dwaf/dwaf/internal/dwafproxy"
)

func main() {
	config := dwafconfig.GetConfiguration()
	config.ReverseProxy.Server.Handler = dwafproxy.New(config.ReverseProxy.URL)

	err := config.ReverseProxy.Server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
