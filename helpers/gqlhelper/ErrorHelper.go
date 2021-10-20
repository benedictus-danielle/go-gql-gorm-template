package gqlhelper

import "github.com/vektah/gqlparser/v2/gqlerror"

func CreateError(message string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
	}
}
