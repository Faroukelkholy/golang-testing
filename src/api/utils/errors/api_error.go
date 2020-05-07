package errors

type ApiError struct {
	StatusCode  int
	Message string
	Error   string
}
