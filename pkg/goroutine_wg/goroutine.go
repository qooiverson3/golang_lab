package goroutinewg

import (
	"fmt"
	"sync"
)

// User is about user information
type User struct {
	ID int
}

// Says _
func (user User) Says(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("I'm user%d\n", user.ID)
}
