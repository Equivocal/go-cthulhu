package cthulhu

import (
	"net/http"

	log "github.com/amedia/go-logstash"
)

type Resource struct {
	path    string
	handler http.Handler
}

type Cthulhu struct {
	resources map[string]http.Handler
}

func (c *Cthulhu) RegisterResource(path string, handler http.Handler) {
	if c.resources == nil {
		c.resources = make(map[string]http.Handler)
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
