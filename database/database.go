package database

import (
	"errors"
	"log"
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

func CreateTotalTime() (*TotalTime, error) {
	totalTime := &TotalTime{
		StartTime: time.Now(),
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
	result := DB.Where("finish_time < start_time").First(&totalTime)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &totalTime, nil
}

func CreateWorkTime(totalTimeID int64) (*WorkTime, error) {
	workTime := &WorkTime{
		StartTime:   time.Now(),
		TotalTimeID: totalTimeID,
	}
	err := DB.Create(workTime).Error
	return workTime, err
}

func GetWorkTime(id int64) (*WorkTime, error) {
	var workTime WorkTime
	err := DB.Preload("Projects").First(&workTime, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("WorkTime not found")
	}
	return &workTime, err
}

func CreateProject(workTimeID int64) (*Project, error) {
	project := &Project{
		StartTime: time.Now(),
	}
	err := DB.Create(project).Error
	return project, err
}

func GetProject(id int64) (*Project, error) {
	var project Project
	err := DB.Preload("Tasks").Preload("Cost").First(&project, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Project not found")
	}
	return &project, err
}

func CreateTask(projectID int64, description string, deadline time.Time) (*Task, error) {
	task := &Task{
		ProjectID:   projectID,
		Description: description,
		Deadline:    deadline,
	}
	err := DB.Create(task).Error
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
