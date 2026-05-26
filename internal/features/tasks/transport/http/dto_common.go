package tasks_transport_http

import (
	"github.com/Daty26/todo-app/internal/core/domain"
	"time"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func taskDTOFromDomain(taskDomain domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           taskDomain.ID,
		Version:      taskDomain.Version,
		Title:        taskDomain.Title,
		Description:  taskDomain.Description,
		Completed:    taskDomain.Completed,
		CreatedAt:    taskDomain.CreatedAt,
		CompletedAt:  taskDomain.CompletedAt,
		AuthorUserID: taskDomain.AuthorUserID,
	}
}
func taskDTOsFromDomains(tasksDomain []domain.Task) []TaskDTOResponse {
	dtos := make([]TaskDTOResponse, len(tasksDomain))
	for i, task := range tasksDomain {
		dtos[i] = taskDTOFromDomain(task)
	}
	return dtos
}
