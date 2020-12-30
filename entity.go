package waboorrt

import (
	"github.com/osiewers/waboorrt-go/types"
)

type EntityType string

const (
	EntityTypeBot = "BOT"
)

type Entity struct {
	types.Pos
	Type EntityType `json:"type"`
}
