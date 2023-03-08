//session builder
package session

import (
	"context"
)

type stringKey string
type intKey int

var userID stringKey
var admin intKey

func SetUserId(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, userID, val)
}

func SetAdmin(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, btoi(isAdmin))
}

func GetUserId(ctx context.Context) int {
	if val := ctx.Value(userID); val != nil {
		if i, ok := val.(int); ok {
			return i
		}
	}
	return 0
}
func GetAdmin(ctx context.Context) bool {
	if val := ctx.Value(admin); val != nil {
		if i, ok := val.(bool); ok {
			return i
		}
	}
	return false
}
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
