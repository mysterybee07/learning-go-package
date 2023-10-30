/*context:
The context package is used for managing and carrying request-scoped values across API boundaries and between processes.
It is often used for cancellation, deadlines, and carrying metadata during the execution of a program.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go doSomething(ctx)

	//Simulate a cancellation signal

	time.Sleep(5 * time.Second)
	cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Task cancelled.")
	}
}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task received cancellation signal.")
			return
		default:
			fmt.Println("Task is still running....")
			time.Sleep(1 * time.Second)
		}
	}
}
