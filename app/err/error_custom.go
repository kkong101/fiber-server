package err

import "errors"

// Error constants
var (
	ErrInvalidKey      = errors.New("key is invalid")
	ErrInvalidKeyType  = errors.New("key is of invalid type")
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")
)

const (
	ValidationErrorMalformed = iota
	ValidationErrorSignatureInvalid
	ValidationErrorAudience
	ValidationErrorExpired
	ValidationErrorIssuedAt
	ValidationErrorIssuer
	ValidationErrorNotValidYet
	ValidationErrorId
	ValidationErrorClaimsInvalid
)

type CustomError struct {
	Inner  error
	Errors uint8
	text   string
}

func NewCustomError(errorText string, errorFlags uint8) *CustomError {
	return &CustomError{
		Errors: errorFlags,
		text:   errorText,
	}
}

func (e CustomError) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.text != "" {
		return e.text
	} else {
		return "ERROR"
	}
}
