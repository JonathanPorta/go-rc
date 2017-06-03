package gpio

import (
	"github.com/kidoman/embd"

	_ "github.com/kidoman/embd/host/all"
)

const (
	ON = iota
	OFF
)

func Reset() {
	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()
	targetPins := []int{15, 18, 17}
	for _, targetPin := range targetPins {

		embd.SetDirection(targetPin, embd.Out)
		embd.DigitalWrite(targetPin, embd.Low)
	}
}

func WriteToPin(targetPin int, state int) {
	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	embd.SetDirection(targetPin, embd.Out)
	switch state {
	case ON:
		embd.DigitalWrite(targetPin, embd.High)
	case OFF:
		embd.DigitalWrite(targetPin, embd.Low)
	}
}
