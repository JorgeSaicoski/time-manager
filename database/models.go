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
	Closed      bool
	Trustworthy bool `gorm:"default:true"`
}

type Project struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	Name      string
	StartTime time.Time
	Duration  time.Duration
	Closed    bool
	Cost      *Cost      `gorm:"foreignKey:ProjectID"`
	WorkTimes []WorkTime `gorm:"many2many:work_time_projects;"`
	Tasks     []Task     `gorm:"foreignKey:ProjectID"`
}

type WorkTimeProject struct {
	ID         int64     `gorm:"primaryKey"`
	WorkTimeID int64     `gorm:"primaryKey"`
	ProjectID  int64     `gorm:"primaryKey"`
	StartTime  time.Time `gorm:"not null"`
	Duration   time.Duration
	Closed     bool
	WorkTime   WorkTime `gorm:"foreignKey:WorkTimeID"`
	Project    Project  `gorm:"foreignKey:ProjectID"`
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
