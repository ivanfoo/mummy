package main

import (
	"flag"

	"github.com/ivanfoo/mummy/services"

	"github.com/gin-gonic/gin"
)

var (
	listenAddr = flag.String("listen", ":443", "listen address")
	certFile   = flag.String("cert", "", "TLS certificate file path")
	keyFile    = flag.String("key", "", "TLS key file path")
)

func main() {
	router := gin.New()

	router.GET("/status", services.GetAllStatus)

	if *certPem == "" && *keyPerm == "" {
		router.Run(*listenAddr)
	} else {
		router.RunTLS(*listenAddr, *certPem, *keyPerm)
	}
}
