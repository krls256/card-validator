package errors

import (
	"errors"
	"github.com/krls256/card-validator/utils"
)

type errorWithCode struct {
	err  error
	code int
}

func NewErrorWithCode(err string, code int) errorWithCode {
	return errorWithCode{
		err:  errors.New(err),
		code: code,
	}
}

func WrapErrorWithCode(err error, code int) errorWithCode {
	return errorWithCode{
		err:  err,
		code: code,
	}
}

func (e errorWithCode) Error() string {
	return e.err.Error()
}

func (e errorWithCode) Code() int {
	return e.code
}

type manyUnwrap interface {
	Unwrap() []error
}

func ErrorToCodes(err error) (codes []int) {
	errorToCodes(err, &codes)

	return utils.Unique(codes)
}

func errorToCodes(err error, codes *[]int) {
	for err != nil {
		var ewc errorWithCode
		if ok := errors.As(err, &ewc); ok {
			*codes = append(*codes, ewc.Code())
		}

		var mu manyUnwrap
		if ok := errors.As(err, &mu); ok {
			for _, e := range mu.Unwrap() {
				errorToCodes(e, codes)
			}

			return
		}

		err = errors.Unwrap(err)
	}
}

type ErrorWithCode interface {
	error
	Code() int
}
