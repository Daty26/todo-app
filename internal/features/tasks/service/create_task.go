package tasks_service

import (
	"context"
	"fmt"
	"github.com/Daty26/todo-app/internal/core/domain"
)

func (s *TaskService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %w", err)
	}
	createdTask, err := s.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("repository create task: %w", err)
	}
	return createdTask, nil
}
