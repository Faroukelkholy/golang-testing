package errors

type ApiError struct {
	Status  int
	Message string
	Error   string
}
