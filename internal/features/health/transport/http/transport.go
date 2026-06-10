package health_transport_http

import (
	"context"
	"net/http"

	core_postgres_pool "github.com/Daty26/todo-app/internal/core/repository/postgres/pool"
	core_http_server "github.com/Daty26/todo-app/internal/core/transport/http/server"
)

type HealthHTTPHandler struct {
	pool core_postgres_pool.Pool
}

func NewHealthHTTPHandler(pool core_postgres_pool.Pool) *HealthHTTPHandler {
	return &HealthHTTPHandler{
		pool: pool,
	}
}

func (h *HealthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/healthz",
			Handler: h.Healthz,
		},
		{
			Method:  http.MethodGet,
			Path:    "/readyz",
			Handler: h.Readyz,
		},
	}
}

func (h *HealthHTTPHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func (h *HealthHTTPHandler) Readyz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), h.pool.OpTimeout())
	defer cancel()

	var result int
	if err := h.pool.QueryRow(ctx, "SELECT 1").Scan(&result); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(`{"status":"unavailable"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ready"}`))
}
