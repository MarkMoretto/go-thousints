package thousints

const (
    DefaultSliceCapacity int = 1e6
)

func NewRuneSlice() []rune {
    return make([]rune, 0, DefaultSliceCapacity)
}


// Reverse direction of string.
func ReverseString(s string, size int) string {
    rz := []rune(s)
    for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
        rz[i], rz[j] = rz[j], rz[i]
    }
    return string(rz)
}
