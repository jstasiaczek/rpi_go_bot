package main

import (
	"fmt"
	"strings"
	"time"

	"rpi_go_bot/controller"
	"rpi_go_bot/motor"

	"github.com/sstallion/go-hid"
)

func readPS3ControllerInput(device *hid.Device, c chan controller.PS3ControllerHid, quit chan bool) {
	for {
		buf := make([]byte, 1024)
		_, err := device.Read(buf)
		if err != nil {
			continue
		}
		PS3ControllerState := controller.NewPS3BytesToControllerHid(buf)

		select {
		case c <- PS3ControllerState:
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func watchForPS3ControllerInputChanges(newState, oldState controller.PS3ControllerHid) bool {
	if controller.IsAnalogStateChanged(oldState.AnalogLeftY, newState.AnalogLeftY) {
		return true
	}
	if controller.IsAnalogStateChanged(oldState.AnalogRightY, newState.AnalogRightY) {
		return true
	}
	return false
}

func main() {
	var foundDevice *hid.DeviceInfo

	for {
		hid.Enumerate(hid.VendorIDAny, hid.ProductIDAny, func(info *hid.DeviceInfo) error {
			fmt.Printf("%s: ID %04x:%04x %s %s\n",
				info.Path, info.VendorID, info.ProductID, info.MfrStr, info.ProductStr)
			if strings.Contains(info.ProductStr, "PLAYSTATION") {
				fmt.Println("Found Playstation controller")
				foundDevice = info
			}
			return nil
		})

		if foundDevice == nil {
			fmt.Println("No Playstation controller found")
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	oldState := controller.PS3ControllerHid{}
	device, err := hid.OpenFirst(foundDevice.VendorID, foundDevice.ProductID)

	if err != nil {
		fmt.Println(err)
		return
	}

	motor.Setup()

	defer device.Close()

	c := make(chan controller.PS3ControllerHid)
	quit := make(chan bool)

	go readPS3ControllerInput(device, c, quit)

	for {
		inputs := <-c
		if watchForPS3ControllerInputChanges(inputs, oldState) {
			oldState = inputs
			handleInputsChange(inputs)
		}
	}
}

var step float32 = 100 / float32(115)

func getSpeedIntValue(value float32) int {
	if value > 100 {
		return 100
	}
	if value < 0 {
		return 0
	}
	return int(value)
}

func getSpeedFromAnalogValue(val byte) int {
	value := float32(val)
	if value < 115 {
		return getSpeedIntValue((float32(115) - value) * step)
	} else if value > 140 {
		return getSpeedIntValue((value - float32(140)) * step)
	} else {
		return 0
	}
}

func handleInputsChange(inputs controller.PS3ControllerHid) {
	if inputs.AnalogLeftY < 115 && inputs.AnalogRightY < 115 {
		fmt.Println("Forward")
		avgSpeed := (getSpeedFromAnalogValue(inputs.AnalogLeftY) + getSpeedFromAnalogValue(inputs.AnalogRightY)) / 2
		fmt.Println("avgSpeed", avgSpeed)
		motor.Motor1Speed(avgSpeed)
		motor.Motor2Speed(avgSpeed)
		motor.Forward()
	}
	if inputs.AnalogLeftY > 140 && inputs.AnalogRightY > 140 {
		fmt.Println("Backward")
		avgSpeed := (getSpeedFromAnalogValue(inputs.AnalogLeftY) + getSpeedFromAnalogValue(inputs.AnalogRightY)) / 2
		fmt.Println("avgSpeed", avgSpeed)
		motor.Motor1Speed(avgSpeed)
		motor.Motor2Speed(avgSpeed)
		motor.Backward()
	}
	if inputs.AnalogLeftY > 115 && inputs.AnalogLeftY < 140 && inputs.AnalogRightY > 115 && inputs.AnalogRightY < 140 {
		fmt.Println("Stop")
		motor.Motor1Speed(0)
		motor.Motor2Speed(0)
		motor.Stop()
	}
	if inputs.AnalogLeftY < 115 && inputs.AnalogRightY > 140 {
		fmt.Println("Left")
		motor.Motor1Speed(getSpeedFromAnalogValue(inputs.AnalogLeftY))
		motor.Motor2Speed(getSpeedFromAnalogValue(inputs.AnalogRightY))
		motor.Left()
	}
	if inputs.AnalogLeftY > 140 && inputs.AnalogRightY < 115 {
		fmt.Println("Right")
		motor.Motor1Speed(getSpeedFromAnalogValue(inputs.AnalogLeftY))
		motor.Motor2Speed(getSpeedFromAnalogValue(inputs.AnalogRightY))
		motor.Right()
	}
}
