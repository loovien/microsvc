package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*100))

	c := time.After(time.Second * 4)

	fmt.Printf("what's wrong!??\n")
	select {
	case a := <-ctx.Done():
		fmt.Printf("xxxxxxxxxxxxxxx\n")
		fmt.Println(a)
	case <-c:
		cancel()
		fmt.Println(<-ctx.Done())
		fmt.Println(<-ctx.Done())
		fmt.Println(<-ctx.Done())
		fmt.Printf("endof timeout:!\n")
	}

}
