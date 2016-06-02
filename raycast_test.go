// public domain

package raycast_test

import (
	"fmt"

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
