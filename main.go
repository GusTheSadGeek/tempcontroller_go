package main

/*
A blinker example using go-rpio library.
Requires administrator rights to run
Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.
*/

import (
	"github.com/GusTheSadGeek/tempcontroller_go/tempcontroller"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ts := tempcontroller.NewTempSensor("Test", "/tmp/wibble")
	ts.SetUpdatePeriod(1)
	go ts.Run()


	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<- gracefulStop
}
