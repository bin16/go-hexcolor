package hexcolor

// char rune to byte
func rToB(r uint8) uint8 {
	if r >= 'a' && r <= 'f' {
		return r - 'a' + 10
	}

	if r >= 'A' && r <= 'F' {
		return r - 'A' + 10
	}

	if r >= '0' && r <= '9' {
		return r - '0'
	}

	return 0x0
}

// byte to char rune
func bToR(b uint8) uint8 {
	if b <= 9 {
		return b + '0'
	}

	return b - 10 + 'A'
}
