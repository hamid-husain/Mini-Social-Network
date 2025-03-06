package errors

type APIError struct {
	Message    string
	StatusCode int
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *APIError) Status() int {
	return e.StatusCode
}

func NewAPIError(message string, statusCode int) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
	}
}
