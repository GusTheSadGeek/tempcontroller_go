package tempcontroller

import (
	"os"
	"time"

	"fmt"
	"os/signal"
	"syscall"
)

type RelayController struct {
	Name         string
	relay        *Relay
	sensor       *TempSensor
	updatePeriod time.Duration
}

func NewRelayController(name string, relay *Relay, ts *TempSensor) *RelayController {
	rc := new(RelayController)
	rc.Name = name
	rc.relay = relay
	rc.sensor = ts
	rc.updatePeriod = time.Second * 60
	return rc
}

func log(temp float64, relay int) {
	tm := time.Now().UTC().Format("2006-01-02T15:04:05.000000000Z07:00")
	line := fmt.Sprintf("time:%s\ttemp:%v\trelay:%v\n", tm, temp, relay)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("/var/log/temps/temp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("%v",err)
	}
	defer f.Close()

	if _, err := f.Write([]byte(line)); err != nil {
		fmt.Printf("%v",err)
	}
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
				log(rc.sensor.Current(), rc.relay.State())
			}

			if rc.sensor.TriggerOff() {
				rc.relay.TurnOff()
				log(rc.sensor.Current(), rc.relay.State())
			}
			break
		}
	}
}
