package tasks_transport_http

import (
	"fmt"
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_request "github.com/Daty26/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Daty26/todo-app/internal/core/transport/http/response"
	"net/http"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks godoc
// @Summary List tasks
// @Description View the list of tasks with optional user filter and pagination
// @Tags tasks
// @Produce json
// @Param user_id query int false "ID of the user that owns the tasks"
// @Param limit query int false "The size of the page with tasks"
// @Param offset query int false "Offset of the page with tasks"
// @Success 200 {array} TaskDTOResponse "Successfully received the list of tasks"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks [get]
func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPesponseHandler(log, w)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/limit/offset query params")
		return
	}
	tasks, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}
	response := GetTasksResponse(taskDTOsFromDomains(tasks))
	responseHandler.JSONResponse(response, http.StatusOK)

}
func getUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
		userIDQueryParamKey = "user_id"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `user_id` param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `limit` param: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `offset` param: %w", err)
	}
	return userID, limit, offset, nil
}
