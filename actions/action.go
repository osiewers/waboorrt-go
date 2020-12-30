package actions

import (
	"waboorrt/waboorrt/actions/constants"
	"waboorrt/waboorrt/types"
)

type Action interface {
	SetMessageBuffer(messages []string)
}

type ActionImpl struct {
	Name     constants.ActionType `json:"name"`
	Messages []string   `json:"messages"`
}

func (a *ActionImpl) SetMessageBuffer(messages []string) {
	a.Messages = messages
}

type NoOp struct {
	ActionImpl
}

func NewNoOp() Action {
	return &NoOp{
		ActionImpl{
			Name: constants.ActionNoop,
		},
	}
}

type WalkOp struct {
	ActionImpl

	Direction constants.WalkDirection `json:"direction"`
}

func NewWalkOp(dir constants.WalkDirection) Action {
	return &WalkOp{
		ActionImpl: ActionImpl{
			Name: constants.ActionWalk,
		},
		Direction: dir,
	}
}

type ThrowOp struct {
	ActionImpl

	types.Pos
}

func NewThrowOp(p types.Pos) Action {
	return &ThrowOp{
		ActionImpl: ActionImpl{
			Name: constants.ActionThrow,
		},
		Pos: p,
	}
}

type LookOp struct {
	ActionImpl

	Range float64 `json:"range"`
}

func NewLookOp(lookRange float64) Action {
	return &LookOp{
		ActionImpl: ActionImpl{
			Name: constants.ActionLook,
		},
		Range: lookRange,
	}
}
