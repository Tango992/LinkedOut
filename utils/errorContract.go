package utils

import "net/http"

type ErrResponse struct {
	Status int   `json:"status" extensions:"x-order=0"`
	Type   string `json:"type" extensions:"x-order=1"`
	Detail string `json:"detail" extensions:"x-order=2"`
}

var (
	ErrBadRequest = ErrResponse{
		Status: http.StatusBadRequest,
		Type:   "Bad Request",
	}

	ErrUnauthorized = ErrResponse{
		Status: http.StatusUnauthorized,
		Type:   "Unauthorized",
	}

	ErrNotFound = ErrResponse{
		Status: http.StatusNotFound,
		Type:   "Not Found",
	}

	ErrConflict = ErrResponse{
		Status: http.StatusConflict,
		Type:   "Conflict",
	}

	ErrInternalServer = ErrResponse{
		Status: http.StatusInternalServerError,
		Type:   "Internal server error",
	}
)

func (er ErrResponse) Details(detail string) (int, ErrResponse) {
	er.Detail = detail
	return er.Status, er
}
