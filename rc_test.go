// public domain

package raycast_test

import (
	"testing"

	"github.com/soniakeys/raycast"
)

// test cases used in Go example on Rosetta Code

type np struct {
	name string
	pg   raycast.Poly
}

var tpg = []np{
	{"square", raycast.Poly{{0, 0}, {10, 0}, {10, 10}, {0, 10}}},
	{"square hole", raycast.Poly{
		{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0},
		{2.5, 2.5}, {7.5, 2.5}, {7.5, 7.5}, {2.5, 7.5}, {2.5, 2.5}}},
	{"strange", raycast.Poly{
		{0, 0}, {2.5, 2.5}, {0, 10}, {2.5, 7.5},
		{7.5, 7.5}, {10, 10}, {10, 0}, {2.5, 2.5}}},
	{"exagon", raycast.Poly{
		{3, 0}, {7, 0}, {10, 5}, {7, 10}, {3, 10}, {0, 5}}},
}

var tpt = []raycast.XY{{1, 2}, {2, 1}}

var tr = [][]bool{
	{true, true},
	{true, true},
	{false, false},
	{false, false},
}

func TestRC(t *testing.T) {
	for i, pg := range tpg {
		for j, pt := range tpt {
			if in := pt.In(pg.pg); in != tr[i][j] {
				t.Fatal(pg.name, pt, "got", in, "want", !in)
			}
		}
	}
}
