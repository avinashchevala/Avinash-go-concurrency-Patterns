package pipeline

import "fmt"

//pipelines take data, do an operation on that data, and then pass it out
//operations on that data are done at stages
func Pipeline() {
	//input
	nums := []int{2, 3, 4, 5, 6, 7}

	//stage 1
	dataChannel := sliceToChannel(nums)

	//stage 2
	finalChannel := sq(dataChannel)

	//stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}

}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}
