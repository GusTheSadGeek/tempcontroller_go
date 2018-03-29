package main

/*
A blinker example using go-rpio library.
Requires administrator rights to run
Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.
*/

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/GusTheSadGeek/tempcontroller_go/tempcontroller"
)

func main() {
	tempcontroller.RelayInit()
	defer tempcontroller.RelayClose()

	ts := tempcontroller.NewTempSensor("Temp1", "/sys/bus/w1/devices/28-041501b016ff/w1_slave")
	ts.SetUpdatePeriod(60)
	ts.SetTriggerValues(23.5,24.5)

	relay := tempcontroller.NewRelay("Relay1", 33)

	go ts.Run()
	controller := tempcontroller.NewRelayController("Name", relay, ts)
	go controller.Run()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
}
