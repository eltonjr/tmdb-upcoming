package movies

import "time"

// Movie represents an upcoming movie
type Movie struct {
	ID          int        `json:"id"`
	Name        string     `json:"name,omitempty"`
	Poster      string     `json:"poster,omitempty"`
	Genre       string     `json:"genre,omitempty"`
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
	Overview    string     `json:"overview,omitempty"`
}

// Collection will be returned by the server when asked for a list of movies.
// The idea behind this wrapper is to be able to grow the struct and maybe return
// more infos about the movies collection
type Collection struct {
	Movies []Movie `json:"movies"`
}
