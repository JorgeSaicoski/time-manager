package main

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) TestMessage() string {
	return "Test"
}

func (a *App) StartTimer(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)

	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Timer",
		Message: "Time finish",
	})
}
