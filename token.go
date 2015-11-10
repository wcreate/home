package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

func main() {

	log.Debug("Starting server...")

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(rest.AuthMiddleWare())
	m.Get("/", func(c *macaron.Context, as rest.AuthService) string {
		token, _ := as.GenSysToken(c.RemoteAddr(), 15)
		return token
	})

	m.Run("0.0.0.0", 8081)
}
