package health

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/tmdb-upcoming/server/errors"
)

// Check only returns a valid 200 to tell the service is up and running
func Check(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.URL.Query().Get("dependencies") == "" {
		err := checkDependencies()
		if err != nil {
			errors.Respond(w, http.StatusInternalServerError, "unhealthy dependencies", "System dependency 'The Movies DataBase' is unreachable. Please try again later")
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func checkDependencies() error {
	// TODO check dependencies
	return nil
}
