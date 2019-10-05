## Upcoming movies API

#### About

This API serves movie data for a simple frontend.  
It is made in Go and queries The Movie DataBase to get the data.

An OpenAPI spec can be found in [docs/openapi.json](docs/openapi.json).

#### Building and running - with Go

A `Makefile` is provided to perform most tasks.

TMDB requires a Key to access the data. You must `export TMDB_KEY=<your_key>` or unlock the provided key using `make secrets-unlock` with the right key ;) After that, you should be able to run the program.

Make targets:

    Usage:
        make <target>

    Targets:
        help                  Display this help
        deps                  Install go dependencies based on go mod
        build                 Compiles the system's binary
        run                   Run the system locally
        test                  Run unit tests
        secrets-unlock        Unlock the secret variables
        secrets-lock          Lock the secret variables

#### Building and running - with Docker

A `Makefile` is provided to perform most tasks.

You will not need to have Go in your local environment, just Docker is enough. However, the secrets variables will need to be exported anyway.

    Usage:
        make <target>

    Targets:
        help                  Display this help
        secrets-unlock        Unlock the secret variables
        ops-build             Compiles the system using docker
        ops-run               Runs the system using docker


#### Using the API

    # Getting the upcoming movies
    curl http://localhost:9000/v1/movies

    # Getting movies filtered by name
    curl http://localhost:9000/v1/movies?name=snatch

    # Checking if the API is online
    curl http://localhost:9000/v1/health
    

#### Architectural decisions

The project is organized in a Modular Architecture following Go's standard library style.  
Each directory holds every business logic about its subject.  

The only external dependency is `httprouter` to handle http requests. While not completely necessary, it helps when dealing with path params. httprouter is also clean and fast, not adding too much magic in the handlers.
