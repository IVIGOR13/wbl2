package main

import (
	"context"
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	ctx, cancel := context.WithCancel(context.Background())

	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			defer cancel()
			for {
				select {
				case val, opened := <-channel:
					if opened {
						out <- val
					} else {
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}(channel)
	}

	go func() {
		<-ctx.Done()
		close(out)
	}()

	return out
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(1*time.Second),
		sig(4*time.Second),
	)

	fmt.Printf("Done after %v\n", time.Since(start))

	// example 2

	sig2 := func(after time.Duration, n int) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for i := 1; i <= n; i++ {
				time.Sleep(after)
				fmt.Println(i, "/", n)
			}
		}()
		return c
	}

	start = time.Now()

	<-or(
		sig2(1*time.Second, 7),
		sig2(2*time.Second, 4),
		sig2(3*time.Second, 4),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
