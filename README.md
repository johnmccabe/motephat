[![Travis Badge]][Travis]
[![Go Report Card Badge]][Go Report Card]
[![GoDoc Badge]][GoDoc]

![Mote](go-mote-phat-logo.png)
Buy the Mote PHAT & accessories here: https://shop.pimoroni.com/products/mote-phat

This library is based on both the Pimoroni Python [`mote-phat` library](https://github.com/pimoroni/mote-phat), and @alexellis's [`blinkt_go` library](https://github.com/alexellis/blinkt_go).

**WARNING THIS IS WORK IN PROGRESS, ISSUES/PULL REQUESTS STILL WELCOME THOUGH**
![WIP](http://i.imgur.com/vWBepKi.gif)

# Prerequisites

You should have Go version 1.8+ installed and your `GOPATH` configured.

You will need to have the `wiringpi` package installed.
```
sudo apt-get install -qy wiringpi
```

# Installation

Install the mote library with `go get`, like so:

```bash
go get -u github.com/johnmccabe/motephat
```
You can of course use your own choice of depedency management tool, Glide, Dep etc.


# Examples

You can run the supplied example programs (ported from their Python equivalents) as follows (installing `glide` first which is used to pull down the examples dependencies).
```
go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/johnmccabe/mote
glide install
```
Then running each example as follows.
```
go run examples/rgb/rgb.go 255 0 0
```



*The Golang Gopher was created by [Renée French](http://reneefrench.blogspot.co.uk/) and is [Creative Commons Attributions 3.0](https://creativecommons.org/licenses/by/3.0/) licensed.*

[Travis]: https://travis-ci.org/johnmccabe/motephat
[Travis Badge]: https://travis-ci.org/johnmccabe/motephat.svg?branch=master
[Go Report Card]: https://goreportcard.com/report/github.com/johnmccabe/motephat
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/johnmccabe/motephat
[GoDoc]: https://godoc.org/github.com/johnmccabe/motephat
[GoDoc Badge]: https://godoc.org/github.com/johnmccabe/motephat?status.svg
