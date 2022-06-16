# Note API

A simple note API.

# Installation

Install [go](https://go.dev/dl/) and [postgres](https://www.postgresql.org/download/).

# Build and run



You must start your postgres database. 
On macOS, intel chip let set an alias for `PGDATA` in zshrc if you use zsh or other config file if you use other one.
```sh
export PGDATA="/usr/local/var/postgres/"
```

if you install postgres via homebrew set `PGDATA`
```sh
export PGDATA="/opt/homebrew/var/postgres"
```

If you have a trouble go this [stackoverflow](https://stackoverflow.com/a/33015649) questions.


Create new database
```sh
psql -h localhost
CREATE DATABASE gonote;
```

Ready to go! Just run this command and your server is live in `localhost:3000`
```sh
make run
```

# Docker
Docker setup is not complete yet.