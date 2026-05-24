package users_transport_http

import (
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_reponse "github.com/Daty26/todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Daty26/todo-app/internal/core/transport/http/utils"
	"net/http"
)

func (h *UsersHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_reponse.NewHTTPesponseHandler(log, w)

	userID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}
	if err = h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}
	responseHandler.NoContentResponse()
}
