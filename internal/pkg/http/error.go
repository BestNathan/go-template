package http

import "strings"

type ErrorCode int

const (
	ErrCodeCommon ErrorCode = iota + 40000000
	ErrCodeInvalidParameters
	ErrCodeResourceNotExist
	ErrCodeInvalidRequestMethod
	ErrCodeMissingParameters
	ErrCodeInvalidParametersType
	ErrCodeSessionTimeout   ErrorCode = 40000401
	ErrCodeNoPermission     ErrorCode = 40000403
	ErrCodeResourceConflict ErrorCode = 40000409
)

type ErrorWithDetail interface {
	error

	Detail() string
}

type Error struct {
	Code       ErrorCode
	Err        error
	DetailInfo string
}

func NewCommonError(err error, detail ...string) *Error {
	return &Error{Code: ErrCodeCommon, Err: err, DetailInfo: strings.Join(detail, "\n")}
}

func NewResourceNotExistError(err error, detail ...string) *Error {
	return &Error{Code: ErrCodeResourceNotExist, Err: err, DetailInfo: strings.Join(detail, "\n")}
}

// errors interface
func (e Error) Error() string {
	return e.Err.Error()
}

// ErrorWithDetail interface
func (e Error) Detail() string {
	if e.DetailInfo == "" {
		return e.Error()
	}

	return e.DetailInfo
}
