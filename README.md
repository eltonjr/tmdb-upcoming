## Upcoming movies

#### About

This is a webapp that displays upcoming movies. It's built on top of The Movies DataBase.

The stack used is mainly Go for the backend and Vuejs for the frontend.

A more extensive frontend documentation can be found [here](web/);
Backend documentation can be found [here](server/);

#### Building and running

You need to provide the secret key for TBDM api.

		cd server && make secrets-unlock # a password is prompted
		# OR
		echo "TMDB_KEY=<your_key>" > server/secrets.env

		# then
		docker-compose up --build

#### TODO

- i18n
- Adding golang-ci to lint the backend
