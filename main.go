package main

import (
	//"log"

	//"github.com/gin-gonic/autotls"
	"github.com/ssoyyoung.p/BlackHeart-Golang/router"
)

func main() {
	debug := false
	r := router.Router()

	if debug {
		r.Run()
	} else {
		//log.Fatal(autotls.Run(r, "sparker.kr", "blackheart.sparker.kr"))
		r.RunTLS(":8080", 
		"cert.pem", 
		"privkey.pem")

	}

}
