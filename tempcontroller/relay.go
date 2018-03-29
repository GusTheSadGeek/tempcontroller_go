package tempcontroller

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

type Relay struct {
	Name string `yaml:"name"`
	Pin  int    `yaml:"relaypin"`
}

func NewRelay(name string, pin int) *Relay{
	r := new(Relay)
	r.Name = name
	r.Pin = pin
	return r
}

func (s *Relay) TurnOn() error {
	fmt.Println("ON")
	return nil
}

func (s *Relay) TurnOff() error {
	fmt.Println("OFF")
	return nil
}




func Run2() {

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

