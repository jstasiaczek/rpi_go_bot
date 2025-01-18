package motor

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	ena  = rpio.Pin(13)
	enb  = rpio.Pin(18)
	in1  = rpio.Pin(6)
	in2  = rpio.Pin(11)
	in3  = rpio.Pin(5)
	in4  = rpio.Pin(0)
	freq = 5000
)

func Setup() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// ena.Output()
	// ena.High()
	// enb.Output()
	// enb.High()
	in1.Output()
	in1.Low()
	in2.Output()
	in2.Low()
	in3.Output()
	in3.Low()
	in4.Output()
	in4.Low()

	ena.Mode(rpio.Pwm)
	enb.Mode(rpio.Pwm)
	ena.Freq(freq)
	enb.Freq(freq)
	ena.DutyCycle(100, 100)
	enb.DutyCycle(100, 100)
}

func engine1(direction int) {
	if direction > 0 {
		in1.Low()
		in2.High()
		return
	}
	if direction < 0 {
		in1.High()
		in2.Low()
	} else {
		in1.Low()
		in2.Low()
	}
}

func engine2(direction int) {
	if direction > 0 {
		in3.High()
		in4.Low()
		return
	}
	if direction < 0 {
		in3.Low()
		in4.High()
	} else {
		in3.Low()
		in4.Low()
	}
}

func Forward() {
	fmt.Println("Forward")
	engine1(1)
	engine2(1)
}

func Motor1Speed(speed int) {
	ena.DutyCycle(uint32(speed), 100)
}

func Motor2Speed(speed int) {
	enb.DutyCycle(uint32(speed), 100)
}

func Stop() {
	engine1(0)
	engine2(0)
}

func Backward() {
	engine1(-1)
	engine2(-1)
}

func Left() {
	engine1(-1)
	engine2(1)
}

func Right() {
	engine1(1)
	engine2(-1)
}
