package app

import (
	"context"
)

type App struct {
	Name          string
	shutdownFuncs []OnShutdownFunc
}

type OnShutdownFunc func()

type Onstart func(ctx context.Context, app *App) error

func Start(onStart Onstart) {
	ctx := context.Background()

	a := &App{
		Name: "hello",
	}
	a.OnShutdown(func() {
		// do something
	})

	if err := onStart(ctx, a); err != nil {
		panic(err)
	}
}

func (a *App) OnShutdown(onShutdown func()) {
	a.shutdownFuncs = append([]OnShutdownFunc{onShutdown}, a.shutdownFuncs...)
}

func shutdown(ctx context.Context, a *App) {
	for _, shutdownFunc := range a.shutdownFuncs {
		shutdownFunc()
	}
}
