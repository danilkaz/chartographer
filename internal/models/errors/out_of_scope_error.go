package errors

type OutOfScopeError struct {
}

func (e OutOfScopeError) Error() string {
	return "Request out of bounds of charta"
}
