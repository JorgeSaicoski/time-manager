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
	WorkTimes  []*WorkTime
	BreakTime  *BreakTime
}

type WorkTime struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	TotalTimeID int64
	StartTime   time.Time
	Duration    time.Duration
	Projects    []*Project
	Brb         *Brb
}

type BreakTime struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	StartTime time.Time
	Duration  time.Duration
}

type Project struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	StartTime time.Time
	Duration  time.Duration
	Cost      *Cost
	WorkTimes []WorkTime `gorm:"foreignKey:ProjectID"`
	Tasks     []Task     `gorm:"foreignKey:ProjectID"`
}

type Task struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	ProjectID   int64
	Deadline    time.Time
	Description string
}

type Cost struct {
	gorm.Model
	ID       int64 `gorm:"primaryKey"`
	Time     time.Time
	HourCost int
}

type Brb struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	StartTime time.Time
	Duration  time.Duration
}
