package waboorrt

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"

	"waboorrt/waboorrt/actions"
)

type Bot interface {
	NextAction(state *GameState) actions.Action
	Debug(msg string)
	EnableDebug()
	msgBuffer() []string
}

func Run(bot Bot) error {
	s := rpc.NewServer()
	c := newCustomCodec()
	s.RegisterCodec(c, "application/json")
	s.RegisterCodec(c, "application/json;charset=UTF-8")

	if err := s.RegisterService(newWrapper(bot), "Bot"); err != nil {
		return err
	}

	r := mux.NewRouter()
	r.Handle("/jsonrpc", s)

	return http.ListenAndServe(":4000", r)
}

type BotDebugger struct {
	messages []string
	debug    bool
}

func (d *BotDebugger) EnableDebug() {
	d.debug = true
}

func (d *BotDebugger) Debug(msg string) {
	if d.debug {
		d.messages = append(d.messages, msg)
	}
}

func (d *BotDebugger) msgBuffer() []string {
	ret := d.messages
	d.messages = []string{}
	return ret
}
