package database

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("time_manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func updateProjectDuration(projectID int64) error {
	var project Project

	err := DB.First(&project, projectID).Error
	if err != nil {
		return fmt.Errorf("failed to retrieve project: %w", err)
	}

	var totalDuration int64
	err = DB.Model(&WorkTimeProject{}).
		Where("project_id = ?", projectID).
		Select("COALESCE(SUM(duration), 0)").
		Scan(&totalDuration).Error
	if err != nil {
		return fmt.Errorf("failed to calculate total duration: %w", err)
	}

	project.Duration = time.Duration(totalDuration)

	if err := DB.Save(&project).Error; err != nil {
		return fmt.Errorf("failed to update project duration: %w", err)
	}

	return nil
}

func getUnfinishedWorkTime() (*WorkTime, error) {
	var workTime WorkTime
	result := DB.Where("closed = ?", false).First(&workTime)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("No unfinished WorkTime found.")
			return nil, nil
		}
		return nil, result.Error
	}
	return &workTime, nil
}

func getUnfinishedWorkTimeProjectAndFinish() (*WorkTimeProject, error) {
	var workTimeProject WorkTimeProject
	result := DB.Where("closed = ?", false).First(&workTimeProject)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) || strings.Contains(result.Error.Error(), "no such column: closed") {
			log.Println("No unfinished WorkTimeProject found or column 'closed' missing.")
			return nil, nil
		}
	}
	workTimeProject.Closed = true
	endTime := time.Now()
	workTimeProject.Duration = endTime.Sub(workTimeProject.StartTime)

	if err := DB.Save(&workTimeProject).Error; err != nil {
		return nil, fmt.Errorf("failed to update active WorkTimeProject: %w", err)
	}

	err := updateProjectDuration(workTimeProject.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to update project duration: %w", err)
	}

	return &workTimeProject, nil
}

func CreateTotalTime() (*TotalTime, error) {
	totalTime := &TotalTime{
		StartTime: time.Now(),
		Closed:    false,
	}
	err := DB.Create(totalTime).Error
	return totalTime, err
}

func GetTotalTime(id int64) (*TotalTime, error) {
	var totalTime TotalTime
	err := DB.Preload("WorkTimes").Preload("BreakTime").First(&totalTime, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("TotalTime not found")
	}
	return &totalTime, err
}

func GetUnfinishedTotalTime() (*TotalTime, error) {
	var totalTime TotalTime
	result := DB.Where("closed = ?", false).First(&totalTime)
	if result.Error != nil {
		fmt.Println("109")
		fmt.Println(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("No unfinished TotalTime found.")
			return nil, nil
		}
		return nil, result.Error
	}

	return &totalTime, nil
}

func FinishTotalTime(id int64) (*TotalTime, error) {
	var totalTime TotalTime

	err := DB.First(&totalTime, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("TotalTime not found")
	} else if err != nil {
		return nil, fmt.Errorf("error retrieving TotalTime: %w", err)
	}

	totalTime.FinishTime = time.Now()
	totalTime.Closed = true

	_, err = FinishWorkTime()
	if err != nil {
		return nil, fmt.Errorf("error finishing WorkTime: %w", err)
	}

	if err := DB.Save(&totalTime).Error; err != nil {
		return nil, fmt.Errorf("failed to finish TotalTime: %w", err)
	}

	return &totalTime, nil
}

func CreateWorkTime(totalTimeID int64) (*WorkTime, error) {
	workTime, _ := getUnfinishedWorkTime()

	if workTime != nil {
		return workTime, nil
	}

	workTime = &WorkTime{
		StartTime:   time.Now(),
		TotalTimeID: totalTimeID,
	}

	err := DB.Create(workTime).Error
	if err != nil {
		return nil, err
	}

	var totalTime TotalTime
	if err := DB.Preload("WorkTimes").First(&totalTime, totalTimeID).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve TotalTime after creating WorkTime: %w", err)
	}

	totalTime.WorkTimes = append(totalTime.WorkTimes, *workTime)

	err = DB.Save(&totalTime).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update TotalTime with new WorkTime: %w", err)
	}

	return workTime, nil
}

func GetWorkTime(id int64) (*WorkTime, error) {
	var workTime WorkTime
	err := DB.Preload("Projects").First(&workTime, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("WorkTime not found")
	}
	return &workTime, err
}

