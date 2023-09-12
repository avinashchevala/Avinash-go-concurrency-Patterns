package forSelect

import "fmt"

func ForSelect() {

	charChannel := make(chan string, 3)
	//3 is the limit,  this makes it a buffered channel. this is to make asynchronus communication
	// with buffered chnnels we can use queue like functionality
	//but only till capacity is met. one filled buffered channel has same functionality as unbuffered
	chars := []string{"avi", "arrow", "bobby"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)
	for result := range charChannel {
		fmt.Println(result)
	}

}
