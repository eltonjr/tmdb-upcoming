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
		ID          int            `json:"id"`
		Title       string         `json:"title"`
		PosterPath  string         `json:"poster_path"`
		GenreIDs    []int          `json:"genre_ids"`
		Genres      []genres.Genre `json:"genres"`
		ReleaseDate string         `json:"release_date"`
		Overview    string         `json:"overview"`
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
		c.Movies = append(c.Movies, m.toSlimMovieOutput())
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

	return m.toFullMovieOutput(), nil
}

func (m ResultMovie) toSlimMovieOutput() Movie {
	return m.toMovieOutput(true)
}

func (m ResultMovie) toFullMovieOutput() Movie {
	return m.toMovieOutput(false)
}

func (m ResultMovie) toMovieOutput(slim bool) Movie {
	movie := Movie{
		ID:     m.ID,
		Name:   m.Title,
		Poster: m.PosterPath,
	}

	if len(m.GenreIDs) > 0 {
		movie.Genre = genres.Get(m.GenreIDs[0])
	}

	if len(m.Genres) > 0 {
		movie.Genre = m.Genres[0].Name
	}

	if !slim {
		r, err := time.Parse("2006-01-02", m.ReleaseDate)
		if err == nil {
			movie.ReleaseDate = &r
		}

		movie.Overview = m.Overview
	}

	return movie
}
