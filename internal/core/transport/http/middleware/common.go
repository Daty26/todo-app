package core_http_middleware

import (
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_reponse "github.com/Daty26/todo-app/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var reqIDHeader = "X-Request-ID"

func CORS(allowerdOriginsList []string) Middleware {
	allowedOrigins := make(map[string]struct{})
	for _, origin := range allowerdOriginsList{
		allowedOrigins[origin] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(reqIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(reqIDHeader, requestID)
			w.Header().Set(reqIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Header.Get(reqIDHeader)
			l := log.With(
				zap.String("request_id", reqID),
				zap.String("url", r.URL.String()),
			)
			ctx := core_logger.ToContext(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_reponse.NewResponseWriter(w)
			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.String("HTTP_method", r.Method),
				zap.Time("time", before.UTC()),
			)
			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Now().Sub(before)),
			)
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_reponse.NewHTTPesponseHandler(log, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
