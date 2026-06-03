package tasks_transport_http

import (
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_request "github.com/Daty26/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Daty26/todo-app/internal/core/transport/http/response"
	"net/http"
)

// DeleteTask godoc
// @Summary Delete task
// @Description Delete an existing task by id
// @Tags tasks
// @Param id path int true "ID of the task that needs to be deleted"
// @Success 204 "Successfully deleted task"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPesponseHandler(log, w)
	id, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get taskID value")
	}

	if err = h.tasksService.DeleteTask(ctx, id); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete task")
		return
	}
	responseHandler.NoContentResponse()
}
