package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/tmdb-upcoming/server/errors"
)

// GetAll is the handler for /movies.
// It returns a list of movies given some limit and some offset.
// Limit default is 20. Offset default is 0.
// It's also possible to filter movies by name
func GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := r.URL.Query()
	page := q.Get("page")
	name := q.Get("name")

	movies, err := getMovies(name, page)
	if err != nil {
		errors.Respond(w, http.StatusInternalServerError, "unable to get movies", err.Error())
		return
	}

	ms, err := json.Marshal(movies)
	if err != nil {
		errors.Respond(w, http.StatusInternalServerError, "invalid payload", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", ms)
}

// GetOne is a handler for /movies/{id}.
// It returns a Movie given a valid ID
func GetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_id := p.ByName("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		errors.Respond(w, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	movie, err := getMovie(id)
	if err != nil {
		errors.Respond(w, http.StatusInternalServerError, "unable to get movie", err.Error())
		return
	}

	m, err := json.Marshal(movie)
	if err != nil {
		errors.Respond(w, http.StatusInternalServerError, "invalid payload", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", m)
}
