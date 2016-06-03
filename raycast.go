// public domain

// Raycast shows an implementation of the ray casting point-in-polygon
// algorithm for testing if a point is inside a closed polygon.  Also
// known as the "crossing number" or the "even-odd rule" algorithm.
package raycast

import "math"

// code originally developed for posting on Rosetta Code.

// XY is a 2D point in the Cartesian plane.
type XY struct {
	X, Y float64
}

// Poly represents a closed polygon.  Pairs of consecutive points represent
// endpoints of segments.  The last and first point represent an additional
// segment.  That is, the last point does not need to repeat the first to
// close the polygon.
type Poly []XY

// In returns true if pt is inside pg.
//
// The result is accurate until pt is within about 2 units of least precision
// (ULP) of pg.  If pt is within 2 ULP of pg, the method may return true or
// false.
func (pt XY) In(pg Poly) bool {
	if len(pg) < 3 {
		return false
	}
	a := pg[0]
	in := rayIntersectsSegment(pt, pg[len(pg)-1], a)
	for _, b := range pg[1:] {
		if rayIntersectsSegment(pt, a, b) {
			in = !in
		}
		a = b
	}
	return in
}

func rayIntersectsSegment(p, a, b XY) bool {
	if a.Y > b.Y {
		a, b = b, a
	}
	// up to 2 ULP bump here.
	for p.Y == a.Y || p.Y == b.Y {
		p.Y = math.Nextafter(p.Y, math.Inf(1))
	}
	if p.Y < a.Y || p.Y > b.Y {
		return false
	}
	if a.X > b.X {
		if p.X > a.X {
			return false
		}
		if p.X < b.X {
			return true
		}
	} else {
		if p.X > b.X {
			return false
		}
		if p.X < a.X {
			return true
		}
	}
	return (p.Y-a.Y)/(p.X-a.X) >= (b.Y-a.Y)/(b.X-a.X)
}
