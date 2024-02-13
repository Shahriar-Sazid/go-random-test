package ed

import (
	"fmt"
	"testing"
)

func TestEditDistance(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected float32
	}{
		{"sadarghat", "sodorghat", 1.0},
		{"sadarghat", "shodorghat", 1.3},
		{"shodorghat", "sadarghat", 1.5},
		{"nabinagar", "nobinagor", 1},
		{"bancharampur", "bamcharampur", 0.7},
		{"lalbag", "lalbagh", 0.3},
		{"lalbagh", "lalbag", 0.5},
		{"azimpur", "ajimpur", 0.5},
		{"gouripur", "gowrepur", 1},
		{"ziranibajar", "jiranibazar", 1},
		{"biddaloy", "biddaloi", 0.5},
		{"ganabhaban", "gonovobon", 3.0},
		{"kodomtoli", "kadamtoly", 1.5},
		{"bhuighor", "bhoghor", 1.5},
		{"bhuighor", "bhuigon", 1.5},
		{"peerganj", "pirganj", 1.5},
		{"dhaka", "dhopa", 1.5},
		{"camila", "cumilla", 1.6},
		{"charfashion", "charfashioin", 0.7},
		{"kushtia", "kushita", 2},
		{"kalaroa", "kalarmar", 2},
		{"forajikanda", "forajipara", 3},
		{"sat", "sath", 0.3},
		{"baddanogor", "boxnogor", 3.5},
		{"baddanogor", "badanogor", 1},
		{"modhubag", "modhubazar", 2.5},
		{"munsiganj", "monshgnj", 2.5},
	}

	for _, tc := range testCases {
		actual := IncrementalED(tc.s, tc.t, 0, max(len(tc.s), len(tc.t)))
		if actual != tc.expected {
			t.Errorf("EditDistance(%s, %s): expected %.1f, got %.1f", tc.s, tc.t, tc.expected, actual)
		} else {
			fmt.Printf("%s,%s: %.1f\n", tc.s, tc.t, actual)
		}
	}
}
