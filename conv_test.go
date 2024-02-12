package thousints

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
        {"1000", 1000},
    }
    for _, tc := range testCases {
        actual = Str2Int[int](tc.given)
        if actual != tc.expect {
            t.Errorf("Wanted: %d, Got: %d", tc.expect, actual)
        }
    }
}

func TestInt2String(t *testing.T) {
    var actual string
    testCases := []struct{
        given int
        expect string
    }{
        {0, "0"},
        {1, "1"},
        {1000, "1000"},
    }
    for _, tc := range testCases {
        actual = Int2String[int](tc.given)
        if actual != tc.expect {
            t.Errorf("Wanted: %s, Got: %s", tc.expect, actual)
        }
    }
}
