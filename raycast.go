// public domain

// Raycast shows an implementation of the ray casting point-in-polygon
// (PNPoly) algorithm for testing if a point is inside a closed polygon.
// Also known as the crossing number or the even-odd rule algorithm.
//
// The implementation follows
// https://www.ecse.rpi.edu/Homepages/wrf/Research/Short_Notes/pnpoly.html
package raycast

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
// Segments of the polygon are allowed to cross.  In this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If pt is exactly on a segment or vertex of pg, the method may return true or
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

// Segment intersect expression from
// https://www.ecse.rpi.edu/Homepages/wrf/Research/Short_Notes/pnpoly.html
//
// Currently the compiler inlines the function by default.
func rayIntersectsSegment(p, a, b XY) bool {
	return (a.Y > p.Y) != (b.Y > p.Y) &&
		p.X < (b.X-a.X)*(p.Y-a.Y)/(b.Y-a.Y)+a.X
}
