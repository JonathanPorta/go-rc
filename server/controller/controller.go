package controller

import (
	"fmt"
	"log"

	"github.com/jonathanporta/go-rc/server/gpio"
	//"github.com/jonathanporta/go-rc/server/gpio"
)

type Controller interface {
	On([]int)
	Off([]int)
}

type CarController struct {
	leftFrontMotor  HBridgeMotor
	rightFrontMotor HBridgeMotor
	leftBackMotor   HBridgeMotor
	rightBackMotor  HBridgeMotor
}

func (c *CarController) AddMotor(name string, enablePin int, forwardPin int, backwardPin int) {
	gc := GPIOController{}
	switch name {
	case "FrontLeft":
		fmt.Printf("NewHBridgeMotor(gc, enablePin, forwardPin, backwardPin)::NewHBridgeMotor(gc, %v, %v, %v)", enablePin, forwardPin, backwardPin)
		c.leftFrontMotor = NewHBridgeMotor(gc, enablePin, forwardPin, backwardPin)
		// c.leftFrontMotor = lfm
	case "FrontRight":
		frm := NewHBridgeMotor(gc, enablePin, forwardPin, backwardPin)
		c.rightFrontMotor = frm
	case "BackLeft":
		blm := NewHBridgeMotor(gc, enablePin, forwardPin, backwardPin)
		c.leftBackMotor = blm
	case "BackRight":
		brm := NewHBridgeMotor(gc, enablePin, forwardPin, backwardPin)
		c.rightBackMotor = brm
	default:
		fmt.Printf("Unknown motor configuration '%s' provided. Should be one of 'FrontLeft', 'FrontRight', 'BackLeft', or 'BackRight':\n%v\n%v\n%v\n", name, enablePin, forwardPin, backwardPin)
		log.Fatalf("Please correct your configuration... Exiting...")
	}
}

func (c *CarController) Left() {
	fmt.Println("MoveLeft")

	c.leftFrontMotor.Backward()
	c.leftBackMotor.Backward()
	c.rightFrontMotor.Forward()
	c.rightBackMotor.Forward()
}
func (c *CarController) Right() {
	fmt.Println("MoveRight")

	c.leftFrontMotor.Forward()
	c.leftBackMotor.Forward()
	c.rightFrontMotor.Backward()
	c.rightBackMotor.Backward()
}
func (c CarController) Forward() {
	fmt.Println("MoveForward")

	c.leftFrontMotor.Forward()
	c.rightFrontMotor.Forward()
	c.leftBackMotor.Forward()
	c.rightBackMotor.Forward()
}
func (c CarController) Backward() {
	fmt.Println("MoveBackward")

	c.leftFrontMotor.Backward()
	c.rightFrontMotor.Backward()
	c.leftBackMotor.Backward()
	c.rightBackMotor.Backward()
}
func (c CarController) Stop() {
	fmt.Println("STOP")

	c.leftFrontMotor.Stop()
	c.rightFrontMotor.Stop()
	c.leftBackMotor.Stop()
	c.rightBackMotor.Stop()
}

func (c *CarController) On(pins []int) {

}
func (c *CarController) Off(pins []int) {

}

func NewCarController(motors []MotorConfiguration) CarController {
	c := CarController{}
	for _, m := range motors {
		c.AddMotor(m.Name, m.Pins.Enable, m.Pins.Forward, m.Pins.Backward)
	}
	fmt.Printf("\n\n\n\nNewCarController\n\n%v\n\n", c)
	return c
}

type GPIOController struct{}

func (gc *GPIOController) On(pin int) {
	fmt.Printf("Writing 'On' to '%v'.\n", pin)
	gpio.WriteToPin(pin, gpio.ON)
}

func (gc *GPIOController) Off(pin int) {
	fmt.Printf("Writing 'Off' to '%v'.\n", pin)
	gpio.WriteToPin(pin, gpio.OFF)
}
