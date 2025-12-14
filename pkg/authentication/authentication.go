package authentication

import (
	"context"
	"errors"
)

type userIdKeyType struct{}

var userIdKey userIdKeyType

func InjectHttpUserId(userId int64, ctx context.Context) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}

func ExtractHttpUserId(ctx context.Context) (int64, error) {
	value := ctx.Value(userIdKey)
	if value == nil {
		return 0, errors.New("no user id is found in a context")
	}

	userId, ok := value.(int64)
	if !ok {
		return 0, errors.New("the value extracted from a context cannot be casted to int")
	}

	return userId, nil
}
