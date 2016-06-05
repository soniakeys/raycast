// public domain

package raycast_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/soniakeys/raycast"
)

var square raycast.Poly

func ExamplePoly() {
	square = raycast.Poly{{1, 1}, {1, 3}, {3, 3}, {3, 1}}
}

// WP notes a ray passing through a "side" vertex is an interesting test case.
// The test case selected for ExampleXY shows the function working properly
// in this case.

func ExampleXY_In() {
	triangle := raycast.Poly{{0, 0}, {0, 2}, {2, 1}}
	pt := raycast.XY{1, 1}
	fmt.Println(pt.In(triangle))
	// Output:
	// true
}

func ExampleXY_In_hourglass() {
	// hg is an hourglass-shape, constructed with crossing segments.
	// Both regions are on the perimiter so In returns true for points
	// in both regions.
	hg := raycast.Poly{{0, 0}, {2, 4}, {0, 4}, {2, 0}}
	top := raycast.XY{1, 3}
	bottom := raycast.XY{1, 1}
	fmt.Println(top.In(hg))
	fmt.Println(bottom.In(hg))
	// Output:
	// true
	// true
}

func ExampleXY_In_star() {
	// star is a five-pointed star constructed of five crossing segments.
	// Regions of the points of the star are on the perimeter; In returns
	// true for these regions.  The center region is not on the perimeter
	// but borders the perimeter regions.  By the two-coloring then, In
	// returns false for the center region.
	star := raycast.Poly{{0, 3}, {4, 3}, {1, 0}, {2, 5}, {3, 0}}
	top := raycast.XY{2, 4}
	center := raycast.XY{2, 2}
	fmt.Println(top.In(star))
	fmt.Println(center.In(star))
	// Output:
	// true
	// false
}

// added for coverage
func TestDegenerate(t *testing.T) {
	pg := make(raycast.Poly, 2)
	pt := raycast.XY{}
	// In is coded to return false for all points, even if a point is on the
	// polygon
	if pt.In(pg) {
		t.Fatal()
	}
}

// show that points on vertices may return either in or out
func TestV(t *testing.T) {
	pg := make(raycast.Poly, 8)
	for i := range pg {
		pg[i] = raycast.XY{rand.Float64(), rand.Float64()}
	}
	// log results on all vertices
	for _, pt := range pg {
		t.Log(pt.In(pg), pt)
	}
}
