package images

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eltonjr/tmdb-upcoming/server/env"
)

const configurationPath = "/configuration"

type response struct {
	images struct {
		base_url     string
		poster_sizes []string
	}
}

func init() {
	url := fmt.Sprintf("%s%s", env.Vars.TMDB.Address, configurationPath)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("failed to build request to %s: %v", url, err)
	}

	q := req.URL.Query()
	q.Add("api_key", env.Vars.TMDB.Key)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to get configuration from %s: %v", url, err)
	}

	defer res.Body.Close()

	var c response
	err = json.NewDecoder(res.Body).Decode(&c)
	if err != nil {
		log.Fatalf("failed to read configuration response: %v", err)
	}

	env.Vars.TMDB.Images.BasePath = c.images.base_url
	env.Vars.TMDB.Images.PosterSize = c.images.poster_sizes[1]
}
