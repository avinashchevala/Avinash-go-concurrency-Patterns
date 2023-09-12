package primitives

import (
	"fmt"
)

func someFunc(num string) {
	fmt.Println(num)
}

// go  routines, channels, select statement, for select loop
func Primitives() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go func() {
		chan1 <- "data1"
	}()
	go func() {
		chan2 <- "data2"
	}()

	//msg := <-chan1
	select {
	case msgFromChan1 := <-chan1:
		fmt.Println(msgFromChan1)
	case msgFrimChan2 := <-chan2:
		fmt.Println(msgFrimChan2)
	}
}
