package users_transport_http

import (
	"fmt"
	"github.com/Daty26/todo-app/internal/core/domain"
	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_request "github.com/Daty26/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Daty26/todo-app/internal/core/transport/http/response"
	core_http_types "github.com/Daty26/todo-app/internal/core/transport/http/types"
	"net/http"
	"strings"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" swaggertype:"string" example:"Max Verstappen"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string" example:"+770765343112"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("`FullName` can't be null")
		}
		fulllNameLen := len([]rune(*r.FullName.Value))
		if fulllNameLen < 3 || fulllNameLen > 100 {
			return fmt.Errorf("`FullName` must be between 3 and 100 symbols")
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("`PhoneNumber` must be between 10 and 15 symbols")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must start with '+' symbol")
			}
		}
	}
	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser godoc
// @Summary Patch user
// @Description Partially update an existing user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID of the user that needs to be patched"
// @Param request body PatchUserRequest true "PatchUser request body"
// @Success 200 {object} PatchUserResponse "Successfully patched user"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPesponseHandler(log, w)

	userID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}
	var request PatchUserRequest
	if err = core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode  and validate HTTP request")
		return
	}

	userPatch := userPatchFromRequest(request)
	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)

	log.Debug(
		fmt.Sprintf("PatchUserRequest fields:\n FullName: '%v'\nPhoneNumber: '%v' ",
			request.FullName,
			request.PhoneNumber,
		))
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(request.FullName.ToDomain(), request.PhoneNumber.ToDomain())
}
