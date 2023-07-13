package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type Hook interface {
	// Start the hook.
	WithSignals(signal ...syscall.Signal) Hook

	// Close register
	Close(funcs ...func())
}

type hook struct {
	ctx chan os.Signal
}

func NewHook() Hook {
	hook := &hook{
		ctx: make(chan os.Signal, 1),
	}
	return hook.WithSignals(syscall.SIGINT, syscall.SIGTERM)
}

func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, sig := range signals {
		signal.Notify(h.ctx, sig)
	}
	return h
}

func (h *hook) Close(funcs ...func()) {
	select {
	case <-h.ctx:
	}
	signal.Stop(h.ctx)
	for _, f := range funcs {
		f()
	}
}
