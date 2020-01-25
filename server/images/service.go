package images

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eltonjr/tmdb-upcoming/server/env"
)

const configurationPath = "/configuration"

// Response is the mapped schema from the configuration endpoint
type Response struct {
	Images struct {
		BaseURL     string   `json:"base_url"`
		PosterSizes []string `json:"poster_sizes"`
	} `json:"images"`
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

	var c Response
	err = json.NewDecoder(res.Body).Decode(&c)
	if err != nil {
		log.Fatalf("failed to read configuration response: %v", err)
	}

	env.Vars.TMDB.Images.BasePath = c.Images.BaseURL
	env.Vars.TMDB.Images.PosterSize = c.Images.PosterSizes[3] // TODO make this size dynamic
}
