package graph

import (
	"artkelo/middlewares"
	"artkelo/utils"
	"context"
	"slices"

	graphqllib "github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func LogOperations(
	ctx context.Context,
	next graphqllib.OperationHandler,
) graphqllib.ResponseHandler {
		req := graphqllib.GetOperationContext(ctx)
		logger := middlewares.GetAPIContext(ctx).Logger
		logger.Debug(
			"Graphql request",
			zap.String("RequestId", middlewares.GetRequestIdFromContext(ctx)),
			zap.String("OperationName", req.OperationName),
			zap.Any("fields", req.Variables),
		)
		return next(ctx)
	}

func serializeError(error *gqlerror.Error) map[string]string {
	buffer := make(map[string]string)
	buffer["message"] = error.Message
	buffer["path"] = error.Path.String()
	return buffer
}

func LogResponses(
	ctx context.Context,
	next graphqllib.ResponseHandler,
) *graphqllib.Response {
	res := next(ctx)
	logger := middlewares.GetAPIContext(ctx).Logger
	fields := []zap.Field {
		zap.String("RequestId", middlewares.GetRequestIdFromContext(ctx)),
	}
	if len(res.Errors) != 0 {
		errors := slices.Collect(utils.Map(slices.Values(res.Errors), serializeError))
		fields = append(fields, zap.Any("errors", errors))
	} else {
		fields = append(fields, zap.Any("body", res.Data))
	}
	logger.Debug(
		"Graphql response",
		fields...,
	)
	return res
}
