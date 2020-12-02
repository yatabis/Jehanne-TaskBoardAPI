package domain

type TaskStatus string

const (
	TaskStatusUnowned   TaskStatus = "未所有"
	TaskStatusWaiting   TaskStatus = "待機中"
	TaskStatusCompleted TaskStatus = "完了"
)

func (s *TaskStatus) Init() {
	if *s == "" {
		*s = TaskStatusWaiting
	}
}

type TaskCategory string

const TaskCategoryInbox TaskCategory = "Inbox"

func (c *TaskCategory) Init() {
	if *c == "" {
		*c = TaskCategoryInbox
	}
}

type OpenTask struct {
	ID              uint         `json:"id"`
	Name            string       `json:"name"`
	Category        TaskCategory `json:"category"`
	Repeating       bool         `json:"repeating"`
	WorkOn          Date         `json:"work_on"`
	Deadline        Date         `json:"deadline"`
	Status          TaskStatus   `json:"status"`
	PerformanceTime Minutes      `json:"performance_time"`
	CreatedAt       Datetime     `json:"created_at"`
	UpdatedAt       Datetime     `json:"updated_at"`
}

type OpenTaskOptions struct {
	Name      string       `json:"name"`
	Category  TaskCategory `json:"category"`
	Repeating bool         `json:"repeating"`
	WorkOn    Date         `json:"work_on"`
	Deadline  Date         `json:"deadline"`
	Status    TaskStatus   `json:"status"`
}
