package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/wcreate/home/tools"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc"
	"gopkg.in/macaron.v1"
)

func main() {

	log.Debug("Starting server...")
	tkits.SyncDB()

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(tools.SpiderFunc())

	// user centre
	wuc.InitHandles(m)

	m.Run(tkits.WebListenIP, tkits.WebPort)
}
