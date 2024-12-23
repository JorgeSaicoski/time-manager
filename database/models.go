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
	Brb        *Brb       `gorm:"foreignKey:TotalTimeID;constraint:OnDelete:CASCADE"`
	Closed     bool
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
	Duration  time.Duration
	HourCost  int
}

type BreakTime struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	TotalTimeID int64
	StartTime   time.Time
	Duration    time.Duration
	Active      bool `gorm:"default:true"`
}

type Brb struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	TotalTimeID int64
	StartTime   time.Time
	Duration    time.Duration
	Active      bool `gorm:"default:true"`
}

type ResolutionTracker struct {
	gorm.Model
	ID       int64 `gorm:"primaryKey"`
	Day      time.Time
	Category string `gorm:"size:255"`
	Closed   bool
	Units    []ResolutionUnit `gorm:"foreignKey:TrackerID"`
}

type ResolutionUnit struct {
	gorm.Model
	ID         int64 `gorm:"primaryKey"`
	TrackerID  int64
	Tracker    ResolutionTracker `gorm:"foreignKey:TrackerID"`
	Identifier string            `gorm:"size:255"`
	Resolved   bool
}
