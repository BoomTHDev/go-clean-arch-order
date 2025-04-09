package controller

type (
	Response struct {
		Success bool   `json:"success"`
		Data    any    `json:"data,omitempty"`
		Error   string `json:"error,omitempty"`
	}
)

func SuccessResponse(data any) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func ErrorResponse(errorMsg string) Response {
	return Response{
		Success: false,
		Error:   errorMsg,
	}
}
