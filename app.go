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
	Timers    []Timer
}

type Timer struct {
	Duration time.Duration
	Message  string
}

type StartDayResponse struct {
	Message   string              `json:"message"`
	TotalTime *database.TotalTime `json:"totalTime,omitempty"`
}

type MessageWorkTimeResponse struct {
	Message  string             `json:"message"`
	WorkTime *database.WorkTime `json:"workTime,omitempty"`
}

type MessageWorkTimeProjectResponse struct {
	Message         string                    `json:"message"`
	Project         *database.Project         `json:"project,omitempty"`
	WorkTimeProject *database.WorkTimeProject `json:"workTimeProject,omitempty"`
}

type MessageProjectResponse struct {
	Message string            `json:"message"`
	Project *database.Project `json:"project,omitempty"`
}

type TaskResponse struct {
	Message string         `json:"message"`
	Task    *database.Task `json:"task,omitempty"`
}

type ProjectsResponse struct {
	Projects     []database.Project `json:"projects"`
	Total        int64              `json:"total"`
	CurrentPage  int                `json:"currentPage"`
	ItemsPerPage int                `json:"itemsPerPage"`
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

func (a *App) StartTimer(seconds int, message string) string {
	if len(a.Timers) >= 3 {
		return "Cannot start more than 3 timers at the same time."
	}

	newTimer := time.Duration(seconds) * time.Second
	a.Timers = append(a.Timers, Timer{
		Duration: newTimer,
		Message:  message,
	})

	go func() {
		<-time.After(newTimer)

		// Emit the event when the timer finishes
		runtime.EventsEmit(a.ctx, "timerFinished", fmt.Sprintf("Reminder: '%s' finished after %d minutes.", message, seconds/60))

		// Show a dialog when the timer finishes
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Alert!",
			Message: message,
		})

		a.RemoveTimer(message)
	}()

	return "Timer started"
}

func (a *App) RemoveTimer(message string) {
	for i, t := range a.Timers {
		if t.Message == message {
			a.Timers = append(a.Timers[:i], a.Timers[i+1:]...)
			break
		}
	}
}

func (a *App) GetStartTimes() []Timer {
	return a.Timers
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

func (a *App) GetProjectByID(projectID int64) MessageProjectResponse {
	project, err := database.GetProject(projectID)
	if err != nil {
		log.Printf("Error while geting project: %v", err)
		return MessageProjectResponse{
			Message: "Project not found. Error",
		}
	}
	message := fmt.Sprintf("Project: %s", project.Name)
	return MessageProjectResponse{
		Message: message,
		Project: project,
	}

}

func (a *App) AssociateProjectToWorkTime(projectID int64) MessageWorkTimeProjectResponse {
	workTimeProject, err := database.AssociateProjectToWorkTime(projectID)
	if err != nil {
		log.Printf("Error while associating Project to Work Time: %v", err)
		return MessageWorkTimeProjectResponse{
			Message:         "Error while associating Project to Work Time",
			Project:         nil,
			WorkTimeProject: nil,
		}
	}
	project, err := database.GetProject(workTimeProject.ProjectID)
	if err != nil {
		log.Printf("Error while geting Project: %v", err)
		return MessageWorkTimeProjectResponse{
			Message:         "Error while geting Project",
			Project:         nil,
			WorkTimeProject: nil,
		}
	}

	message := fmt.Sprintf("Project %s associated to Work Time", project.Name)
	return MessageWorkTimeProjectResponse{
		Message:         message,
		Project:         project,
		WorkTimeProject: workTimeProject,
	}
}

func (a *App) GetAllProjects(page int, pageSize int) ProjectsResponse {
	projects, total, err := database.GetAllProjects(page, pageSize)
	if err != nil {
		log.Printf("Error retrieving projects: %v", err)
		return ProjectsResponse{
			Projects:     nil,
			Total:        0,
			CurrentPage:  page,
			ItemsPerPage: pageSize,
		}
	}

	return ProjectsResponse{
		Projects:     projects,
		Total:        total,
		CurrentPage:  page,
		ItemsPerPage: pageSize,
	}
}

func (a *App) CreateTask(projectID int64, description string, deadline string) TaskResponse {
	if deadline == "" {
		log.Println("No deadline provided, skipping task creation")
		return TaskResponse{
			Message: "Please provide a valid deadline.",
			Task:    nil,
		}
	}

	parsedDeadline, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		log.Printf("Error parsing deadline: %v", err)
		return TaskResponse{
			Message: "Invalid deadline format. Please use YYYY-MM-DD",
			Task:    nil,
		}
	}

	task, err := database.CreateTask(projectID, description, parsedDeadline)
	if err != nil {
		log.Printf("Error retrieving projects: %v", err)
		return TaskResponse{
			Message: "Error Creating Task",
			Task:    nil,
		}
	}

	return TaskResponse{
		Message: "Task created and associated with the project",
		Task:    task,
	}
}

func (a *App) GetUnfinishedWorkTimeProjectWithoutSendingError() *database.WorkTimeProject {
	workTimeProject, err := database.GetUnfinishedWorkTimeProject()
	if err != nil {
		log.Printf("Error retrieving workTimeProject: %v", err)
		return nil
	}
	return workTimeProject
}
