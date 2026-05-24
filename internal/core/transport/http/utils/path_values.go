package core_http_utils

import (
	"fmt"
	core_errors "github.com/Daty26/todo-app/internal/core/errors"
	"net/http"
	"strconv"
)

func GetIntPathValues(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key='%s' in path values: %w", key, core_errors.ErrInvalidArgument)
	}
	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("path value='%s' by key='%s' not a valid integer: %v: %w",
			pathValue,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	return val, nil
}
