package create_task

type TaskDto struct {
	Description string
	Priority    int
}

func NewTaskDto(description string, priority int) TaskDto {
	return TaskDto{
		Description: description,
		Priority:    priority,
	}
}
