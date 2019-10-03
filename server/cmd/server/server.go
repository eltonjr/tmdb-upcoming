package main

import (
	"log"
	"net/http"

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
	if err := http.ListenAndServe(env.Vars.Server.Address, r); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
