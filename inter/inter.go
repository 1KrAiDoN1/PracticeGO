package inter

import "fmt"

type Device interface {
	On() string
	Off() string
	Status() string
}

func NewLamp() *Lamp {
	return &Lamp{}
}

type Lamp struct {
	Brightness int
}

type Condithioner struct {
	temperature int
}

type Rozetka struct {
	Power bool
}

func (l *Lamp) On() string {
	l.Brightness = 100
	return "Лампа зажглась"
}

func (l *Lamp) Off() string {
	l.Brightness = 0
	return "Лампа погасла"
}
func (l *Lamp) Status() string {
	return fmt.Sprintf("Яркость: %d", l.Brightness)
}

func (c *Condithioner) On() string {
	c.temperature = 22
	return "Кондиционер включен"
}
func (c *Condithioner) Off() string {
	c.temperature = 0
	return "Кондиционер выключен"
}
func (c *Condithioner) Status() string {
	return fmt.Sprintf("Температура: %d°C", c.temperature)
}

func (r *Rozetka) On() string {
	r.Power = true
	return "Розетка включена"
}
func (r *Rozetka) Off() string {
	r.Power = false
	return "Розетка выключена"
}
func (r *Rozetka) Status() string {
	if r.Power {
		return "Питание: on"
	} else {
		return "Питание: off"
	}
}

func ControlDevice(dev Device, command string) string {
	if command == "on" {
		return dev.On()
	}
	if command == "off" {
		return dev.Off()
	}
	if command == "status" {
		return dev.Status()
	} else {
		return "Неизвестная операция"
	}
}
