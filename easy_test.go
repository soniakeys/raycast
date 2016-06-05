// public domain

package raycast_test

import (
	"github.com/soniakeys/raycast"
	"testing"
)

// easy_test.go contains some test cases from other pnp implementations
// found on the internet.  The cases selected here are "easy" in that none
// fall exactly on segments or vertices.  All are clearly in or out.
type pgCase struct {
	name string
	pg   raycast.Poly
	pts  []ptCase
}

type ptCase struct {
	pt   raycast.XY
	want bool
}

func TestEasy(t *testing.T) {
	for _, pg := range easy {
		for _, pt := range pg.pts {
			if got := pt.pt.In(pg.pg); got != pt.want {
				t.Fatal(pg.name, pt.pt, "got", got, "want", pt.want)
			}
		}
	}
}

var easy = []pgCase{
	// test cases from http://rosettacode.org/wiki/Ray-casting_algorithm#Go
	{"rc square",
		raycast.Poly{{0, 0}, {10, 0}, {10, 10}, {0, 10}},
		[]ptCase{
			{raycast.XY{5, 5}, true},
			{raycast.XY{5, 8}, true},
			{raycast.XY{-10, 5}, false},
			{raycast.XY{8, 5}, true},
			{raycast.XY{1, 2}, true},
			{raycast.XY{2, 1}, true},
		}},
	{"rc square hole", // (there's a 0-width isthmus)
		raycast.Poly{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0},
			{2.5, 2.5}, {7.5, 2.5}, {7.5, 7.5}, {2.5, 7.5}, {2.5, 2.5}},
		[]ptCase{
			{raycast.XY{5, 5}, false},
			{raycast.XY{5, 8}, true},
			{raycast.XY{-10, 5}, false},
			{raycast.XY{8, 5}, true},
			{raycast.XY{1, 2}, true},
			{raycast.XY{2, 1}, true},
		}},
	{"rc strange", // (there's a 0-width spit)
		raycast.Poly{{0, 0}, {2.5, 2.5}, {0, 10}, {2.5, 7.5},
			{7.5, 7.5}, {10, 10}, {10, 0}, {2.5, 2.5}},
		[]ptCase{
			{raycast.XY{5, 5}, true},
			{raycast.XY{5, 8}, false},
			{raycast.XY{-10, 5}, false},
			{raycast.XY{8, 5}, true},
			{raycast.XY{1, 2}, false},
			{raycast.XY{2, 1}, false},
		}},
	{"rc exagon", // (sic) "hexagon"
		raycast.Poly{{3, 0}, {7, 0}, {10, 5}, {7, 10}, {3, 10}, {0, 5}},
		[]ptCase{
			{raycast.XY{5, 5}, true},
			{raycast.XY{5, 8}, true},
			{raycast.XY{-10, 5}, false},
			{raycast.XY{8, 5}, true},
			{raycast.XY{10, 10}, false},
			{raycast.XY{1, 2}, false},
			{raycast.XY{2, 1}, false},
		}},

	// https://github.com/JamesMilnerUK/pip-go.
	{"jm rectangle",
		raycast.Poly{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
		[]ptCase{
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
		}},

	// https://github.com/substack/point-in-polygon
	{"ss box",
		raycast.Poly{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
		[]ptCase{
			{raycast.XY{1.5, 1.5}, true},
			{raycast.XY{1.2, 1.9}, true},
			{raycast.XY{0, 1.9}, false},
			{raycast.XY{1.5, 2}, false},
			{raycast.XY{1.5, 2.2}, false},
			{raycast.XY{3, 5}, false},
		}},

	// https://github.com/sromku/polygon-contains-point
	{"sr simple",
		raycast.Poly{{1, 3}, {2, 8}, {5, 4}, {5, 9}, {7, 5}, {6, 1}, {3, 1}},
		[]ptCase{
			{raycast.XY{5.5, 7}, true},
			{raycast.XY{4.5, 7}, false},
		}},
	{"sr holes",
		raycast.Poly{
			{1, 2}, {1, 6}, {8, 7}, {8, 1},
			{2, 3}, {5, 5}, {6, 2}, {8, 1},
			{6, 6}, {7, 6}, {7, 5}, {8, 1}},
		[]ptCase{
			{raycast.XY{6, 5}, true},
			{raycast.XY{4, 3}, false},
			{raycast.XY{6.5, 5.8}, false},
		}},
}
