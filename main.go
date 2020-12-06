package main

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/router"
)

func main() {
	debug := true
	r := router.Router()

	if debug {
		r.Run()
	} else {
		r.RunTLS(":8080", "cert.pem", "privkey.pem")
	}

}
