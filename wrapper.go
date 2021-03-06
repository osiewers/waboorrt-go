package waboorrt

import (
	"net/http"

	"github.com/osiewers/waboorrt-go/actions"
)

type wrapper struct {
	bot Bot
}

func newWrapper(bot Bot) *wrapper {
	return &wrapper{
		bot: bot,
	}
}

func (w *wrapper) NextAction(r *http.Request, args *GameState, result *actions.Action) error {
	ret := w.bot.NextAction(args)
	ret.SetMessageBuffer(w.bot.msgBuffer())
	*result = ret
	return nil
}

func (w *wrapper) Health(r *http.Request, args *interface{}, result *bool) error {
	*result = true
	return nil
}
