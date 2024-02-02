package gothousints

import (
	"testing"
)



func TestStr2Int(t *testing.T) {
    var actual int
    testCases := []struct{
        given string
        expect int
    }{
        {"0", 0},
        {"1", 1},
    }
    for _, tc := range testCases {
        actual = Str2Int[int](tc.given)
        if actual != tc.expect {
            t.Errorf("Wanted: %d, Got: %d", tc.expect, actual)
        }
    }
}
