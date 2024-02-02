package gothousints


// Convert an digit (integer in therange 0 - 9) to a rune.
// Note: There is no check to ensure that a value is a single digit.
func Int2Rune[N int | int32 | int64](n N) rune {
    return '0' + rune(n)
}

// Convert a string to int.
func Str2Int[N int | int32 | int64](s string) N {
	var out N
	for _, el := range s {
		out = out * 10 + N(el - '0')
	}
	return out
}

// Convert an integer to string.
func Int2String[N int | int32 | int64](v N) string {
    // Special case of only zero.
    if v == 0 {
        return "0"
    }
	rz := NewRuneSlice()
	for v > 0 {
		rz = append(rz, rune(v % 10))
		v /= 10
	}
	return string(rz)
}
