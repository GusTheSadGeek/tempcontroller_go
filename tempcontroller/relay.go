package tempcontroller

import (
	"fmt"
	"os"
	"github.com/stianeikeland/go-rpio"
)

//Rev 2 and 3 Raspberry Pi                        Rev 1 Raspberry Pi (legacy)
//+-----+---------+----------+---------+-----+      +-----+--------+----------+--------+-----+
//| BCM |   Name  | Physical | Name    | BCM |      | BCM | Name   | Physical | Name   | BCM |
//+-----+---------+----++----+---------+-----+      +-----+--------+----++----+--------+-----+
//|     |    3.3v |  1 || 2  | 5v      |     |      |     | 3.3v   |  1 ||  2 | 5v     |     |
//|   2 |   SDA 1 |  3 || 4  | 5v      |     |      |   0 | SDA    |  3 ||  4 | 5v     |     |
//|   3 |   SCL 1 |  5 || 6  | 0v      |     |      |   1 | SCL    |  5 ||  6 | 0v     |     |
//|   4 | GPIO  7 |  7 || 8  | TxD     | 14  |      |   4 | GPIO 7 |  7 ||  8 | TxD    |  14 |
//|     |      0v |  9 || 10 | RxD     | 15  |      |     | 0v     |  9 || 10 | RxD    |  15 |
//|  17 | GPIO  0 | 11 || 12 | GPIO  1 | 18  |      |  17 | GPIO 0 | 11 || 12 | GPIO 1 |  18 |
//|  27 | GPIO  2 | 13 || 14 | 0v      |     |      |  21 | GPIO 2 | 13 || 14 | 0v     |     |
//|  22 | GPIO  3 | 15 || 16 | GPIO  4 | 23  |      |  22 | GPIO 3 | 15 || 16 | GPIO 4 |  23 |
//|     |    3.3v | 17 || 18 | GPIO  5 | 24  |      |     | 3.3v   | 17 || 18 | GPIO 5 |  24 |
//|  10 |    MOSI | 19 || 20 | 0v      |     |      |  10 | MOSI   | 19 || 20 | 0v     |     |
//|   9 |    MISO | 21 || 22 | GPIO  6 | 25  |      |   9 | MISO   | 21 || 22 | GPIO 6 |  25 |
//|  11 |    SCLK | 23 || 24 | CE0     | 8   |      |  11 | SCLK   | 23 || 24 | CE0    |   8 |
//|     |      0v | 25 || 26 | CE1     | 7   |      |     | 0v     | 25 || 26 | CE1    |   7 |
//|   0 |   SDA 0 | 27 || 28 | SCL 0   | 1   |      +-----+--------+----++----+--------+-----+
//|   5 | GPIO 21 | 29 || 30 | 0v      |     |
//|   6 | GPIO 22 | 31 || 32 | GPIO 26 | 12  |
//|  13 | GPIO 23 | 33 || 34 | 0v      |     |
//|  19 | GPIO 24 | 35 || 36 | GPIO 27 | 16  |
//|  26 | GPIO 25 | 37 || 38 | GPIO 28 | 20  |
//|     |      0v | 39 || 40 | GPIO 29 | 21  |
//

func X(phyiscalPin int) int {
	switch phyiscalPin {
	case 3:
		return 2
	case 5:
		return 3
	case 7:
		return 4
	case 11:
		return 11
	case 13:
		return 27
	case 15:
		return 22
	case 19:
		return 10
	case 23:
		return 11
	case 27:
		return 0
	case 29:
		return 5
	case 31:
		return 6
	case 33:
		return 13
	case 35:
		return 19
	case 37:
		return 26
	case 8:
		return 14
	case 10:
		return 15
	case 12:
		return 18
	case 22:
		return 25
	case 24:
		return 8
	case 26:
		return 7
	case 28:
		return 1
	case 32:
		return 12
	case 36:
		return 16
	case 38:
		return 20
	case 40:
		return 21
	}
	return 999
}

var debug = 0

type Relay struct {
	Name  string
	pin   rpio.Pin
	state int
}

func RelayInit() {
	if debug == 0 {
		if err := rpio.Open(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func (r *Relay) State() int {
	return r.state
}

func RelayClose() {
	if debug == 0 {
		rpio.Close()
	}
}

func NewRelay(name string, pin int) *Relay {
	r := new(Relay)
	r.Name = name
	q := X(pin)
	fmt.Printf("%v\n", q)
	r.pin = rpio.Pin(q)
	r.pin.Output()
	r.state = 0
	if debug == 0 {
		//r.TurnOff()
		//time.Sleep(time.Second * 5)
		//r.TurnOn()
		//time.Sleep(time.Second * 5)
		r.TurnOff()
	}
	return r
}

func (r *Relay) TurnOn() error {
	if r.state != 1 {
		fmt.Printf("%s ON\n", r.Name)
		if debug == 0 {
			r.pin.Low()
		}
		r.state = 1
	}
	return nil
}

func (r *Relay) TurnOff() error {
	if r.state != 0 {
		fmt.Printf("%s OFF\n", r.Name)
		if debug == 0 {
			r.pin.High()
		}
		r.state = 0
	}
	return nil
}
