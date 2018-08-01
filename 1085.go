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
	var x, y, w, h int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	l = strings.TrimSpace(l)
	pos := strings.Split(l, " ")
	x, _ = strconv.Atoi(pos[0])
	y, _ = strconv.Atoi(pos[1])
	w, _ = strconv.Atoi(pos[2])
	h, _ = strconv.Atoi(pos[3])

	minX := math.Min(float64(x), float64(w-x))
	minY := math.Min(float64(y), float64(h-y))
	fmt.Println(math.Min(minX, minY))
}
