package algorithm

import (
	"fmt"
)

func innerFilter(input <-chan int) <-chan int{
	prime := <- input

	fmt.Println(prime)
	output := make(chan int)

	go func(input <-chan int, output chan int){
		for {
			i :=<- input
			if i%prime != 0{
				output <- i
			}
		}
	}(input, output)

	return output
}

func firstChannel() <-chan int{
	output := make(chan int)
	go func(output chan int){
		for i:=2;i<100;i++{
			output <- i
		}
	}(output)

	return output
}

func PrintPrime(){
	input := firstChannel()

	for {
		output := innerFilter(input)

		input = output
	}
}
