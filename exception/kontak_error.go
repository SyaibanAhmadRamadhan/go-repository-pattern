package exception

type BadRequestError struct {
	Error any
}

func NewBadRequestError(error any) BadRequestError {
	return BadRequestError{Error: error}
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
