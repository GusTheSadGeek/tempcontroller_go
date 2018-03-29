package tempcontroller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type TempSensor struct {
	name            string
	path            string
	currentValue    float64
	updatePeriod    time.Duration
	triggerOnValue  float64
	triggerOffValue float64
}

func NewTempSensor(name string, path string) *TempSensor {
	ts := new(TempSensor)
	ts.name = name
	ts.path = path
	ts.currentValue = 0
	ts.updatePeriod = time.Second * 1
	ts.triggerOnValue = 0
	ts.triggerOffValue = 0
	return ts
}

func (s *TempSensor) Current() float64 {
	return s.currentValue
}

func (s *TempSensor) SetUpdatePeriod(period time.Duration) {
	s.updatePeriod = time.Second * period
}

func (s *TempSensor) TriggerOn() bool {
	if s.triggerOnValue > s.triggerOffValue {
		return s.currentValue > s.triggerOnValue
	} else {
		return s.currentValue < s.triggerOnValue
	}
}

func (s *TempSensor) TriggerOff() bool {
	if s.triggerOnValue > s.triggerOffValue {
		return s.currentValue < s.triggerOffValue
	} else {
		return s.currentValue > s.triggerOffValue
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (s *TempSensor) readRawTempData() (string, error) {
	lines, err := readLines(s.path)
	if err != nil {
		return "", err
	}
	if len(lines) != 2 {
		return "", errors.New("unexpected number of lines return from temp sensor")
	}
	parts := strings.Split(lines[1], "t=")
	if len(parts) != 2 {
		return "", errors.New("unexpected format returned from temp sensor")
	}
	return parts[1], nil
}

func (s *TempSensor) readTemp() error {
	tempData, err := s.readRawTempData()
	if err != nil {
		return err
	}
	f, err := strconv.ParseFloat(tempData, 64)
	if err != nil {
		return err
	}
	s.currentValue = f / 1000
	return nil
}

func (s TempSensor) Run() {

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	var ticker = time.Tick(s.updatePeriod)

	for {
		select {
		case <-gracefulStop:
			return
			break

		case <-ticker:
			var err = s.readTemp()
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("%v\n", s.currentValue)
			}

			break
		}
	}
}
