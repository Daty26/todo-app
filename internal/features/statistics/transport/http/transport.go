package statistics_transport_http

import (
	"context"
	"github.com/Daty26/todo-app/internal/core/domain"
	core_http_server "github.com/Daty26/todo-app/internal/core/transport/http/server"
	"net/http"
	"time"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}
type StatisticsService interface {
	GetStatistics(ctx context.Context, userID *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(service StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{statisticsService: service}
}
func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}

}
