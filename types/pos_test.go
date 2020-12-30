package types

import (
	"testing"

	"waboorrt/waboorrt/actions"
)

func TestPos_CardinalDirection(t *testing.T) {
	p1 := Pos{0, 0}
	p2 := Pos{1, 0}
	p3 := Pos{0, 1}

	if p1.CardinalDirection(p2) != actions.WalkEast {
		t.Fail()
	}

	if p2.CardinalDirection(p1) != actions.WalkWest {
		t.Fail()
	}

	if p1.CardinalDirection(p3) != actions.WalkSouth {
		t.Fail()
	}

	if p3.CardinalDirection(p1) != actions.WalkNorth {
		t.Fail()
	}
}
