package waboorrt

import (
	"waboorrt/waboorrt/types"
)

type MeInfo struct {
	types.Pos
	Coins     float64 `json:"coins"`
	ViewRange float64 `json:"view_range"`
	Name      string  `json:"name"`
	Health    float64 `json:"health"`
}
