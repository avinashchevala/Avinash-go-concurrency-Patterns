package done

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing Work")
			time.Sleep(time.Second * 1)
		}
	}
}

// used to avoid goroutine leak
func Done() {
	done := make(chan bool)
	go doWork(done)
	time.Sleep(time.Second * 3)
	close(done)
}
