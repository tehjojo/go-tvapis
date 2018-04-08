# go-tvmaze
[![GoDoc](https://godoc.org/github.com/mrobinsn/go-tvmaze/tvmaze?status.svg)](https://godoc.org/github.com/mrobinsn/go-tvmaze/tvmaze)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrobinsn/go-tvmaze)](https://goreportcard.com/report/github.com/mrobinsn/go-tvmaze)
[![Build Status](https://travis-ci.org/mrobinsn/go-tvmaze.svg?branch=master)](https://travis-ci.org/mrobinsn/go-tvmaze)
[![Coverage Status](https://coveralls.io/repos/github/mrobinsn/go-tvmaze/badge.svg?branch=master)](https://coveralls.io/github/mrobinsn/go-tvmaze?branch=master)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)


> TVMaze API bindings for Go (golang)

## Documentation
[GoDoc](https://godoc.org/github.com/mrobinsn/go-tvmaze/tvmaze)

## Features
- Search shows by name, TVMaze ID, TVDB ID, or TVRage ID
- Get episodes for a show
- Get the next episode for a show

## Installation
To install the package, run `go get github.com/mrobinsn/go-tvmaze/tvmaze`

To use it in application, import `"github.com/mrobinsn/go-tvmaze/tvmaze"`

## Library Usage
```
show, _ := tvmaze.DefaultClient.GetShowWithID("315") // Archer
episode, _ := c.GetNextEpisode(show)
```

## Contributing
Pull requests welcome.
