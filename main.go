package main

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"periph.io/x/periph/conn/gpio"
)

func main() {
	adaptor := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(adaptor, "16")

	offset := 13

	work := func(){
		servo.Move(uint8(27))
	}
