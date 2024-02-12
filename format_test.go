package thousints

import (
	"testing"
)

func TestFormatByThousandsDefaultSep(t *testing.T) {
    testCases := []struct{
        given int
        expect string
    }{
        {1, "1"},
        {100, "100"},
        {1000, "1,000"},
        {1000000, "1,000,000"},
    }
    for _, tc := range testCases {
        actual := AddSeparatorsTo(tc.given)
        if actual != tc.expect {
            t.Errorf("Wanted: %s, Got: %s", tc.expect, actual)
        }
    }
}

func TestFormatByThousandsAltSep(t *testing.T) {
    testCases := []struct{
        given int
        sep rune
        expect string
    }{
        {1, '_', "1"},
        {100, '_', "100"},
        {1000, '_', "1_000"},
        {1000000, '_', "1_000_000"},
    }
    for _, tc := range testCases {
        actual := AddSeparatorsTo(tc.given, tc.sep)
        if actual != tc.expect {
            t.Errorf("Wanted: %s, Got: %s", tc.expect, actual)
        }
    }
}
