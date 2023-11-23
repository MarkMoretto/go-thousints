package gothousints

import (
	"testing"
)

func TestNewRuneSlice(t *testing.T) {
    nrs := NewRuneSlice()
    actual := cap(nrs)
    if actual != DefaultSliceCapacity {
        t.Errorf("Wanted: %d, Got: %d", DefaultSliceCapacity, actual)
    }
}

func TestReverseString(t *testing.T) {
    testCases := []struct{
        given string
        size int
        expect string
    }{
        {"abcd", 4, "dcba"},
        {"", 0, ""},
        {"aaaa", 4, "aaaa"},
    }
    for _, tc := range testCases {
        actual := ReverseString(tc.given, tc.size)
        if actual != tc.expect {
            t.Errorf("Wanted: %s, Got: %s", tc.expect, actual)
        }
    }
}
