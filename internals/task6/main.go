package task6

func Sum(d []int, ch chan int) {
	// defer close(ch)
	total := 0
	for _, v := range d {
		total += v
	}
	ch <- total
}
