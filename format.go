package thousints

const (
    DefaultSeparator = ','
)

type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Main method to add specified separators to an integer value.
func AddSeparatorsTo[N Integer](v N, seperator ...rune) string {
    var (
        numStr, rNumStr string
        charCount int
        sep rune
    )

    rz := NewRuneSlice()

    // Convert integer to string.
    numStr = Int2String(v)
    charCount = len(numStr)

    // Reverse the string.
    rNumStr = ReverseString(numStr, charCount)

    // Default separator is comma.
    sep = DefaultSeparator
    if len(seperator) == 1 {
        sep = seperator[0]
    }

    // Add specified separator every third character.
    for idx, ch := range rNumStr {
        if idx > 0 && idx%3 == 0 {
            rz = append(rz, sep)
        }
        rz = append(rz, ch)
    }

    // Reverse sting and return result.
    return ReverseString(string(rz), len(rz))
}
