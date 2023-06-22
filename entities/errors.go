package entities

type ErrorKind string

const (
	INVALID_PARAMETER ErrorKind = "invalid_parameter"
	UNEXPECTED        ErrorKind = "unexpected"
)

func (e ErrorKind) ToString() string {
	return string(e)
}
