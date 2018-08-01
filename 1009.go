package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	A   int
	B   int
	Mul int
}

func (p *Pair) get1st() int {
	if p.B == 1 {
		return p.A
	}

	p.A = (p.A * p.Mul) % 10
	if p.A == 0 {
		return 10
	}

	p.B--
	return p.get1st()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	n, _ := strconv.ParseInt(l, 10, 64)

	pairs := []*Pair{}
	for i := 1; int64(i) <= n; i++ {
		scanner.Scan()
		l = scanner.Text()
		nums := strings.Split(strings.TrimSpace(l), " ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		p := &Pair{a, b, a}
		pairs = append(pairs, p)
	}

	for _, p := range pairs {
		fmt.Println(p.get1st())
	}
}
