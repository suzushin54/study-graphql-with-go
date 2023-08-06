package graph

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/suzushin54/study-graphql-with-go/internal"
	"github.com/suzushin54/study-graphql-with-go/middlewares/auth"
)

var Directive = internal.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if _, ok := auth.ExtractUserName(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
