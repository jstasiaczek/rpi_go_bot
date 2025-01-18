package controller

func ValueInRange(current, incomming, step byte) bool {
	min := current - step
	max := current + step
	return incomming >= min && incomming <= max
}

func ValueInRangeDefault(current, incomming byte) bool {
	return ValueInRange(current, incomming, 10)
}

func IsButtonPressed(value byte) bool {
	return value > 0
}

func ButtonStateChanged(current, incomming byte) bool {
	return IsButtonPressed(current) != IsButtonPressed(incomming)
}

func IsAnalogStateChanged(current, incomming byte) bool {
	return !ValueInRangeDefault(current, incomming)
}
