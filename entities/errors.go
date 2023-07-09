package entities

type ErrorKind string

const (
	INVALID_PARAMETER ErrorKind = "invalid_parameter"
	UNEXPECTED        ErrorKind = "unexpected"
	NOT_FOUND         ErrorKind = "not_found"
)

func (e ErrorKind) ToString() string {
	return string(e)
}

func (e ErrorKind) Error() string {
	return e.ToString()
}
