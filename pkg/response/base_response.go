package response

type BaseResponse[T any] struct {
	Result *T  `json:"result"`
	Error  any `json:"error"`
	Status int `json:"status"`
}

func SuccessResponse[T any](result *T, status int) BaseResponse[T] {
	if status < 200 || status > 299 {
		status = 200
	}
	return BaseResponse[T]{
		Result: result,
		Error:  nil,
		Status: status,
	}
}

func ErrorResponse[T any](err error, status int) BaseResponse[T] {
	if status < 400 || status > 599 {
		status = 500
	}
	return BaseResponse[T]{
		Result: nil,
		Error:  err,
		Status: status,
	}
}
