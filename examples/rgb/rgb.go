/*
A simple static RGB colour setting example.

This is a port of https://github.com/pimoroni/mote/blob/master/python/examples/rgb.py

You can set the colour of your Mote sticks by passing the RGB values as arguments.

  go run rgb.go 255 255 255
*/
package main

import (
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/johnmccabe/motephat"
)

func main() {
	mote := motephat.NewMote()
	mote.Setup()

	mote.ConfigureChannel(1, 16, false)
	mote.ConfigureChannel(2, 16, false)
	mote.ConfigureChannel(3, 16, false)
	mote.ConfigureChannel(4, 16, false)

	r, _ := strconv.Atoi(os.Args[1])
	g, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])

	fmt.Printf("r: %d, g: %d, b: %d\n", r, g, b)

	for channel := 1; channel < 5; channel++ {
		for pixel := 0; pixel < 16; pixel++ {
			mote.SetPixel(channel, pixel, r, g, b)
		}
		time.Sleep(10 * time.Millisecond)
	}

	mote.Show()
}
