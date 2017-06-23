/*
Package motephat represents a connected Pimoroni Mote Phat device, communicating over GPIO.

It allows you to configure the 4 channels and set individual pixels, see the `examples` subdirectory for soem demo applications using the library.

It is a port of the Pimoroni Mote Python library (https://github.com/pimoroni/mote-phat), and is based on @alexellis' Blinkt Go library (https://github.com/alexellis/blinkt_go)

The Mote device can be obtained directly from Pimoroni (https://shop.pimoroni.com/products/motephat).
*/
package motephat

import (
	"fmt"

	"github.com/alexellis/rpi"
)

const DAT int = 10
const CLK int = 11

var CHANNEL_PINS = [...]int{8, 7, 25, 24}

var _gpioSetup = false

// NumChannels is the number of available channel connections on the Mote device
const NumChannels = 4

// NumPixelsPerChannel is the maximum addressable number of pixels across a single channel
const NumPixelsPerChannel = 16

const BRIGHTNESS int = 15

// const MAX_BRIGHTNESS = 0b01111
const MAX_BRIGHTNESS int = 15

// Mote represents a connected Pimoroni Mote device
type Mote struct {
	Channels [NumChannels]*Channel
}

// Pixel represents a single RGB pixel
type Pixel struct {
	Red, Green, Blue int
}

// Channel represents a single channel on the Mote board
type Channel struct {
	Pixels          []Pixel
	GammaCorrection bool
}

// NewMote creates a connection to a Mote Phat device, communicating via GPIO.
//
func NewMote() *Mote {
	mote := Mote{}
	return &mote
}

// ConfigureChannel configures a channel, taking the following parameters.
//
//   - channel: Channel, either 1, 2, 3 or 4 corresponding to numbers on Mote
//   - numPixels: Number of pixels to configure for this channel
//   - gammaCorrection: Whether to enable gamma correction
func (m *Mote) ConfigureChannel(channel, numPixels int, gammaCorrection bool) error {
	fmt.Println("entering ConfigureChannel")
	defer fmt.Println("exiting ConfigureChannel")
	if channel > NumChannels || channel < 1 {
		return fmt.Errorf("channel index must be between 1 and 4")
	}
	if numPixels > NumPixelsPerChannel {
		return fmt.Errorf("number of pixels can not be more than %d", NumPixelsPerChannel)
	}

	p := []Pixel{}
	for i := 0; i < numPixels; i++ {
		p = append(p, Pixel{0, 0, 0})
	}
	c := Channel{
		Pixels:          p,
		GammaCorrection: gammaCorrection,
	}
	m.Channels[channel-1] = &c

	return nil
}

// SetPixel Set the RGB colour of a single pixel, on a single channel, taking the following parameters.
//
//   - channel: Channel, either 1, 2, 3 or 4 corresponding to numbers on Mote
//   - index: Index of pixel to set, from 0 up
//   - r: Amount of red: 0-255
//   - g: Amount of green: 0-255
//   - b: Amount of blue: 0-255
func (m *Mote) SetPixel(channel, index, r, g, b int) error {
	fmt.Println("entering SetPixel")
	defer fmt.Println("exiting SetPixel")
	if channel > NumChannels || channel < 1 {
		return fmt.Errorf("channel index must be between 1 and 4")
	}
	if m.Channels[channel-1] == nil {
		return fmt.Errorf("please set up channel %d before using it", channel)
	}
	if index >= len(m.Channels[channel-1].Pixels) {
		return fmt.Errorf("Pixel index must be < %d", m.Channels[channel-1].Pixels)
	}
	m.Channels[channel-1].Pixels[index] = Pixel{r & 0xff, g & 0xff, b & 0xff}
	return nil
}

func selectChannel(c int) {
	for x := 0; x < NumChannels; x++ {
		mode := rpi.HIGH
		if x == c {
			mode = rpi.LOW
		}
		fmt.Printf("setting channel %d to %d\n", x, mode)
		rpi.DigitalWrite(rpi.GpioToPin(CHANNEL_PINS[x]), mode)
	}
}

// Setup initializes GPIO via WiringPi base library.
func (m *Mote) Setup() {
	rpi.WiringPiSetup()
	rpi.PinMode(rpi.GpioToPin(DAT), rpi.OUTPUT)
	rpi.PinMode(rpi.GpioToPin(CLK), rpi.OUTPUT)
	for x := 0; x < NumChannels; x++ {
		rpi.PinMode(rpi.GpioToPin(CHANNEL_PINS[x]), rpi.OUTPUT)
	}
}

// Show sends the pixel buffer to the hardware.
func (m *Mote) Show() {
	fmt.Println("entering Show")
	defer fmt.Println("exiting Show")
	for channel, data := range m.Channels {
		if data == nil {
			fmt.Printf("skipping empty channel %d\n", channel+1)
			continue
		}
		selectChannel(channel)
		sof()
		for _, pixel := range data.Pixels {
			// 0b11100000 (224)
			bitwise := 224
			writeByte(bitwise | BRIGHTNESS)
			writeByte(pixel.Blue)
			writeByte(pixel.Green)
			writeByte(pixel.Red)
			fmt.Printf("written bytes, r: %d, g: %d, b: %d\n", pixel.Red, pixel.Green, pixel.Blue)
		}
		eof()
	}
}

// pulse sends a pulse through the DAT/CLK pins
func pulse(pulses int) {
	rpi.DigitalWrite(rpi.GpioToPin(DAT), 0)
	for i := 0; i < pulses; i++ {
		rpi.DigitalWrite(rpi.GpioToPin(CLK), 1)
		rpi.DigitalWrite(rpi.GpioToPin(CLK), 0)
	}
}

// eof end of file or signal, from Python library
func eof() {
	pulse(36)
}

// sof start of file (name from Python library)
func sof() {
	pulse(32)
}

func writeByte(val int) {
	for i := 0; i < 8; i++ {
		// 0b10000000 = 128
		rpi.DigitalWrite(rpi.GpioToPin(DAT), val&128)
		rpi.DigitalWrite(rpi.GpioToPin(CLK), 1)
		val = val << 1
		rpi.DigitalWrite(rpi.GpioToPin(CLK), 0)
	}
}
