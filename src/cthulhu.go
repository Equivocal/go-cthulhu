package cthulhu

import (
	"net/http"

	log "github.com/amedia/go-logstash"
)

type Resource struct {
	path    string
	handler RequestHandler
}

type Cthulhu struct {
	resources map[string]RequestHandler
}

func (c *Cthulhu) RegisterResource(path string, handler RequestHandler) {
	if c.resources == nil {
		c.resources = make(map[string]RequestHandler)
	}
	c.resources[path] = handler
}

func (c *Cthulhu) Start() {
	log.Info("Initializing application..")
	for k, v := range c.resources {
		log.Infof("Registered resource '%s'", k)
		http.Handle(k, v)
	}
}
