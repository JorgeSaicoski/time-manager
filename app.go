package main

import (
	"context"
	"time"
)

type TotalTime struct {
	StartTime  time.Time
	FinishTime time.Time
	WorkTimes  []*WorkTime
	BreakTime  *BreakTime
}

type Timer struct {
	Time    time.Time
	Message string
}

type WorkTime struct {
	StartTime time.Time
	Duration  time.Duration
	Projects  []Project
	Brb       *Brb
}

type BreakTime struct {
	StartTime time.Time
	Duration  time.Duration
}

type Project struct {
	StartTime time.Time
	Duration  time.Duration
	Cost      *Cost
	WorkTimes []WorkTime
	Tasks     []Task
}

type Task struct {
	Project     Project
	Deadline    time.Time
	Description string
}

type Cost struct {
	Time     time.Time
	HourCost int
}

type Brb struct {
	StartTime time.Time
	Duration  time.Duration
}

type App struct {
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
