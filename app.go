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
	Duration   time.Duration
	Message    string
	FinishTime time.Time
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

type MessageCostResponse struct {
	Message string         `json:"message"`
	Cost    *database.Cost `json:"cost,omitempty"`
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

type DaySummary struct {
	WorkTimesStarted      []database.WorkTime        `json:"workTimesStarted"`
	WorkTimesCrossingDays []database.WorkTime        `json:"workTimesCrossingDays"`
	TotalTime             time.Duration              `json:"totalTime"`
	WorkTimeProjects      []database.WorkTimeProject `json:"workTimeProjects"`
	Breaks                []database.BreakTime       `json:"breaks"`
	Brbs                  []database.Brb             `json:"brbs"`
	Message               string                     `json:"message"`
}

type CurrentTimersResponse struct {
	WorkTime  *database.WorkTime  `json:"workTime,omitempty"`
	BreakTime *database.BreakTime `json:"breakTime,omitempty"`
	Brb       *database.Brb       `json:"brb,omitempty"`
	TotalTime *database.TotalTime `json:"totalTime,omitempty"`
}

type ResolutionMessageResponse struct {
	Message string                      `json:"message"`
	Tracker *database.ResolutionTracker `json:"tracker,omitempty"`
	Units   []database.ResolutionUnit   `json:"units,omitempty"`
	Unit    *database.ResolutionUnit    `json:"unit,omitempty"`
	Success bool                        `json:"success"`
}

type SearchResultResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Results []database.SearchResult `json:"results"`
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
	for _, timer := range a.Timers {
		if timer.Message == message {
			return "A timer with this message already exists."
		}
	}

	if len(a.Timers) >= 3 {
		return "Cannot start more than 3 timers at the same time."
	}

	newTimer := time.Duration(seconds) * time.Second

	finishTime := time.Now().Add(newTimer)

	a.Timers = append(a.Timers, Timer{
		Duration:   newTimer,
		Message:    message,
		FinishTime: finishTime,
	})

	go func() {
		<-time.After(newTimer)

		runtime.EventsEmit(a.ctx, "timerFinished", fmt.Sprintf("Reminder: '%s' finished after %d minutes.", message, seconds/60))

		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Alert!",
			Message: message,
		})

		a.RemoveTimer(message)
	}()

	return fmt.Sprintf("Timer started and will finish at %s", finishTime.Format("15:04:05"))
}