func FinishWorkTime() (*WorkTime, error) {
	workTime, err := getUnfinishedWorkTime()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve unfinished WorkTime: %w", err)
	}

	if workTime == nil {
		return nil, nil
	}
	workTime.Closed = true
	endTime := time.Now()
	workTime.Duration = endTime.Sub(workTime.StartTime)
	if err := DB.Save(&workTime).Error; err != nil {
		return nil, fmt.Errorf("failed to finish TotalWork: %w", err)
	}

	workTimeProject, err := getUnfinishedWorkTimeProjectAndFinish()
	if err != nil {
		return nil, fmt.Errorf("failed to finish WorkTimeProject: %w", err)
	}

	if workTimeProject == nil {
		log.Println("No unfinished WorkTimeProject found.")
	}

	return workTime, nil
}

func CreateProject(name string) (*Project, error) {

	project := &Project{
		Name:      name,
		StartTime: time.Now(),
	}

	if err := DB.Create(project).Error; err != nil {
		return nil, fmt.Errorf("failed to create Project: %w", err)
	}

	return project, nil
}

func AssociateProjectToWorkTime(projectID int64) (*WorkTimeProject, error) {
	unfinishedWorkTimeProject, err := GetUnfinishedWorkTimeProject()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve unfinished WorkTimeProject: %w", err)
	}

	if unfinishedWorkTimeProject != nil && unfinishedWorkTimeProject.ProjectID == projectID {
		log.Printf("WorkTimeProject already associated with project ID %d, returning existing project", projectID)
		return unfinishedWorkTimeProject, nil
	}

	workTime, err := getUnfinishedWorkTime()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve unfinished WorkTime: %w", err)
	}

	if workTime == nil {
		return nil, fmt.Errorf("no active WorkTime found to associate with the project")
	}

	if unfinishedWorkTimeProject != nil {
		_, err = getUnfinishedWorkTimeProjectAndFinish()
		if err != nil {
			return nil, fmt.Errorf("failed to finish previous WorkTimeProject: %w", err)
		}
	}

	workTimeProject := &WorkTimeProject{
		WorkTimeID: workTime.ID,
		ProjectID:  projectID,
		StartTime:  time.Now(),
		Closed:     false,
	}

	if err := DB.Create(workTimeProject).Error; err != nil {
		return nil, fmt.Errorf("failed to create association between Project and WorkTime: %w", err)
	}

	return workTimeProject, nil
}

func GetProject(id int64) (*Project, error) {
	var project Project

	err := DB.Preload("Tasks").Preload("Cost").First(&project, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("project not found")
	}

	err = updateProjectDuration(id)
	if err != nil {
		return nil, fmt.Errorf("failed to update project duration: %w", err)
	}

	err = DB.Preload("Tasks").Preload("Cost").First(&project, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("project not found")
	}

	return &project, nil
}

