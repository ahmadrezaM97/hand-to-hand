package main

import (
	"fmt"
	"sync"

	"github.com/ahmadrezam97/hand-to-hand/internal/channels/bank"
)

func getDiff() bool {
	acc := bank.Account{Value: 1997}

	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < 1000000; i++ {
			acc.Increment()
		}
		defer wg.Done()
	}(&wg)
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < 1000000; i++ {
			acc.Decrement()
		}
		defer wg.Done()
	}(&wg)

	wg.Wait()
	return acc.Value == 1997
}

func check() int {
	cnt := 0
	for i := 0; i < 10; i++ {
		if getDiff() == false {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(check())
}
