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
	// Response is the mapped schema from the movie endpoint
	Response struct {
		Results []ResultMovie `json:"results"`
	}

	// ResultMovie is the mapped schema from the movie endpoint
	ResultMovie struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		PosterPath  string `json:"poster_path"`
		GenreIDs    []int  `json:"genre_ids"`
		ReleaseDate string `json:"release_date"`
		Overview    string `json:"overview"`
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

	var ms Response
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return Collection{}, fmt.Errorf("failed to read movies response: %v", err)
	}

	var c Collection
	for _, m := range ms.Results {
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

	var m ResultMovie
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to read movie response: %v", err)
	}

	return m.toMovieOutput(), nil
}

func (m ResultMovie) toMovieOutput() Movie {
	r, err := time.Parse("2006-01-02", m.ReleaseDate)
	var release *time.Time
	if err == nil {
		release = &r
	}

	var genre string
	if len(m.GenreIDs) > 0 {
		genre = genres.Get(m.GenreIDs[0])
	}

	return Movie{
		ID:          m.ID,
		Name:        m.Title,
		Poster:      m.PosterPath,
		Genre:       genre,
		ReleaseDate: release,
		Overview:    m.Overview,
	}
}
