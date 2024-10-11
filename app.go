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

type StartDayResponse struct {
	Message   string              `json:"message"`
	TotalTime *database.TotalTime `json:"totalTime,omitempty"`
}

type MessageWorkTimeResponse struct {
	Message  string             `json:"message"`
	WorkTime *database.WorkTime `json:"workTime,omitempty"`
}

type MessageProjectResponse struct {
	Message string            `json:"message"`
	Project *database.Project `json:"project,omitempty"`
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

func (a *App) StartDay() StartDayResponse {
	unfinishedTotalTime, err := database.GetUnfinishedTotalTime()
	if err != nil {
		log.Printf("Error checking unfinished TotalTime: %v", err)
		return StartDayResponse{
			Message: "Error checking for unfinished days",
		}
	}

	if unfinishedTotalTime != nil {
		a.TotalTime = unfinishedTotalTime
		return StartDayResponse{
			Message:   "Unfinished day found",
			TotalTime: unfinishedTotalTime,
		}
	}

	totalTime, err := database.CreateTotalTime()
	if err != nil {
		log.Printf("Error creating TotalTime: %v", err)
		return StartDayResponse{
			Message: "Error starting the day",
		}
	}

	a.TotalTime = totalTime
	return StartDayResponse{
		Message:   "Day started",
		TotalTime: totalTime,
	}
}

func (a *App) FinishDay() string {
	unfinishedTotalTime, err := database.GetUnfinishedTotalTime()
	database.FinishWorkTime()
	if err != nil {
		log.Printf("Error checking unfinished TotalTime: %v", err)
		return "Error checking for unfinished days"
	}
	if unfinishedTotalTime == nil {
		log.Printf("None unfinished time found to close")
		return "None unfinished time found to close"
	}

	if _, err := database.FinishTotalTime(unfinishedTotalTime.ID); err != nil {
		log.Printf("Error finishing TotalTime: %v", err)
		return "error finishing Total Time"
	}

	return "Total Time finished"
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

func (a *App) StartWorkTime() MessageWorkTimeResponse {
	newWorkTime, err := database.CreateWorkTime(a.TotalTime.ID)
	if err != nil {
		log.Printf("Error creating WorkTime: %v", err)
		return MessageWorkTimeResponse{
			Message: "Error creating work time",
		}
	}
	return MessageWorkTimeResponse{
		Message:  "Work Time!",
		WorkTime: newWorkTime,
	}
}

func (a *App) TakeBreak() string {
	_, err := database.FinishWorkTime()
	if err != nil {
		log.Printf("Error retrieving WorkTime: %v", err)
		return "Work Time Not Found"
	}

	if a.TotalTime.BreakTime == nil {
		log.Printf("a.TotalTime.BreakTime is nil, initializing new BreakTime")
		a.TotalTime.BreakTime = &database.BreakTime{}
	}

	a.TotalTime.BreakTime.StartTime = time.Now()

	a.TotalTime.BreakTime.TotalTimeID = a.TotalTime.ID

	if err := database.DB.Save(a.TotalTime.BreakTime).Error; err != nil {
		log.Printf("Error saving updated BreakTime: %v", err)
		return "Failed to start Break"
	}

	return "Break started!"

}

func (a *App) EndBreak() MessageWorkTimeResponse {
	endTime := time.Now()

	breakDuration := endTime.Sub(a.TotalTime.BreakTime.StartTime)

	a.TotalTime.BreakTime.Duration += breakDuration

	newWorkTime, err := database.CreateWorkTime(a.TotalTime.ID)
	if err != nil {
		log.Printf("Error creating WorkTime: %v", err)
		return MessageWorkTimeResponse{
			Message: "Work time not created. Error",
		}
	}

	message := fmt.Sprintf("Break ended! Total break time: %v", a.TotalTime.BreakTime.Duration)

	return MessageWorkTimeResponse{
		Message:  message,
		WorkTime: newWorkTime,
	}

}

func (a *App) CreateProject(name string) MessageProjectResponse {
	newProject, err := database.CreateProject(name)
	if err != nil {
		log.Printf("Error creating WorkTime: %v", err)
		return MessageProjectResponse{
			Message: "Project not created. Error",
		}
	}

	message := fmt.Sprintf("Project Created: %s", newProject.Name)
	return MessageProjectResponse{
		Message: message,
		Project: newProject,
	}

}

func (a *App) AssociateProjectToWorkTime(projectId int64) string {
	workTimeProject, err := database.AssociateProjectToWorkTime(projectId)
	if err != nil {
		log.Printf("Error while associating Project to Work Time: %v", err)
		return "Error while associating Project to Work Time"
	}

	message := fmt.Sprintf("Project %s associated to Work Time", workTimeProject.Project.Name)
	return message
}
