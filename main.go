package main

import (
	"fmt"

	done "main.go/doneChannel"
	pipeline "main.go/pipelines"
	"main.go/primitives"
)

func main() {
	fmt.Println("main")
	primitives.Primitives()
	done.Done()
	pipeline.Pipeline()
}
