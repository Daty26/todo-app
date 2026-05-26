package statistics_service

import (
	"context"
	"github.com/Daty26/todo-app/internal/core/domain"
	"time"
)

type StatisticsService struct {
	statisticsRepository StatisticsRespository
}
type StatisticsRespository interface {
	GetTasks(ctx context.Context, userID *int, from *time.Time, to *time.Time) ([]domain.Task, error)
}

func NewStatisticService(statisticsRepository StatisticsRespository) *StatisticsService {
	return &StatisticsService{statisticsRepository: statisticsRepository}
}
