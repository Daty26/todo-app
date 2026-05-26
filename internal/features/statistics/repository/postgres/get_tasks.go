package statistics_postgres_repository

import (
	"context"
	"fmt"
	"github.com/Daty26/todo-app/internal/core/domain"
	"strings"
	"time"
)

func (r *StatisticsRepository) GetTasks(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder
	queryBuilder.WriteString(`
	SELECT id, version, title, description, completed, created_at, completed_at, user_id
	FROM todoapp.tasks
`)
	args := []any{}
	conditions := []string{}
	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("user_id=$%d", len(args)+1))
		args = append(args, userID)
	}
	if from != nil {
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)+1))
		args = append(args, from)
	}
	if to != nil {
		conditions = append(conditions, fmt.Sprintf("completed_at<$%d", len(args)+1))
		args = append(args, to)
	}
	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}
	fmt.Println(queryBuilder.String())
	rows, err := r.pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()
	var taskModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		err = rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorUserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan tasks: %w", err)
		}
		taskModels = append(taskModels, taskModel)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}
	taskDomains := taskDomainFromModels(taskModels)
	return taskDomains, nil

}