func (a *App) RemoveTimer(message string) string {
	for i, t := range a.Timers {
		if t.Message == message {
			a.Timers = append(a.Timers[:i], a.Timers[i+1:]...)
			break
		}
	}

	return "Timer removed"
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
	if err != nil {
		log.Printf("Error checking unfinished TotalTime: %v", err)
		return "Error checking for unfinished days"
	}
	if unfinishedTotalTime == nil {
		log.Printf("None unfinished time found to close")
		return "None unfinished time found to close"
	}
	_, err = database.FinishWorkTime()
	if err != nil {
		log.Printf("Error finishing WorkTime: %v", err)
		return "Error finishing WorkTime"
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

	if err := database.SaveTotalTime(a.TotalTime); err != nil {
		log.Printf("Error saving current TotalTime: %v", err)
		return MessageWorkTimeResponse{
			Message:  "No Current Total Time, can't start work time",
			WorkTime: nil,
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

	a.TotalTime.BreakTime.Active = true

	a.TotalTime.BreakTime.TotalTimeID = a.TotalTime.ID

	if err := database.DB.Save(a.TotalTime.BreakTime).Error; err != nil {
		log.Printf("Error saving updated BreakTime: %v", err)
		return "Failed to start Break"
	}

	if err := database.SaveTotalTime(a.TotalTime); err != nil {
		log.Printf("Error saving current TotalTime: %v", err)
		return "No Current Total Time, can't take a Break"
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

	if err := database.SaveTotalTime(a.TotalTime); err != nil {
		log.Printf("Error saving current TotalTime: %v", err)
		return MessageWorkTimeResponse{
			Message:  "Error finding the current TotalTime and saving it",
			WorkTime: nil,
		}
	}

	message := fmt.Sprintf("Break ended! Total break time: %v", a.TotalTime.BreakTime.Duration)

	return MessageWorkTimeResponse{
		Message:  message,
		WorkTime: newWorkTime,
	}

}

func (a *App) TakeBrb() string {
	_, err := database.FinishWorkTime()
	if err != nil {
		log.Printf("Error retrieving WorkTime: %v", err)
		return "Work Time Not Found"
	}

	if a.TotalTime.Brb == nil {
		log.Printf("a.TotalTime.Brb is nil, initializing new BreakTime")
		a.TotalTime.Brb = &database.Brb{}
	}

	a.TotalTime.Brb.StartTime = time.Now()

	a.TotalTime.Brb.Active = true

	a.TotalTime.Brb.TotalTimeID = a.TotalTime.ID

	if err := database.DB.Save(a.TotalTime.Brb).Error; err != nil {
		log.Printf("Error saving updated Brb: %v", err)
		return "Failed to start Break"
	}

	if err := database.SaveTotalTime(a.TotalTime); err != nil {
		log.Printf("Error saving current TotalTime: %v", err)
		return "No Current Total Time, can't take a BRB"
	}

	return "Brb started!"

}

func (a *App) EndBrb() MessageWorkTimeResponse {
	endTime := time.Now()

	breakDuration := endTime.Sub(a.TotalTime.Brb.StartTime)

	a.TotalTime.Brb.Duration += breakDuration

	newWorkTime, err := database.CreateWorkTime(a.TotalTime.ID)
	if err != nil {
		log.Printf("Error creating WorkTime: %v", err)
		return MessageWorkTimeResponse{
			Message: "Work time not created. Error",
		}
	}

	if err := database.SaveTotalTime(a.TotalTime); err != nil {
		log.Printf("Error saving current TotalTime: %v", err)
		return MessageWorkTimeResponse{
			Message:  "Error finding the current TotalTime and saving it",
			WorkTime: nil,
		}
	}

	message := fmt.Sprintf("Brb ended! Total break time: %v", a.TotalTime.Brb.Duration)

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

func (a *App) ChangeProjectClose(projectID int64) MessageProjectResponse {
	project, err := database.ChangeProjectClose(projectID)
	if err != nil {
		log.Printf("Error while changing project close: %v", err)
		return MessageProjectResponse{
			Message: "Project close not changed. Error",
		}
	}
	message := fmt.Sprintf("Project %s closed changed", project.Name)
	return MessageProjectResponse{
		Message: message,
		Project: project,
	}

}

func (a *App) UpdateProjectName(projectID int64, newName string) MessageProjectResponse {
	project, err := database.ChangeProjectName(projectID, newName)
	if err != nil {
		log.Printf("Error while changing project name: %v", err)
		return MessageProjectResponse{
			Message: "Project name not changed. Error",
		}
	}
	message := fmt.Sprintf("Project %s name changed", project.Name)
	return MessageProjectResponse{
		Message: message,
		Project: project,
	}
}

func (a *App) CalculateAndSaveProjectCost(projectID int64, hourCost int) MessageCostResponse {
	cost, err := database.SaveCost(projectID, hourCost)
	if err != nil {
		log.Printf("Error while saving project cost: %v", err)
		return MessageCostResponse{
			Message: "Project cost not saved. Error",
		}
	}
	message := fmt.Sprintf("Project cost saved: %d", cost.HourCost)
	return MessageCostResponse{
		Message: message,
		Cost:    cost,
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

func (a *App) GetAllProjects(page int, pageSize int, closedFilter *bool, orderBy string, orderDirection string) ProjectsResponse {
	projects, total, err := database.GetAllProjects(page, pageSize, closedFilter, orderBy, orderDirection)
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

func (a *App) DeleteTask(taskID int64) TaskResponse {
	err := database.DeleteTask(taskID)
	if err != nil {
		return TaskResponse{
			Message: "Failed to delete Task.",
		}
	}

	return TaskResponse{
		Message: "Task deleted successfully.",
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

func (a *App) CalculateWorkTimeForDay(day time.Time) (time.Duration, error) {
	workTimes, err := database.GetWorkTimesForDay(day)
	if err != nil {
		log.Printf("Error fetching work times: %v", err)
		return 0, err
	}

	var totalDuration time.Duration

	for _, workTime := range workTimes {
		totalDuration += workTime.Duration

		if totalDuration > 24*time.Hour {
			err := database.UpdateWorkTimeTrustworthy(&workTime, false)
			if err != nil {
				log.Printf("Error updating trustworthy status: %v", err)
			}
			totalDuration = 24 * time.Hour
		}
	}

	return totalDuration, nil
}

func (a *App) GetDaySummary(dayString string) DaySummary {
	day, err := time.Parse("2006-01-02", dayString)
	if err != nil {
		log.Printf("Error parsing day: %v", err)
		return DaySummary{
			Message: "Invalid date format",
		}
	}

	var summary DaySummary

	workTimes, err := database.GetWorkTimesForDay(day)
	if err != nil {
		log.Printf("Error fetching work times: %v", err)
		summary.Message = "Error fetching work times"
		return summary
	}
	summary.WorkTimesStarted = workTimes

	workTimesCrossingDays, err := database.GetWorkTimesCrossingDays(day)
	if err != nil {
		log.Printf("Error fetching work times crossing days: %v", err)
		summary.Message = "Error fetching work times crossing days"
		return summary
	}
	summary.WorkTimesCrossingDays = workTimesCrossingDays

	var totalTime time.Duration
	for _, workTime := range workTimes {
		totalTime += workTime.Duration
	}
	if totalTime > 24*time.Hour {
		totalTime = 24 * time.Hour
	}
	summary.TotalTime = totalTime

	var workTimeProjects []database.WorkTimeProject
	for _, workTime := range workTimes {
		wtpList, err := database.GetWorkTimeProjectsByWorkTimeID(workTime.ID)
		if err != nil {
			log.Printf("Error fetching WorkTimeProjects for WorkTime %d: %v", workTime.ID, err)
			continue
		}
		workTimeProjects = append(workTimeProjects, wtpList...)
	}

	summary.WorkTimeProjects = workTimeProjects

	breakTimes, err := database.GetBreakTimesForDay(day)
	if err != nil {
		log.Printf("Error fetching break times: %v", err)
		summary.Message = "Error fetching break times"
		return summary
	}
	summary.Breaks = breakTimes

	brbs, err := database.GetBrbsForDay(day)
	if err != nil {
		log.Printf("Error fetching brb sessions: %v", err)
		summary.Message = "Error fetching brb sessions"
		return summary
	}
	summary.Brbs = brbs

	summary.Message = "Day summary successfully fetched"
	return summary

}

func (a *App) UpdateWorkTimeDuration(workTimeID int64, newDurationSeconds int64) MessageWorkTimeResponse {
	workTime, err := database.GetWorkTime(workTimeID)
	if err != nil {
		log.Printf("Error retrieving WorkTime: %v", err)
		return MessageWorkTimeResponse{
			Message: "Work time not found",
		}
	}

	workTime.Duration = time.Duration(newDurationSeconds) * time.Second

	if err := database.DB.Omit("Projects").Save(workTime).Error; err != nil {
		log.Printf("Error updating WorkTime duration: %v", err)
		return MessageWorkTimeResponse{
			Message: "Failed to update WorkTime duration",
		}
	}

	message := fmt.Sprintf("WorkTime duration updated to %v", workTime.Duration)
	return MessageWorkTimeResponse{
		Message:  message,
		WorkTime: workTime,
	}
}

func (a *App) UpdateBreakTimeDuration(breakTimeID int64, newDurationSeconds int64) string {
	breakTime, err := database.GetBreakTime(breakTimeID)
	if err != nil {
		log.Printf("Error retrieving BreakTime: %v", err)
		return "Break time not found"
	}

	breakTime.Duration = time.Duration(newDurationSeconds) * time.Second

	if err := database.DB.Save(breakTime).Error; err != nil {
		log.Printf("Error updating BreakTime duration: %v", err)
		return "Failed to update BreakTime duration"
	}

	return fmt.Sprintf("BreakTime duration updated to %v", breakTime.Duration)
}

func (a *App) UpdateBrbDuration(brbID int64, newDurationSeconds int64) string {
	brb, err := database.GetBrb(brbID)
	if err != nil {
		log.Printf("Error retrieving Brb: %v", err)
		return "Brb time not found"
	}

	brb.Duration = time.Duration(newDurationSeconds) * time.Second

	if err := database.DB.Save(brb).Error; err != nil {
		log.Printf("Error updating Brb duration: %v", err)
		return "Failed to update Brb duration"
	}

	return fmt.Sprintf("Brb duration updated to %v", brb.Duration)
}

func (a *App) UpdateWorkTimeProjectDuration(workTimeProjectID int64, newDurationSeconds int64) MessageWorkTimeProjectResponse {
	workTimeProject, err := database.GetWorkTimeProjectByID(workTimeProjectID)
	if err != nil {
		log.Printf("Error retrieving WorkTimeProject: %v", err)
		return MessageWorkTimeProjectResponse{
			Message: "WorkTimeProject not found",
		}
	}

	workTimeProject.Duration = time.Duration(newDurationSeconds) * time.Second

	if err := database.DB.Save(workTimeProject).Error; err != nil {
		log.Printf("Error updating WorkTimeProject duration: %v", err)
		return MessageWorkTimeProjectResponse{
			Message: "Failed to update WorkTimeProject duration",
		}
	}

	message := fmt.Sprintf("WorkTimeProject duration updated to %v", workTimeProject.Duration)
	return MessageWorkTimeProjectResponse{
		Message:         message,
		WorkTimeProject: workTimeProject,
	}
}

func (a *App) GetCurrentActiveTimers() (*CurrentTimersResponse, error) {
	currentTimers, err := database.GetCurrentActiveTimers()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch active timers: %w", err)
	}

	return &CurrentTimersResponse{
		WorkTime:  currentTimers.WorkTime,
		BreakTime: currentTimers.BreakTime,
		Brb:       currentTimers.Brb,
		TotalTime: currentTimers.TotalTime,
	}, nil
}

func (a *App) GetOrCreateTodayResolutionTracker() ResolutionMessageResponse {
	tracker, err := database.GetOrCreateTodayResolutionTracker()
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to retrieve or create today's tracker.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "Successfully retrieved today's tracker.",
		Tracker: tracker,
		Success: true,
	}
}

func (a *App) UpdateResolutionTrackerCategory(trackerID int64, newCategory string) ResolutionMessageResponse {
	err := database.UpdateResolutionTrackerCategory(trackerID, newCategory)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to update tracker category.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "Category updated successfully.",
		Success: true,
	}
}

func (a *App) CloseResolutionTracker(trackerID int64) ResolutionMessageResponse {
	err := database.CloseResolutionTracker(trackerID)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to close tracker.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "Tracker closed successfully.",
		Success: true,
	}
}

func (a *App) FindResolutionTrackerByDay(day string) ResolutionMessageResponse {
	parsedDay, err := time.Parse("2006-01-02", day)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Invalid date format, use YYYY-MM-DD.",
			Success: false,
		}
	}

	tracker, err := database.FindOrCreateResolutionTrackerByDay(parsedDay)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "No tracker found for the given day.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "Tracker found successfully.",
		Tracker: tracker,
		Success: true,
	}
}

func (a *App) CreateResolutionUnit(identifier string) ResolutionMessageResponse {
	unit, err := database.CreateResolutionUnit(identifier)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to create ResolutionUnit.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "ResolutionUnit created successfully.",
		Unit:    unit,
		Success: true,
	}
}

func (a *App) GetUnitsByResolutionTracker(trackerID int64) ResolutionMessageResponse {
	units, err := database.GetUnitsByResolutionTracker(trackerID)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to retrieve ResolutionUnits.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "ResolutionUnits retrieved successfully.",
		Units:   units,
		Success: true,
	}
}

func (a *App) GetUnitsTrackerByDay(dayString string) ResolutionMessageResponse {
	parsedDay, err := time.Parse("2006-01-02", dayString)
	if err != nil {
		log.Printf("Error parsing day: %v", err)
		return ResolutionMessageResponse{
			Message: "Invalid date format",
			Success: false,
		}
	}

	tracker, err := database.FindOrCreateResolutionTrackerByDay(parsedDay)
	if err != nil {
		log.Printf("Error fetching Tracker By Day: %v", err)
		return ResolutionMessageResponse{
			Message: "Error fetching tracker for the day",
			Success: false,
		}
	}

	if tracker == nil {
		log.Printf("Null: %v", parsedDay)
		return ResolutionMessageResponse{
			Message: "No tracker data available for this day.",
			Units:   nil,
			Success: true,
		}
	}

	units, err := database.GetUnitsByResolutionTracker(tracker.ID)
	if err != nil {
		log.Printf("Error fetching units: %v", err)
		return ResolutionMessageResponse{
			Message: "Error fetching units for the day",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "Units retrieved successfully",
		Units:   units,
		Success: true,
	}
}

func (a *App) CreateUnitByDay(identifier string, dayString string) ResolutionMessageResponse {
	parsedDay, err := time.Parse("2006-01-02", dayString)
	if err != nil {
		log.Printf("Error parsing day: %v", err)
		return ResolutionMessageResponse{
			Message: "Invalid date format",
			Success: false,
		}
	}

	unit, err := database.CreateResolutionUnitByDay(identifier, parsedDay)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to create ResolutionUnit.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "ResolutionUnit created successfully.",
		Unit:    unit,
		Success: true,
	}
}

func (a *App) DeleteResolutionUnit(unitID int64) ResolutionMessageResponse {
	err := database.DeleteResolutionUnit(unitID)
	if err != nil {
		return ResolutionMessageResponse{
			Message: "Failed to delete ResolutionUnit.",
			Success: false,
		}
	}

	return ResolutionMessageResponse{
		Message: "ResolutionUnit deleted successfully.",
		Success: true,
	}
}

func (a *App) Search(term string) SearchResultResponse {
	results, err := database.SearchItems(term)
	if err != nil {
		return SearchResultResponse{
			Success: false,
			Message: err.Error(),
			Results: nil,
		}
	}

	return SearchResultResponse{
		Success: true,
		Message: "Search completed successfully.",
		Results: results,
	}
}
