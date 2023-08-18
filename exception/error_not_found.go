package exception

type NotFoundErrors struct {
	Error string
}

func NewNotFoundErrors(error string) NotFoundErrors {
	return NotFoundErrors{Error: error}
}
