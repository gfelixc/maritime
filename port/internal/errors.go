package internal

import "fmt"

type DomainErrorCode string

const (
	InvalidUNLOC       DomainErrorCode = "InvalidUNLOC"
	InvalidCountry     DomainErrorCode = "InvalidCountry"
	InvalidName        DomainErrorCode = "InvalidName"
	InvalidCoordinates DomainErrorCode = "InvalidCoordinates"
	InvalidTimezone    DomainErrorCode = "InvalidTimezone"
	InvalidProvince    DomainErrorCode = "InvalidProvince"
	InvalidCity        DomainErrorCode = "InvalidCity"
	CodeNotProvided    DomainErrorCode = "CodeNotProvided"
)

type DomainError struct {
	code    DomainErrorCode
	message string
	wrapped error
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s] - %s", e.code, e.message)
}

func (e *DomainError) Unwrap() error {
	return e.wrapped
}

func NewDomainError(code DomainErrorCode, message string) error {
	return &DomainError{
		code:    code,
		message: message,
	}
}

func WrapWithDomainError(code DomainErrorCode, message string, wrapped error) error {
	return &DomainError{
		code:    code,
		message: message,
		wrapped: wrapped,
	}
}
