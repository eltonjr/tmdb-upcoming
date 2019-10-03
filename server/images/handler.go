package images

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/tmdb-upcoming/server/env"
	"github.com/eltonjr/tmdb-upcoming/server/errors"
)

// GetImage will get an image from TMDB and pipe it to the caller
func GetImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	url := fmt.Sprintf("%s/%s/%s", env.Vars.TMDB.Images.BasePath, env.Vars.TMDB.Images.PosterSize, id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		errors.Respond(w, http.StatusInternalServerError, "unable to get image", err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errors.Respond(w, http.StatusInternalServerError, "unable to get image", err.Error())
		return
	}

	defer res.Body.Close()
	io.Copy(w, res.Body)

	w.WriteHeader(res.StatusCode)
}
