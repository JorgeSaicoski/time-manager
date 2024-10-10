package database

import (
	"errors"
	"fmt"
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

	var unfinishedWorkTime WorkTime
	err = DB.Where("total_time_id = ? AND duration = 0", id).First(&unfinishedWorkTime).Error
	if err == nil {
		unfinishedWorkTime.Duration = totalTime.FinishTime.Sub(unfinishedWorkTime.StartTime)

		if err := DB.Save(&unfinishedWorkTime).Error; err != nil {
			return nil, fmt.Errorf("failed to finish WorkTime: %w", err)
		}
		log.Printf("Finished WorkTime with ID %d before finishing TotalTime %d", unfinishedWorkTime.ID, totalTime.ID)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking unfinished WorkTime: %w", err)
	}

	if err := DB.Save(&totalTime).Error; err != nil {
		return nil, fmt.Errorf("failed to finish TotalTime: %w", err)
	}

	return &totalTime, nil
}

func CreateWorkTime(totalTimeID int64) (*WorkTime, error) {
	workTime := &WorkTime{
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
	log.Printf("get work time")
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
