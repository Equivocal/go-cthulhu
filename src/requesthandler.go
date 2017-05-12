package cthulhu

import "net/http"

type RequestHandler func(w http.ResponseWriter, r *http.Request)

func (f RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func Example(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("ExampleHeader", "true")
		h.ServeHTTP(w, r)
	})
}

func Use(h http.Handler, middlewares ...http.Handler) {}
