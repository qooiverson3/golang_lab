package main

import (
	"fmt"
	coroutinewg "golang_lab/pkg/coroutine_wg"
	"sync"
)

func main() {
	users := 100
	user := coroutinewg.User{}

	wg := &sync.WaitGroup{}
	wg.Add(users)
	for i := 1; i <= users; i++ {
		user.ID = i
		go user.Says(wg)
	}
	wg.Wait()
	fmt.Println("Hello everyone~")
}
