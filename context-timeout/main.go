package main

import (
	"context"
	"fmt"
	"time"
)

func generate(ctx context.Context) <-chan int {

	ch := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-ctx.Done():
				close(ch)
				return
			}
		}
	}()

	return ch

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := generate(ctx)

	for i := range c {
		fmt.Printf("count %d \n", i)
	}

}
