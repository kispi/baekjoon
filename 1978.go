package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	l := strings.TrimSpace(scanner.Text())
	strs := strings.Split(l, " ")
	for _, n := range strs {
		i, _ := strconv.ParseInt(n, 10, 64)
		numbers = append(numbers, int(i))
	}

	count := 0
	for _, n := range numbers {
		if isPrime(n) {
			count++
		}
	}
	fmt.Print(count)
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n == 1 || n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(n)))); i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
