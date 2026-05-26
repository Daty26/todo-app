package statistics_transport_http

import (
	"fmt"
	"github.com/Daty26/todo-app/internal/core/domain"
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_request "github.com/Daty26/todo-app/internal/core/transport/http/request"
	core_http_reponse "github.com/Daty26/todo-app/internal/core/transport/http/response"
	"net/http"
	"time"
)

type GetStatisticsResponse struct {
	TaskCreated                int      `json:"task_created"`
	TaskCompleted              int      `json:"task_completed"`
	TaskCompletedRate          *float64 `json:"task_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_reponse.NewHTTPesponseHandler(log, w)
	userID, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/from/to query params")
		return
	}
	statistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")
		return
	}
	response := toDTOFromDomain(statistics)
	responseHandler.JSONResponse(response, http.StatusOK)

}
func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		avgTime = &duration
	}
	return GetStatisticsResponse{
		TaskCreated:                statistics.TaskCreated,
		TaskCompleted:              statistics.TaskCompleted,
		TaskCompletedRate:          statistics.TaskCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)
	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `user_id` query param: %w", err)
	}
	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `from` query param: %w", err)
	}
	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `to` query param: %w", err)
	}
	return userID, from, to, nil
}
