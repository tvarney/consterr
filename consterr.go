package consterr

// Error is a constant error type.
type Error string

// Error returns the string version of the Error.
func (e Error) Error() string {
	return string(e)
}
