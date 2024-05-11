package middlewares

import (
	"net/http"

	"google.golang.org/grpc/status"
)

func ErrorHandler(err error) (int, any) {
	if st, ok := status.FromError(err); ok {
		return http.StatusOK, Body{
			Code: int(st.Code()),
			Message:  st.Message(),
		}
	}

	code := http.StatusInternalServerError
	return http.StatusOK, Body{
		Code: code,
		Message:  err.Error(),
	}
}
