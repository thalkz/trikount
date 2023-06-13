# Trikount

Trikount is a webapp to track expenses in groups, for travels, projects, with flatmates etc. All pages are rendered server-side in HTML/CSS and a tiny bit of Javascript. Trikount should work fine on desktop and mobile, even very old ones.

## Install and start locally

To start Trikount locally, follow these steps
* Clone the repository `git clone https://github.com/thalkz/trikount`
* Cd into the folder `cd trikount/`
* Install dependencies `go install`, you might need to install a C compiler (for compiling sqlite)
* Start with `go run .` (or use [air](https://github.com/cosmtrek/air) for auto-reload)
* Open `http://localhost:8080` on your browser

## Run with Docker

- Start Docker
- Run `docker compose up`

## Deploy

- Create a semver tag `git tag v1.0.5`
- Push & push tags `git push && git push --tags`

This triggers the `deploy.yml` Github Action, which builds the image, deploys it to Github Container Repository and triggers the `./deploy.sh` script on the VPS.