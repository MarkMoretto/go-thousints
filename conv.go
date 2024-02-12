package thousints

import "fmt"

type Integer interface {
	int | int8 | int16 | int32 | int64
}


// Convert a string to int.
func Str2Int[N Integer](s string) N {
	var out N
	for _, el := range s {
        out *= 10
        out += N(el-'0')
	}
	return out
}

// Convert an integer to string.
func Int2String[N Integer](v N) string {
	var outs string

	// Special case of only zero.
	if v == 0 {
		return "0"
	}

	for v > 0 {
		outs = fmt.Sprintf("%d", v%10) + outs
		v /= 10
	}
	return outs
}
