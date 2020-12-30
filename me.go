package waboorrt

import (
	"github.com/osiewers/waboorrt-go/types"
)

type MeInfo struct {
	types.Pos
	Coins     float64 `json:"coins"`
	ViewRange float64 `json:"view_range"`
	Name      string  `json:"name"`
	Health    float64 `json:"health"`
}
