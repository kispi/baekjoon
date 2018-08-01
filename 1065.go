package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isHan(n int64) bool {
	if n >= 1 && n <= 99 {
		return true
	}

	f := strconv.FormatInt(n, 10)
	d := 0
	for i := range f {
		cur, _ := strconv.Atoi(string(f[i]))
		if i > 0 {
			prev, _ := strconv.Atoi(string(f[i-1]))
			if cur-prev != d && i >= 2 {
				return false
			}
			d = cur - prev
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	n, _ := strconv.ParseInt(l, 10, 64)

	total := 0
	for i := 1; int64(i) <= n; i++ {
		if isHan(int64(i)) {
			total++
		}
	}

	fmt.Println(total)
}
