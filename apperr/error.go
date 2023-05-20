package apperr

import "errors"

var (
	ErrAffectedRows = errors.New("The affected rows are not as expected")
)
