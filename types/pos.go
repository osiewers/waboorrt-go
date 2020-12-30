package types

import (
	"math"

	"github.com/osiewers/waboorrt-go/actions/constants"
)

type Pos struct {
	X float64
	Y float64
}

const WalkNone constants.WalkDirection = "NONE"

func (p Pos) Equals(t Pos) bool {
	return p.X == t.X && p.Y == t.Y
}

func (p Pos) Distance(t Pos) float64 {
	return math.Sqrt(math.Pow(math.Abs(p.X-t.X), 2) + math.Pow(math.Abs(p.Y-t.Y), 2))
}

func (p Pos) CardinalDirection(t Pos) constants.WalkDirection {
	dx := p.X - t.X
	dy := p.Y - t.Y

	if math.Abs(dx) < 1 && math.Abs(dy) < 1 {
		return WalkNone
	}

	if math.Abs(dx) >= math.Abs(dy) {
		if dx < 0 {
			return constants.WalkEast
		} else {
			return constants.WalkWest
		}
	} else {
		if dy < 0 {
			return constants.WalkSouth
		} else {
			return constants.WalkNorth
		}
	}
}
