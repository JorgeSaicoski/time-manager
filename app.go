package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TotalTime struct {
	ID         int64
	StartTime  time.Time
	FinishTime time.Time
	WorkTimes  []*WorkTime
	BreakTime  *BreakTime
}

type Timer struct {
	ID      int64
	Time    time.Duration
	Message string
}

type WorkTime struct {
	ID        int64
	StartTime time.Time
	Duration  time.Duration
	Projects  []Project
	Brb       *Brb
}

type BreakTime struct {
	ID        int64
	StartTime time.Time
	Duration  time.Duration
}

type Project struct {
	ID        int64
	StartTime time.Time
	Duration  time.Duration
	Cost      *Cost
	WorkTimes []WorkTime
	Tasks     []Task
}

type Task struct {
	ID          int64
	Project     Project
	Deadline    time.Time
	Description string
}

type Cost struct {
	ID       int64
	Time     time.Time
	HourCost int
}

type Brb struct {
	ID        int64
	StartTime time.Time
	Duration  time.Duration
}

type App struct {
	ID        int64
	ctx       context.Context
	totalTime *TotalTime
	timer     *Timer
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartTotalTime() string {
	a.totalTime = &TotalTime{
		StartTime: time.Now(),
		WorkTimes: []*WorkTime{},
		BreakTime: &BreakTime{},
	}
	return "Day initialized. Have a good journey!"
}

func (a *App) StartTimer(seconds int, message string) {

	a.timer = &Timer{
		Time:    time.Duration(seconds) * time.Second,
		Message: message,
	}

	go func() {
		<-time.After(a.timer.Time)

		runtime.EventsEmit(a.ctx, "timerFinished", fmt.Sprintf("Reminder: '%s' finished after %d minutes.", message, seconds/60))

		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Alert!",
			Message: a.timer.Message,
		})

	}()

}
