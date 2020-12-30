package waboorrt

import (
	"waboorrt/waboorrt/types"
)

type EntityType string

const (
	EntityTypeBot = "BOT"
)

type Entity struct {
	types.Pos
	Type EntityType `json:"type"`
}
