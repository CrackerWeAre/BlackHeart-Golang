package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/ssoyyoung.p/BlackHeart-Golang/router"
)

func main() {
	debug := true
	r := router.Router()

	if debug {
		r.Run()
	} else {
		log.Fatal(autotls.Run(r))
	}

}
