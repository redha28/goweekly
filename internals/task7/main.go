package task7

import "sync"

func Fibonacci(d int, ch chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var result []int
	a, b := 0, 1
	for a <= d {
		result = append(result, a)
		a, b = b, a+b
	}
	ch <- result
}

func EvenNum(nums []int, ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var result []int
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	ch <- result
}
