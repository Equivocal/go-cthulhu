package cthulhu

import "net/http"

func Example(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("ExampleHeader", "true")
		h.ServeHTTP(w, r)
	})
}

func Use(h http.Handler, middlewares ...http.Handler) {}
