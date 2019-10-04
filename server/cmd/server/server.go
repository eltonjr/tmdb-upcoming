package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/tmdb-upcoming/server/env"
	"github.com/eltonjr/tmdb-upcoming/server/health"
	"github.com/eltonjr/tmdb-upcoming/server/images"
	"github.com/eltonjr/tmdb-upcoming/server/movies"
)

func main() {
	r := httprouter.New()

	r.GET("/v1/movies", movies.GetAll)
	r.GET("/v1/movies/:id", movies.GetOne)

	r.GET("/v1/posters/:id", images.GetImage)

	r.GET("/v1/health", health.Check)

	log.Printf("Server is running at %s", env.Vars.Server.Address)
	if err := http.ListenAndServe(env.Vars.Server.Address, middleware(r)); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

// WrapperResponseWriter is needed because http.ResponseWriter
// does not expose the returned status code.
// To log the returned status code, a new implementation of the
// ResponseWriter interface is needed to override the WriteHeader
// method and store the status code.
type wrapperRW struct {
	http.ResponseWriter
	status int
}

func (w *wrapperRW) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		ww := wrapperRW{ResponseWriter: w}
		next.ServeHTTP(&ww, r)
		t2 := time.Now()
		log.Printf("%s %s - %d %v", r.Method, r.URL, ww.status, t2.Sub(t1))
	})
}
