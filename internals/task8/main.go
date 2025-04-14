package task8

import (
	"fmt"
	"sync"
)

func RunTask8() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(numbers)) // buffered channel
	total := 0

	for range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num := <-ch
			mu.Lock()
			total += num
			mu.Unlock()
		}()
	}

	for _, n := range numbers {
		ch <- n
	}
	close(ch)

	wg.Wait()
	fmt.Println("Total penjumlahan:", total)
}
