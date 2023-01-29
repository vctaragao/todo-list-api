package entity

type Task struct {
	Id          int
	Description string
	Priority    int
}

func NewTask(description string, priority int) Task {
	return Task{
		Description: description,
		Priority:    priority,
	}
}
