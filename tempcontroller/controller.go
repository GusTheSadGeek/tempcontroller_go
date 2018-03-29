package tempcontroller

import (
	"os"
	"time"

	"os/signal"
	"syscall"

	"github.com/stianeikeland/go-rpio"
)

type RelayController struct {
	Name         string
	relay        Relay
	sensor       TempSensor
	updatePeriod time.Duration
}

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(10)
)

func NewRelayController(name string, relay Relay, ts TempSensor) *RelayController {
	rc := new(RelayController)
	rc.Name = name
	rc.relay = relay
	rc.sensor = ts
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
			if rc.sensor.TriggerOn() {
				rc.relay.TurnOn()
			}
			if rc.sensor.TriggerOff() {
				rc.relay.TurnOff()
			}
			break
		}
	}
}
