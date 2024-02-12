package thousints

const (
    DefaultSeparator = ','
)

func FormatByThousands(v int, seperator ...rune) string {
    var sep rune

    // String is reversed.
    numstr := Int2String(v)
    rz := NewRuneSlice()

    // Default separator is comma.
    sep = DefaultSeparator
    if len(seperator) == 1 {
        sep = seperator[0]
    }

    for idx, ch := range numstr {
        if idx > 0 && idx%3 == 0 {
            rz = append(rz, sep)
        }
        rz = append(rz, ch)
    }
    return ReverseString(string(rz), len(rz))
}
