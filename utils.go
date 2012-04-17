package blend

type BlendError struct {
	err string
}

func (e BlendError) Error() string {
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
	if x < 0 {
		return 0
	}
	if x > 0xFFFF {
		return 0xFFFF
	}
	return uint16(int(x + 0.5))
}
