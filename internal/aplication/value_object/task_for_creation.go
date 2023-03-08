package value_object

type TaskForCreation struct {
	Description string
	Priority    int
}

func NewTaskForCreation(description string, priority int) TaskForCreation {
	return TaskForCreation{
		Description: description,
		Priority:    priority,
	}
}
