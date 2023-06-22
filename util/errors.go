package util

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/daichi1002/go-graphql/entities"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewGqlError(ctx context.Context, err error, errorKind entities.ErrorKind) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": errorKind.ToString(),
		},
	}
}
