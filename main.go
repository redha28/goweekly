package main

import (
	"fmt"
	"goweekly/internals/task1"
	"goweekly/internals/task2"
	"goweekly/internals/task3"
	"goweekly/internals/task4"
	"goweekly/internals/task5"
	"goweekly/internals/task6"
	"goweekly/internals/task7"
	"goweekly/internals/task8"
	"sync"
)

func printSeparator(title string) {
	fmt.Println()
	fmt.Println("==========================================")
	fmt.Printf("🔹 %s\n", title)
	fmt.Println("==========================================")
}

func main() {
	// task 1
	printSeparator("Task 1: Print Triangle")
	task1.PrintTriangle(5)

	// task 2
	printSeparator("Task 2: Find Matching Movie Durations")
	duration1, duration2, find := task2.FindMatchMovies(7)
	if find {
		fmt.Printf("🎬 Match movie durations: %d and %d \n", duration1, duration2)
	} else {
		fmt.Println("❌ No matching movie pair found")
	}

	// task 3
	printSeparator("Task 3: Round to Tenth")
	n := 4.66
	rounded := task3.RoundToTenth(n)
	fmt.Printf("🔢 Original: %.2f, Rounded: %.2f\n", n, rounded)

	// task 4
	printSeparator("Task 4: Deret Bilangan")
	deret := task4.DeretBilangan{Limit: 40}
	deret.Prima()
	deret.Ganjil()
	deret.Genap()
	deret.Fibonacci()

	// task 5
	printSeparator("Task 5: Implementasi struct dan method")
	var p task5.Hitung2D = &task5.PersegiPanjang{
		Panjang: 6,
		Lebar:   4,
	}
	fmt.Println("== Persegi Panjang ==")
	fmt.Printf("Luas: %.2f\n", p.Luas())
	fmt.Printf("Keliling: %.2f\n", p.Keliling())

	var b task5.Hitung = &task5.Kubus{
		Panjang: 5,
		Lebar:   3,
		Tinggi:  4,
	}
	fmt.Println("== Kubus ==")
	fmt.Printf("Luas: %.2f\n", b.Luas())
	fmt.Printf("Keliling: %.2f\n", b.Keliling())
	fmt.Printf("Volume: %.2f\n", b.Volume())

	// task 6
	printSeparator("Task 6: go routine channel")
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	ch := make(chan int)

	// Bagi slice menjadi dua bagian
	go task6.Sum(a[:len(a)/2], ch)
	go task6.Sum(a[len(a)/2:], ch)

	// Terima hasil dari kedua goroutine
	x, y := <-ch, <-ch

	fmt.Printf("Total: %d\n", x+y)

	// task 7
	printSeparator("Task 7: go routine channel")
	chslice := make(chan []int)
	var wg sync.WaitGroup
	wg.Add(2)
	go task7.Fibonacci(40, chslice, &wg)
	go task7.EvenNum([]int{10, 8, 3, 7, 24, 21}, chslice, &wg)
	even, fibo := <-chslice, <-chslice
	wg.Wait()
	fmt.Printf("fibo: %v\n", fibo)
	fmt.Printf("even nums: %v\n", even)

	// task 7
	printSeparator("Task 7")
	task8.RunTask8()
}
