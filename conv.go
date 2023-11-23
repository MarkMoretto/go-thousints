package gothousints


// Convert an integer to a rune.
func Int2Rune[N int | int32 | int64](n N) rune {
	return rune('0' + n)
}

// Convert an integer to string.
func Int2String[N int | int32 | int64](v N) string {
	rz := NewRuneSlice()
	for v > 0 {
		// rz = append([]rune{runeify(v % 10)}, rz...)
		rz = append(rz, Int2Rune(v % 10))
		v /= 10
	}
	return string(rz)
}
