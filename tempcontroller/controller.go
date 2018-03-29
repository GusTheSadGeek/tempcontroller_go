package tempcontroller

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
	"os/signal"
	"syscall"
)

type RelayController struct {
	Name       string
	relay	Relay
	sensor TempSensor
	updatePeriod time.Duration
}

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(10)
)

func NewRelayController(name string, relay Relay, ts TempSensor)
{
	rc := new(RelayController)
	rc.Name = name
	rc.Relay = relay
	rc.TempSensor = ts
	rc.updatePeriod = time.Second * 10
	return rc
}

func (rc RelayController) Run() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	var ticker = time.Tick(rc.updatePeriod)

	for {
		select {
		case <-gracefulStop:
			return
			break

		case <-ticker:
			if rc.sensor.triggerOnValue(){
				rc.relay.TurnOn()
			}
			if rc.sensor.triggerOffValue(){
				rc.relay.TurnOff()
			}
			break
		}
	}
}

