# go-tvmaze
[![Build Status](https://travis-ci.org/mrobinsn/go-tvmaze.svg?branch=master)](https://travis-ci.org/mrobinsn/go-tvmaze)
> TVMaze API bindings for Go (golang)

## Documentation
https://godoc.org/github.com/mrobinsn/go-tvmaze/tvmaze

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