func ChangeProjectClose(id int64) (*Project, error) {
	project, error := GetProject(id)

	if error != nil {
		return nil, error
	}

	project.Closed = !project.Closed

	if err := DB.Save(&project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func ChangeProjectName(id int64, newName string) (*Project, error) {
	project, err := GetProject(id)
	if err != nil {
		return nil, err
	}

	project.Name = newName

	if err := DB.Save(&project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func SaveCost(projectID int64, hourCost int) (*Cost, error) {
	// Retrieve the project with the total duration
	project, err := GetProject(projectID)
	if err != nil {
		return nil, err
	}

	cost := Cost{
		ProjectID: projectID,
		Duration:  project.Duration,
		HourCost:  hourCost,
	}

	if err := DB.Where("project_id = ?", projectID).First(&cost).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		cost = Cost{
			ProjectID: projectID,
			Duration:  project.Duration,
			HourCost:  hourCost,
		}
		if err := DB.Create(&cost).Error; err != nil {
			return nil, err
		}
	} else {
		cost.Duration = project.Duration
		cost.HourCost = hourCost
		if err := DB.Save(&cost).Error; err != nil {
			return nil, err
		}
	}

	return &cost, nil
}

func CreateTask(projectID int64, description string, deadline time.Time) (*Task, error) {
	task := &Task{
		ProjectID:   projectID,
		Description: description,
		Deadline:    deadline,
	}
	err := DB.Create(task).Error

	var project Project
	if err := DB.Preload("Tasks").First(&project, projectID).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve TotalTime after creating WorkTime: %w", err)
	}

	project.Tasks = append(project.Tasks, *task)

	return task, err
}

func GetTask(id int64) (*Task, error) {
	var task Task
	err := DB.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Task not found")
	}
	return &task, err
}

func GetAllProjects(page int, pageSize int, closedFilter *bool, orderBy string, orderDirection string) ([]Project, int64, error) {
	var projects []Project
	var totalProjects int64

	query := DB.Model(&Project{})

	if closedFilter != nil {
		query = query.Where("closed = ?", *closedFilter)
	}

	err := query.Count(&totalProjects).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	orderField := "updated_at"
	if orderBy == "duration" {
		orderField = "duration"
	}

	if orderDirection != "asc" && orderDirection != "desc" {
		orderDirection = "desc"
	}

	orderClause := fmt.Sprintf("%s %s", orderField, strings.ToUpper(orderDirection))

	err = query.Order(orderClause).Offset(offset).Limit(pageSize).Find(&projects).Error
	if err != nil {
		return nil, 0, err
	}

	return projects, totalProjects, nil
}

func GetUnfinishedWorkTimeProject() (*WorkTimeProject, error) {
	var workTimeProject WorkTimeProject
	result := DB.Where("closed = ?", false).First(&workTimeProject)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("No unfinished WorkTimeProject found.")
			return nil, nil
		}
		return nil, result.Error
	}
	return &workTimeProject, nil
}

func GetWorkTimesForDay(day time.Time) ([]WorkTime, error) {
	var workTimes []WorkTime

	dayStart := day.UTC().Truncate(24 * time.Hour)
	dayEnd := dayStart.Add(24 * time.Hour)

	result := DB.Where("start_time >= ? AND start_time < ?", dayStart, dayEnd).Find(&workTimes)
	if result.Error != nil {
		return nil, result.Error
	}

	return workTimes, nil
}

func UpdateWorkTimeTrustworthy(workTime *WorkTime, trustworthy bool) error {
	workTime.Trustworthy = trustworthy
	return DB.Save(workTime).Error
}

func GetWorkTimesCrossingDays(day time.Time) ([]WorkTime, error) {
	var workTimes []WorkTime
	dayStart := day.Truncate(24 * time.Hour)

	result := DB.Where("start_time >= ? AND start_time < ? AND duration > ?", dayStart, dayStart.Add(24*time.Hour), 24*time.Hour).Find(&workTimes)
	if result.Error != nil {
		return nil, result.Error
	}

	return workTimes, nil
}

func GetBreakTimesForDay(day time.Time) ([]BreakTime, error) {
	var breakTimes []BreakTime
	dayStart := day.Truncate(24 * time.Hour)

	result := DB.Where("start_time >= ? AND start_time < ?", dayStart, dayStart.Add(24*time.Hour)).Find(&breakTimes)
	if result.Error != nil {
		return nil, result.Error
	}

	return breakTimes, nil
}

func GetBrbsForDay(day time.Time) ([]Brb, error) {
	var brbs []Brb
	dayStart := day.Truncate(24 * time.Hour)

	result := DB.Where("start_time >= ? AND start_time < ?", dayStart, dayStart.Add(24*time.Hour)).Find(&brbs)
	if result.Error != nil {
		return nil, result.Error
	}

	return brbs, nil
}

func GetWorkTimeProjectsByWorkTimeID(workTimeID int64) ([]WorkTimeProject, error) {
	var workTimeProjects []WorkTimeProject
	result := DB.Where("work_time_id = ?", workTimeID).Preload("Project").Find(&workTimeProjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return workTimeProjects, nil
}

func GetBreakTime(id int64) (*BreakTime, error) {
	var breakTime BreakTime
	err := DB.First(&breakTime, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("BreakTime not found")
	}
	return &breakTime, err
}

func GetWorkTimeProjectByID(id int64) (*WorkTimeProject, error) {
	var workTimeProject WorkTimeProject
	err := DB.First(&workTimeProject, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("WorkTimeProject not found")
	}
	return &workTimeProject, err
}
