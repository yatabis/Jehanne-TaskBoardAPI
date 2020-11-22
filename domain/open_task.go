package domain

import "time"

type TaskStatus string

const (
	TaskStatusUnowned   TaskStatus = "未所有"
	TaskStatusWaiting   TaskStatus = "待機中"
	TaskStatusCompleted TaskStatus = "完了"
)

type OpenTask struct {
	ID              uint       `json:"id"`
	Name            string     `json:"name"`
	Category        string     `json:"category"`
	Repeating       bool       `json:"repeating"`
	WorkOn          Date       `json:"work_on"`
	Deadline        Date       `json:"deadline"`
	Status          TaskStatus `json:"status"`
	PerformanceTime Minutes    `json:"performance_time"`
	CreatedAt       Datetime   `json:"created_at"`
	UpdatedAt       Datetime   `json:"updated_at"`
}

type OpenTaskOptions struct {
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	Repeating bool       `json:"repeating"`
	WorkOn    Date       `json:"work_on"`
	Deadline  Date       `json:"deadline"`
	Status    TaskStatus `json:"status"`
}

type Date time.Time
type Datetime time.Time
type Minutes time.Duration
