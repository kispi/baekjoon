package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	A     int
	B     int
	Sum   int
	Prime bool
}

var possiblePairs [][]*Pair

func remove(numbers []int, i int) []int {
	if len(numbers) <= 1 {
		return []int{}
	}

	return append(numbers[:i], numbers[i+1:]...)
}

func NewPair(a int, b int) *Pair {
	return &Pair{
		A:     a,
		B:     b,
		Sum:   a + b,
		Prime: isPrime(a + b),
	}
}

func main() {
	numbers := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		l := scanner.Text()
		if i == 1 {
			l = strings.TrimSpace(l)
			strs := strings.Split(l, " ")
			for _, n := range strs {
				i, _ := strconv.ParseInt(n, 10, 64)
				numbers = append(numbers, int(i))
			}
		}
	}

	pairs := []*Pair{}
	process(numbers, pairs)

	result := []int{}
	cur := 0
	for _, ps := range possiblePairs {
		p := ps[0]
		if cur != p.B {
			result = append(result, p.B)
			cur = p.B
		}
	}

	sort.Ints(result)
	fmt.Println(strings.Replace(strings.Replace(fmt.Sprintf("%v", result), "[", "", 1), "]", "", 1))
}

func process(numbers []int, pairs []*Pair) {
	if len(numbers) < 2 {
		prime := 0
		for _, p := range pairs {
			if p.Prime {
				prime++
			}
		}
		if prime == len(pairs) && !alreadyExists(pairs) {
			possiblePairs = append(possiblePairs, pairs)
		}
		return
	}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			arrCp := make([]int, len(numbers))
			copy(arrCp, numbers)
			a := arrCp[i]
			b := arrCp[j]
			if isPrime(a + b) {
				arrCp = remove(arrCp, int(j))
				arrCp = remove(arrCp, int(i))

				newPairs := append(pairs, NewPair(a, b))
				process(arrCp, newPairs)
			}
		}
	}
	return
}

func alreadyExists(arg []*Pair) bool {
	for _, pairs := range possiblePairs {
		s1 := convertToIntSliceFromPairs(arg)
		s2 := convertToIntSliceFromPairs(pairs)
		if isSameIntSlice(s1, s2) {
			return true
		}
	}
	return false
}

func isSameIntSlice(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func convertToIntSliceFromPairs(pairs []*Pair) []int {
	me := []int{}
	for _, p := range pairs {
		me = append(me, p.Sum)
	}

	sort.Ints(me)
	return me
}

func isPrime(n int) bool {
	if n == 1 || n%2 == 0 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(n)))); i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
