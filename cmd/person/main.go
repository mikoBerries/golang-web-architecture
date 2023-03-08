package main

import (
	"context"
	"fmt"

	s "github.com/golang-web-architecture/session"
)

func main() {

	ctx := context.Background()
	defer ctx.Done()
	ctx = s.SetUserId(ctx, "1")
	ctx = s.SetAdmin(ctx, true)

	fmt.Println(s.GetAdmin(ctx))
	fmt.Println(s.GetUserId(ctx))

}
