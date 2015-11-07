package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc"
	"gopkg.in/macaron.v1"
)

func main() {

	log.Debug("Starting server...")
	tkits.SyncDB()

	cfg := macaron.Config()
	web, err := cfg.GetSection("web")
	if err != nil {
		panic(err)
	}
	ip := web.Key("ip").MustString("0.0.0.0")
	port := web.Key("port").MustInt(8080)

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	// user
	wuc.InitHandles(m)

	m.Run(ip, port)
}
