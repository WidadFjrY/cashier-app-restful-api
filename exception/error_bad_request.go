package exception

type BadRequestErrors struct {
	Error string
}

func NewBadRequestErrors(error string) BadRequestErrors {
	return BadRequestErrors{Error: error}
}
