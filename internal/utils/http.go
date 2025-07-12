package utils

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ParseUintParam(r *http.Request, param string) (uint, error) {
	paramStr := chi.URLParam(r, param)
	if paramStr == "" {
		return 0, fmt.Errorf("missing parameter: %s", param)
	}

	var id uint
	if _, err := fmt.Sscanf(paramStr, "%d", &id); err != nil {
		return 0, fmt.Errorf("%s must be a number", param)
	}

	return id, nil
}
