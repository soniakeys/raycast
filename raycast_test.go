// public domain

package raycast_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/soniakeys/raycast"
)

var square raycast.Poly

func ExamplePoly() {
	square = raycast.Poly{{1, 1}, {1, 3}, {3, 3}, {3, 1}}
}

// WP notes a ray passing through a "side" vertex is an interesting test case.

func ExampleXY_In() {
	triangle := raycast.Poly{{0, 0}, {0, 2}, {2, 1}}
	pt := raycast.XY{1, 1}
	fmt.Println(pt.In(triangle))
	// Output:
	// true
}

// tests added for coverage

func TestDegenerate(t *testing.T) {
	pg := make(raycast.Poly, 2)
	pt := raycast.XY{}
	// In is coded to return false for all points, even if a point is on the
	// polygon
	if pt.In(pg) {
		t.Fatal()
	}
}

func TestRight(t *testing.T) {
	triangle := raycast.Poly{{0, 0}, {0, 2}, {2, 1}}
	pt := raycast.XY{2.5, 1.5}
	if pt.In(triangle) {
		t.Fatal()
	}
}

// A test showing that points on vertices may return either in or out.
func TestV(t *testing.T) {
	pg := raycast.Poly{{0, 0}, {0, 1}, {1, 1}} // a horizonal and vertical
	for i := 0; i < 8; i++ {
		pg = append(pg, raycast.XY{rand.Float64(), rand.Float64()})
	}
	// log results on all vertices
	for _, pt := range pg {
		t.Log(pt.In(pg), pt)
	}
}

// A test showing pathological case near 2ULP tolerance.
func Test2ULP(t *testing.T) {
	// *             *
	// -        ^ (fp rounding)
	// -        | (bump)
	// *        +
	y3 := math.Nextafter(math.Nextafter(math.Nextafter(0, 1), 1), 1)
	pg := raycast.Poly{{0, 0}, {0, y3}, {3, y3}}
	pt := raycast.XY{1.9, 0}
	t.Log(pt.In(pg)) // gives true when it is nearly 2 ULP out.
	// (and that's with only one bump.  it seems errors over 2 ULP may be
	// possible.)
}
