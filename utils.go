package blend

type Error struct {
	err string
}

func (e Error) Error() string {
	return e.err
}

func float64ToUint8(x float64) uint8 {
	if x < 0 {
		return 0
	}
	if x > 255 {
		return 255
	}
	return uint8(int(x + 0.5))
}

func float64ToUint16(x float64) uint16 {
	if x < 0.0 {
		return 0
	}
	if x > 65535.0 {
		return 65535
	}
	return uint16(int(x + 0.5))
}
