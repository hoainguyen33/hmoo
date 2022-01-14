package errors

import (
	"fmt"
)

var (
	ErrQueryInvalid = fmt.Errorf("QUERY_INVALID")
	ErrBodyInvalid  = fmt.Errorf("BODY_INVALID")
	ErrNotFound     = fmt.Errorf("NOT_FOUND")
)
