package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eltonjr/tmdb-upcoming/server/env"
	"github.com/eltonjr/tmdb-upcoming/server/genres"
)

const (
	detailsPath  = "/movie/"
	upcomingPath = "/movie/upcoming"
	searchPath   = "/search/movie"
)

type (
	response struct {
		results []movie
	}

	movie struct {
		id           int
		title        string
		poster_path  string
		genre_ids    []int
		release_date string
		overview     string
	}
)

func getMovies(name, page string) (Collection, error) {
	path := upcomingPath
	if name != "" {
		path = searchPath
	}

	url := fmt.Sprintf("%s%s", env.Vars.TMDB.Address, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Collection{}, fmt.Errorf("failed to build request to %s: %v", url, err)
	}

	q := req.URL.Query()
	q.Add("api_key", env.Vars.TMDB.Key)
	q.Add("query", name)
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Collection{}, fmt.Errorf("failed to get movies from %s: %v", url, err)
	}

	defer res.Body.Close()

	var ms response
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return Collection{}, fmt.Errorf("failed to read movies response: %v", err)
	}

	var c Collection
	for _, m := range ms.results {
		c.Movies = append(c.Movies, m.toMovieOutput())
	}

	return c, nil
}

func getMovie(id int) (Movie, error) {
	url := fmt.Sprintf("%s%s%d", env.Vars.TMDB.Address, detailsPath, id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to build request to %s: %v", url, err)
	}

	q := req.URL.Query()
	q.Add("api_key", env.Vars.TMDB.Key)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to get movie from %s: %v", url, err)
	}

	defer res.Body.Close()

	var m movie
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to read movie response: %v", err)
	}

	return m.toMovieOutput(), nil
}

func (m movie) toMovieOutput() Movie {
	r, err := time.Parse("2006-01-02", m.release_date)
	var release *time.Time
	if err == nil {
		release = &r
	}

	return Movie{
		ID:          m.id,
		Name:        m.title,
		Poster:      m.poster_path,
		Genre:       genres.Get(m.genre_ids[0]),
		ReleaseDate: release,
		Overview:    m.overview,
	}
}
