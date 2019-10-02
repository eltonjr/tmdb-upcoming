package env

import (
	"log"
	"os"
)

type vars struct {
	Server struct {
		Address string
	}

	TMDB struct {
		Address string
		Key     string
	}
}

// Vars holds custom variables needed for the system to work.
// When the system starts, a process checks every needed var
// and crashes if some are missing
var Vars vars

func init() {
	Vars.Server.Address = os.Getenv("SERVER_ADDRESS")
	if Vars.Server.Address == "" {
		log.Fatal("missing SERVER_ADDRESS")
	}

	Vars.TMDB.Address = os.Getenv("TMDB_ADDRESS")
	if Vars.TMDB.Address == "" {
		log.Fatal("missing TMDB_ADDRESS")
	}

	Vars.TMDB.Key = os.Getenv("TMDB_KEY")
	if Vars.TMDB.Key == "" {
		log.Fatal("missing TMDB_KEY")
	}
}
