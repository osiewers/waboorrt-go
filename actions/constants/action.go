package constants

type ActionType string

const (
	ActionNoop  ActionType = "NOOP"
	ActionWalk  ActionType = "WALK"
	ActionThrow ActionType = "THROW"
	ActionLook  ActionType = "LOOK"
)

type WalkDirection string

const (
	WalkNorth WalkDirection = "NORTH"
	WalkWest  WalkDirection = "WEST"
	WalkSouth WalkDirection = "SOUTH"
	WalkEast  WalkDirection = "EAST"
)
