package database

import (
	"time"

	"gorm.io/gorm"
)

type TotalTime struct {
	gorm.Model
	ID         int64 `gorm:"primaryKey"`
	StartTime  time.Time
	FinishTime time.Time
	WorkTimes  []WorkTime `gorm:"foreignKey:TotalTimeID"`
	BreakTime  *BreakTime `gorm:"foreignKey:TotalTimeID;constraint:OnDelete:CASCADE"`
}

type WorkTime struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	TotalTimeID int64
	StartTime   time.Time
	Duration    time.Duration
	Projects    []Project `gorm:"many2many:work_time_projects;"`
}

type Project struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	StartTime time.Time
	Duration  time.Duration
	Closed    bool
	Cost      *Cost      `gorm:"foreignKey:ProjectID"`
	WorkTimes []WorkTime `gorm:"many2many:work_time_projects;"`
}

type BreakTime struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	TotalTimeID int64
	StartTime   time.Time
	Duration    time.Duration
}

type Task struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	ProjectID   int64
	Deadline    time.Time
	Description string
	Closed      bool
}

type Cost struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	ProjectID int64 `gorm:"uniqueIndex"`
	Time      time.Time
	HourCost  int
}

type Brb struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	StartTime time.Time
	Duration  time.Duration
}
