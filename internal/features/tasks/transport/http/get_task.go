package tasks_transport_http

import (
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_request "github.com/Daty26/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Daty26/todo-app/internal/core/transport/http/response"
	"net/http"
)

type GetTaskResponse TaskDTOResponse

// GetTask godoc
// @Summary Get task
// @Description Get specific task by id
// @Tags tasks
// @Produce json
// @Param id path int true "ID of the task"
// @Success 200 {object} GetTaskResponse "Task was successfully found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPesponseHandler(logger, w)
	taskID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get taskID path value")
		return
	}
	task, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}
	reponse := GetTaskResponse(taskDTOFromDomain(task))
	responseHandler.JSONResponse(reponse, http.StatusOK)

}
