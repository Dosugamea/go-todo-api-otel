package response

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	switch v := err.(type) {
	case *echo.HTTPError:
		return ErrorResponse{
			Message: v.Message.(string),
		}
	default:
		return ErrorResponse{
			Message: v.Error(),
		}
	}
}
