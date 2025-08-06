package main

import (
	"fmt"
	"time"
)

func add(input Input) int {
	time.Sleep(2 * time.Second)
	return input.A + input.B
}

type Input struct {
	A int
	B int
}

func timerDecorator(fn func(Input) int) func(Input) int {
	return func(input Input) int {

		timer := time.Now()
		result := fn(input)
		duration := time.Since(timer)
		fmt.Printf("O tempo de execução foi: %d\n", duration.Milliseconds())
		return result
	}
}

func main() {
	resultAddDecorator := timerDecorator(add)

	result0 := resultAddDecorator(Input{A: 20, B: 30})

	fmt.Printf("resultado: %d\n", result0)

}
