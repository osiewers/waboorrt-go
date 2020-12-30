package types

import (
	"testing"

	"github.com/osiewers/waboorrt-go/actions/constants"
)

func TestPos_CardinalDirection(t *testing.T) {
	p1 := Pos{0, 0}
	p2 := Pos{1, 0}
	p3 := Pos{0, 1}

	if p1.CardinalDirection(p2) != constants.WalkEast {
		t.Fail()
	}

	if p2.CardinalDirection(p1) != constants.WalkWest {
		t.Fail()
	}

	if p1.CardinalDirection(p3) != constants.WalkSouth {
		t.Fail()
	}

	if p3.CardinalDirection(p1) != constants.WalkNorth {
		t.Fail()
	}
}
