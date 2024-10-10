package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"time-manager/database"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ID        int64
	ctx       context.Context
	TotalTime *database.TotalTime
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	database.Connect()
	database.MigrateDB()
	log.Println("Database connected and migrated successfully.")
}

func (a *App) FindByID(id int64, model string) (interface{}, string) {
	switch model {
	case "WorkTime":
		workTime, err := database.GetWorkTime(id)
		if err != nil {
			log.Printf("Error creating WorkTime: %v", err)
			return nil, "Work Time Not Found"
		}
		return workTime, "Work Time Found"
	case "Project":
		project, err := database.GetProject(id)
		if err != nil {
			log.Printf("Error creating project: %v", err)
			return nil, "Project Not Found"
		}
		return project, "Project Time Found"

	default:
		return nil, "no model found"
	}
}

func (a *App) StartDay() string {
	unfinishedTotalTime, err := database.GetUnfinishedTotalTime()
	if err != nil {
		log.Printf("Error checking unfinished TotalTime: %v", err)
		return "Error starting the day"
	}

	if unfinishedTotalTime != nil {
		return fmt.Sprintf("Unfinished day found with ID: %d", unfinishedTotalTime.ID)
	}

	totalTime, err := database.CreateTotalTime()
	if err != nil {
		log.Printf("Error creating TotalTime: %v", err)
		return "Error starting the day"
	}

	a.TotalTime = totalTime
	return "Day started"
}

func (a *App) StartTimer(seconds int, message string) string {

	newTimer := time.Duration(seconds) * time.Second

	go func() {
		<-time.After(newTimer)

		runtime.EventsEmit(a.ctx, "timerFinished", fmt.Sprintf("Reminder: '%s' finished after %d minutes.", message, seconds/60))

		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Alert!",
			Message: message,
		})

	}()

	return "Timer Started"

}

func (a *App) StartWorkTime() (string, *database.WorkTime) {
	newWorkTime, err := database.CreateWorkTime(a.TotalTime.ID)
	if err != nil {
		log.Printf("Error creating WorkTime: %v", err)
		return "Work Time not created", nil
	}
	return "Work Time! Good luck!", newWorkTime
}

func (a *App) TakeBreak(WorkTimeID int64) string {
	workTime, err := database.GetWorkTime(WorkTimeID)
	if err != nil {
		log.Printf("Error retrieving WorkTime: %v", err)
		return "Work Time Not Found"
	}

	endTime := time.Now()
	workTime.Duration = endTime.Sub(workTime.StartTime)

	if err := database.DB.Save(workTime).Error; err != nil {
		log.Printf("Error saving updated WorkTime: %v", err)
		return "Failed to update Work Time"
	}

	a.TotalTime.BreakTime.StartTime = time.Now()

	if err := database.DB.Save(a.TotalTime.BreakTime).Error; err != nil {
		log.Printf("Error saving updated BreakTime: %v", err)
		return "Failed to start Break"
	}

	return "Break started!"

}

func (a *App) EndBreak() (string, *database.WorkTime) {
	endTime := time.Now()

	breakDuration := endTime.Sub(a.TotalTime.BreakTime.StartTime)

	a.TotalTime.BreakTime.Duration += breakDuration

	newWorkTime, err := database.CreateWorkTime(a.TotalTime.ID)
	if err != nil {
		log.Printf("Error creating WorkTime: %v", err)
		return "Work Time not created", nil
	}

	message := fmt.Sprintf("Break ended! Total break time: %v", a.TotalTime.BreakTime.Duration)

	return message, newWorkTime

}
