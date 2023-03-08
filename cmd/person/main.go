package main

import (
	"context"
	"fmt"
)

type key string

var session key = "session"
var userId key = "userId"

func main() {
	ctx := context.Background()
	defer ctx.Done()
	ctx = context.WithValue(ctx, userId, 1)
	ctx = context.WithValue(ctx, session, "session")

	fmt.Println(ctx.Value(session))
	fmt.Println(ctx.Value(userId))
}
