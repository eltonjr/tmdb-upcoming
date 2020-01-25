package health

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/tmdb-upcoming/server/env"
	"github.com/eltonjr/tmdb-upcoming/server/errors"
)

const configurationPath = "/configuration"

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
	url := fmt.Sprintf("%s%s", env.Vars.TMDB.Address, configurationPath)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("dependencies misconfigured: %v", err)
	}

	q := req.URL.Query()
	q.Add("api_key", env.Vars.TMDB.Key)
	req.URL.RawQuery = q.Encode()

	_, err = http.DefaultClient.Do(req)

	return err
}
