package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l1 := scanner.Text()
	n, _ := strconv.ParseInt(l1, 10, 64)

	scanner.Scan()
	l2 := scanner.Text()
	f, _ := strconv.Atoi(l2)

	TwoNumbers, _ := strconv.ParseInt(l1[len(l1)-2:], 10, 64)
	startNum := n - TwoNumbers

	for i := 0; i < 100; i++ {
		curNum := startNum + int64(i)
		target := strconv.FormatInt(curNum, 10)
		if curNum%int64(f) == 0 {
			fmt.Println(target[len(target)-2:])
			break
		}
	}
}
