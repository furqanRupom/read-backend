package graphqlresolvers

import (
	  "read-backend/graphql"
	 "read-backend/middlewares"
	"context"
	"errors"
	"net/http"

	graphqllib "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/ast"
)


func CreateHandler() http.Handler {
	c := graphql.Config{Resolvers: &Resolver{}}
	c.Directives.Authenticated = func (ctx context.Context, obj interface {}, next graphqllib.Resolver) (interface{}, error) {
		if middlewares.GetUserFromContext(ctx) == nil {
			return nil, errors.New("Unauthorized")
		}
		return next(ctx)
	}
	c.Directives.Admin = func (ctx context.Context, obj interface {}, next graphqllib.Resolver) (interface{}, error) {
		user := middlewares.GetUserFromContext(ctx)
		if !user.IsAdmin {
			return nil, errors.New("Forbidden")
		}
		return next(ctx)
	}
	srv := handler.New(graphql.NewExecutableSchema(c))
	srv.AroundOperations(graphql.LogOperations)
	srv.AroundResponses(graphql.LogResponses)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}
