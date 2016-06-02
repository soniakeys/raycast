// public domain

package raycast_test

import (
	"testing"

	"github.com/soniakeys/raycast"
)

// test cases taken from test suite of https://github.com/JamesMilnerUK/pip-go.

var rectangle = raycast.Poly{{1, 1}, {1, 2}, {2, 2}, {2, 1}}

var tcs = []struct {
	pt raycast.XY
	in bool
}{
	{raycast.XY{1.1, 1.1}, true},
	{raycast.XY{1.2, 1.2}, true},
	{raycast.XY{1.3, 1.3}, true},
	{raycast.XY{1.4, 1.4}, true},
	{raycast.XY{1.5, 1.5}, true},
	{raycast.XY{1.6, 1.6}, true},
	{raycast.XY{1.7, 1.7}, true},
	{raycast.XY{1.8, 1.8}, true},

	{raycast.XY{-4.9, 1.2}, false},
	{raycast.XY{10.0, 10.0}, false},
	{raycast.XY{-5.0, -6.0}, false},
	{raycast.XY{-13.0, 1.0}, false},
	{raycast.XY{4.9, -1.2}, false},
	{raycast.XY{10.0, -10.0}, false},
	{raycast.XY{5.0, 6.0}, false},
	{raycast.XY{-13.0, 1.0}, false},
}

func TestJM(t *testing.T) {
	for _, tc := range tcs {
		if got := tc.pt.In(rectangle); got != tc.in {
			t.Fatal(tc.pt, "got", got, "want", tc.in)
		}
	}
}
