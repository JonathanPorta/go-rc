package gpio

import (
	"fmt"

	"github.com/kidoman/embd"

	_ "github.com/kidoman/embd/host/all"
)

const (
	ON = iota
	OFF
)

func Reset(targetPins []int) {
	for _, targetPin := range targetPins {
		digitalWrite(targetPin, OFF)
	}
}

func WriteToPins(targetPins []int, state int) {
	for _, targetPin := range targetPins {
		digitalWrite(targetPin, state)
	}
}

func WriteToPin(targetPin int, state int) {
	digitalWrite(targetPin, state)
}

func digitalWrite(targetPin int, state int) {
	defer embd.CloseGPIO()
	if err := embd.InitGPIO(); err != nil {
		//TODO: Return err
		panic(err)
	}
	embd.SetDirection(targetPin, embd.Out)
	switch state {
	case ON:
		fmt.Printf("Pulling pin '%v' high\n", targetPin)
		embd.DigitalWrite(targetPin, embd.High)
	case OFF:
		fmt.Printf("Pulling pin '%v' low\n", targetPin)
		embd.DigitalWrite(targetPin, embd.Low)
	}

}
