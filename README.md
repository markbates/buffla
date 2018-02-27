# Buff.la - URL Shortner Example

This application is an example of writing a URL shortner using [Buffalo](https://gobuffalo.io). The application can be found live on the web at [https://buff.la](https://buff.la).

## Running Locally

### Requirements

* [dep](https://github.com/golang/dep)
* Postgres
* Node/Yarn
* [Buffalo](https://github.com/gobuffalo/buffalo) - v0.11.0 (or later or `development` branch)

### Installation

First, make sure you have all of the above dependencies setup and running (as appropriate).

1. Get the repo: `go get -v github.com/gobuffalo/buffla`.
1. Make sure `database.yml` is configured for your Postgres instance.
1. Run `buffalo setup`.
1. Setup OAuth keys for either GitHub, FaceBook, or Twitter (so you can log into the application). See `actions/auth.go` for what to name the keys.
1. Run `buffalo dev` - to run the application in development.
1. Have Fun!!

---

[Powered by Buffalo](http://gobuffalo.io)
