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

var pins = make(map[int]*embd.DigitalPin)

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
	pin := pins[targetPin]
	if pin == nil {
		fmt.Printf("Pin '%v' not initialized yet\n", targetPin)
		pin, err := embd.NewDigitalPin(targetPin)
		if err != nil {
			fmt.Printf("Unable to init pin: '%v' ", targetPin)
			panic(err)
		}
		pin.SetDirection(embd.Out)
		pins[targetPin] = &pin
	}

	if err := embd.InitGPIO(); err != nil {
		//TODO: Return err
		panic(err)
	}
	//embd.SetDirection(targetPin, embd.Out)
	switch state {
	case ON:
		fmt.Printf("Pulling pin '%v' high\n", targetPin)
		fmt.Printf("Pin: '%v'\n", pin)
		pin.Write(embd.High)
	case OFF:
		fmt.Printf("Pulling pin '%v' low\n", targetPin)
		fmt.Printf("Pin: '%v'\n", pin)
		pin.Write(embd.Low)
	}

}
