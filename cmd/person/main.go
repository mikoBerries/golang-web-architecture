package main

import (
	"context"
	"fmt"
	"runtime"
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
	/*
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
	*/

	ctx, cancelF := context.WithCancel(context.Background())
	defer cancelF()
	// defer time.Sleep(2 * time.Second)
	for i := 0; i < 100; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					runtime.Goexit()
					// return
				default:
					fmt.Println("go routine now :", runtime.NumGoroutine())
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()

	}
	time.Sleep(2 * time.Second)
	cancelF()
	fmt.Println("cancelF called ")
	time.Sleep(1 * time.Second)
	defer fmt.Println("last go routine:", runtime.NumGoroutine())
}
