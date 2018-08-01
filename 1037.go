package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	ints := []int{}

	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	values := strings.Split(line, " ")
	for i := range values {
		v, _ := strconv.Atoi(values[i])
		ints = append(ints, v)
	}

	sort.Ints(ints)
	fmt.Println(ints[0] * ints[len(ints)-1])
}
