package controller

// byte 3, idx 2: select, analog left, analog right, start, dpad up, dpad right, dpad down, dpad left
var IncrementalValues []byte = []byte{1, 2, 4, 8, 16, 32, 64, 128}

type PS3ControllerHid struct {
	DpadUp       byte
	DpadDown     byte
	DpadLeft     byte
	DpadRight    byte
	Traingle     byte
	Circle       byte
	Cross        byte
	Square       byte
	TriggerL1    byte
	TriggerR1    byte
	TriggerL2    byte
	TriggerR2    byte
	AnalogLeftX  byte
	AnalogLeftY  byte
	AnalogRightX byte
	AnalogRightY byte
	AnalogLeft   bool
	AnalogRight  bool
	Start        bool
	Select       bool
}

func readButtonFromByte(b byte) []bool {
	result := make([]bool, 8)
	for i := range IncrementalValues {
		revI := 7 - i
		v := IncrementalValues[revI]
		result[revI] = b >= v
		if b >= v {
			b -= v
		}
	}
	return result
}

func NewPS3BytesToControllerHid(bytes []byte) PS3ControllerHid {
	result := PS3ControllerHid{
		DpadUp:       bytes[14],
		DpadDown:     bytes[16],
		DpadLeft:     bytes[17],
		DpadRight:    bytes[15],
		Traingle:     bytes[22],
		Circle:       bytes[23],
		Cross:        bytes[24],
		Square:       bytes[25],
		TriggerL1:    bytes[20],
		TriggerR1:    bytes[21],
		TriggerL2:    bytes[18],
		TriggerR2:    bytes[19],
		AnalogLeftX:  bytes[6],
		AnalogLeftY:  bytes[7],
		AnalogRightX: bytes[8],
		AnalogRightY: bytes[9],
	}
	buttons := readButtonFromByte(bytes[2])
	result.AnalogLeft = buttons[1]
	result.AnalogRight = buttons[2]
	result.Start = buttons[3]
	result.Select = buttons[0]
	return result
}
