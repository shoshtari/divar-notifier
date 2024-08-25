package jarchi

import (
	"github.com/shoshtari/divar-notifier/internal/divar"
	"github.com/shoshtari/divar-notifier/internal/notify"
)

type Jarchi interface {
	Start() error
	Stop() error
}

type JarchiImp struct {
}

// Start implements Jarchi.
func (j JarchiImp) Start() error {
	panic("unimplemented")
}

// Stop implements Jarchi.
func (j JarchiImp) Stop() error {
	panic("unimplemented")
}

func New(divarClient divar.DivarClient, notifier notify.Notifier) Jarchi {
	return JarchiImp{}
}
