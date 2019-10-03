package genres

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eltonjr/tmdb-upcoming/server/env"
)

type (
	response struct {
		genres []genre
	}

	genre struct {
		id   int
		name string
	}
)

const genresPath = "/genre/movie/list"

func init() {
	url := fmt.Sprintf("%s%s", env.Vars.TMDB.Address, genresPath)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("failed to build request to %s: %v", url, err)
	}

	q := req.URL.Query()
	q.Add("api_key", env.Vars.TMDB.Key)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to get genres from %s: %v", url, err)
	}

	defer res.Body.Close()

	var gs response
	err = json.NewDecoder(res.Body).Decode(&gs)
	if err != nil {
		log.Fatalf("failed to read genres response: %v", err)
	}

	for _, g := range gs.genres {
		collection[g.id] = g.name
	}
}
