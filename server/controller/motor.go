package controller

type MotorConfiguration struct {
	Name string
	Pins struct {
		Enable   int
		Forward  int
		Backward int
	}
}

type Motor interface {
	Forward()
	Backward()
	Stop()
	Enable()
	Disable()
}

type HBridgeMotor struct {
	controller  GPIOController
	enablePin   int
	forwardPin  int
	backwardPin int
}

func (m HBridgeMotor) Forward() {
	m.controller.Off(m.backwardPin)
	m.controller.On(m.forwardPin)
}

func (m *HBridgeMotor) Backward() {
	m.controller.Off(m.forwardPin)
	m.controller.On(m.backwardPin)
}

func (m *HBridgeMotor) Stop() {
	m.controller.Off(m.backwardPin)
	m.controller.Off(m.forwardPin)
}

func (m *HBridgeMotor) Enable() {
	m.controller.Off(m.enablePin)
}

func (m *HBridgeMotor) Disable() {
	m.controller.On(m.enablePin)
}

func NewHBridgeMotor(c GPIOController, enablePin int, forwardPin int, backwardPin int) HBridgeMotor {
	return HBridgeMotor{c, enablePin, forwardPin, backwardPin}
}
