package dto

type TaskDto struct {
	Id          int
	Description string
	Priority    int
}

func NewDto(id int, description string, priority int) TaskDto {
	return TaskDto{
		Id:          id,
		Description: description,
		Priority:    priority,
	}

}

func (d TaskDto) GetId() int {
	return d.Id
}

func (d TaskDto) GetDescription() string {
	return d.Description
}

func (d TaskDto) GetPriority() int {
	return d.Priority
}
