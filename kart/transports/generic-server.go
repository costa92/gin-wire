package transports

import (
	"context"
	"github.com/costa92/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
)

// ServerImp 公共接口
type ServerImp interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// AppInfo is application context value.
type AppInfo interface {
	ID() string
	Name() string
	Version() string
}

type GenericOption func(gs *GenericAPIServer)

func WithTransport(servers []ServerImp) GenericOption {
	return func(gs *GenericAPIServer) {
		gs.Servers = servers
	}
}

func WithSigs(signal []os.Signal) GenericOption {
	return func(gs *GenericAPIServer) {
		gs.sigs = signal
	}
}

// ID returns app instance id.
func (gs *GenericAPIServer) ID() string { return "" }

// Name returns service name.
func (gs *GenericAPIServer) Name() string { return "" }

// Version returns app version.
func (gs *GenericAPIServer) Version() string { return "" }

type GenericAPIServer struct {
	Servers []ServerImp
	cancel  func()
	sigs    []os.Signal
	ctx     context.Context
}

// NewGenericAPIServer 实例化
func NewGenericAPIServer(opts ...GenericOption) *GenericAPIServer {
	gs := &GenericAPIServer{}
	for _, o := range opts {
		o(gs)
	}
	bgCtx := context.Background()
	gs.ctx, gs.cancel = context.WithCancel(bgCtx)
	return gs
}

// Run 开始运行
func (gs *GenericAPIServer) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	wg := sync.WaitGroup{}
	for _, server := range gs.Servers {
		srv := server
		eg.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}

	wg.Wait()
	c := make(chan os.Signal, 1)
	signal.Notify(c, gs.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return gs.Stop()
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (gs *GenericAPIServer) Stop() error {
	if gs.cancel != nil {
		gs.cancel()
	}
	return nil
}

type appKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, s AppInfo) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppInfo)
	return
}
