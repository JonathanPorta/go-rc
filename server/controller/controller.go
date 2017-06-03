package controller

import (
	"fmt"
	//"github.com/jonathanporta/go-rc/server/gpio"
)

const (
	LEFT_FRONT_ENABLE_PIN   = 17
	LEFT_FRONT_FORWARD_PIN  = 17
	LEFT_FRONT_BACKWARD_PIN = 17

	RIGHT_FRONT_ENABLE_PIN   = 17
	RIGHT_FRONT_FORWARD_PIN  = 17
	RIGHT_FRONT_BACKWARD_PIN = 17

	LEFT_REAR_ENABLE_PIN   = 17
	LEFT_REAR_FORWARD_PIN  = 17
	LEFT_REAR_BACKWARD_PIN = 17

	RIGHT_REAR_ENABLE_PIN   = 17
	RIGHT_REAR_FORWARD_PIN  = 17
	RIGHT_REAR_BACKWARD_PIN = 17
)

func on(pin int) {
	//gpio.WriteToPin(pin, gpio.ON)
}
func off(pin int) {
	//gpio.WriteToPin(pin, gpio.OFF)
}

func MoveLeft() {
	//RIGHT_FRONT_ENABLE_PIN
	//RIGHT_REAR_ENABLE_PIN
	fmt.Println("MoveLeft")
}
func MoveRight() {
	//RIGHT_FRONT_ENABLE_PIN
	//RIGHT_REAR_ENABLE_PIN
	fmt.Println("MoveRight")
}
func MoveForward() {
	//RIGHT_FRONT_ENABLE_PIN
	//RIGHT_REAR_ENABLE_PIN
	fmt.Println("MoveForward")
}
func MoveBackward() {
	//RIGHT_FRONT_ENABLE_PIN
	//RIGHT_REAR_ENABLE_PIN
	fmt.Println("MoveBackward")
}
