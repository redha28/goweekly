package task8

import (
	"fmt"
	"sync"
)

func SumNum(num1 int, num2 int, result chan int, wgTask8 *sync.WaitGroup) {
	defer wgTask8.Done()
	sum := num1 + num2
	result <- sum
}

func PrintOddEven(num int) string {
	if num%2 == 0 {
		return fmt.Sprint("Genap")
	}
	return fmt.Sprint("Ganjil")
}
