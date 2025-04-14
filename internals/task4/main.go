package task4

import (
	"fmt"
)

// Struct untuk deret bilangan
type DeretBilangan struct {
	Limit int
}

// Fungsi bantu untuk cek bilangan prima
func isPrime(n int) bool {
	// fmt.Println(n)
	// if n < 2 {
	// return false
	// }
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Method untuk mencetak bilangan prima
func (d DeretBilangan) Prima() {
	fmt.Print("Bilangan Prima: ")
	for i := 2; i <= d.Limit; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// Method untuk mencetak bilangan ganjil
func (d DeretBilangan) Ganjil() {
	fmt.Print("Bilangan Ganjil: ")
	for i := 1; i <= d.Limit; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Method untuk mencetak bilangan genap
func (d DeretBilangan) Genap() {
	fmt.Print("Bilangan Genap: ")
	for i := 0; i <= d.Limit; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Method untuk mencetak deret Fibonacci
func (d DeretBilangan) Fibonacci() {
	fmt.Print("Fibonacci: ")
	a, b := 0, 1
	for a <= d.Limit {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
