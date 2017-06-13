package gpio

import (
	"github.com/kidoman/embd"
	"fmt"

	_ "github.com/kidoman/embd/host/all"
)

const (
	ON = iota
	OFF
)

func Reset(targetPins []int) {
	defer embd.CloseGPIO()

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
	if err := embd.InitGPIO(); err != nil {
        	//TODO: Return err
	        panic(err)
        }
	embd.SetDirection(targetPin, embd.Out)
	switch state {
        case ON:
                fmt.Println("Pulling pin '%s' high", targetPin)
		embd.DigitalWrite(targetPin, embd.High)
        case OFF:
                fmt.Println("Pulling pin '%s' low", targetPin)
		embd.DigitalWrite(targetPin, embd.Low)
        }

}
