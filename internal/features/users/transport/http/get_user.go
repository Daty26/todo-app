package users_transport_http

import (
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_request "github.com/Daty26/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Daty26/todo-app/internal/core/transport/http/response"
	"net/http"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary Get user
// @Description Get specific user by id
// @Tags users
// @Produce json
// @Param id path int true   "ID of the user"
// @Success 200 {object} GetUserResponse "User was successfully found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPesponseHandler(logger, w)
	userID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId path value")
		return
	}
	getUser, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}
	userDomain := GetUserResponse(userDTOFromDomain(getUser))
	responseHandler.JSONResponse(userDomain, http.StatusOK)
}
