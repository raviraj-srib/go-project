package router

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/raviraj-srib/go-project/corporate-directory/src/logger"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//Decorator for printing metadata
		handler = logRoute(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

func logRoute(innerHandler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		innerHandler.ServeHTTP(w, r)

		logger.Debug("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
