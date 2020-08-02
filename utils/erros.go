package utils

type RESTError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequest(message string) *RESTError {
	err := RESTError{
		Message: message,
		Status:  400,
		Error:   "bad request",
	}
	return &err
}

func NotFound(message string) *RESTError {
	err := RESTError{
		Message: message,
		Status:  404,
		Error:   "not fount",
	}
	return &err
}

func InternalErr(message string) *RESTError {
	err := RESTError{
		Message: message,
		Status:  500,
		Error:   "internal server error",
	}
	return &err
}
