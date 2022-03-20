package errors

type NotChangedError struct {
}

func (e NotChangedError) Error() string {
	return "Charta has not been changed"
}
