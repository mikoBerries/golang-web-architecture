package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//hand on exercise 2
	/*
		ctx =context.Background()
		defer cancelF()
		ctx.Done()
		// defer ctx.Done()
		ctx = s.SetUserId(ctx, "1")
		ctx = s.SetAdmin(ctx, true)

		fmt.Println(s.GetAdmin(ctx))
		fmt.Println(s.GetUserId(ctx))
	*/
	//hand on exercise 4

	timeout := 2 * time.Second
	ctx, cancelF := context.WithTimeout(context.Background(), timeout)
	defer cancelF()
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("exit")
	default:
		fmt.Println("finish")
	}

}
