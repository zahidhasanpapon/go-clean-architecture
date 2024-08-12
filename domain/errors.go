package domain

import "errors"

var (
	ErrorInternalServerError = errors.New("internal Server Error")
	ErrorNotFound            = errors.New("your requested Item is not found")
	ErrorConflict            = errors.New("your Item already exist")
	ErrorBadParamInput       = errors.New("given Param is not valid")
)
