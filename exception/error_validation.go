package exception

type ValidationError struct {
	Error string
}

func NewValidationError(error string) ValidationError {
	return ValidationError{Error: error}
}
