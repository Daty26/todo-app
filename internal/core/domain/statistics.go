package domain

import "time"

type Statistics struct {
	TaskCreated                int
	TaskCompleted              int
	TaskCompletedRate          *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TaskCreated:                tasksCreated,
		TaskCompleted:              tasksCompleted,
		TaskCompletedRate:          tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
