package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Values [][]int
}

func (m *Matrix) Print() {
	for _, row := range m.Values {
		fmt.Println(row)
	}
}

func main() {
	var n, m int
	var mat1, mat2 Matrix
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()

	l = strings.TrimSpace(l)
	pos := strings.Split(l, " ")
	n, _ = strconv.Atoi(pos[0])
	m, _ = strconv.Atoi(pos[1])

	fmt.Println(m)

	for i := 0; i < n; i++ {
		scanner.Scan()
		nums := strings.Split(strings.TrimSpace(scanner.Text()), "")
		row := []int{}
		for _, numch := range nums {
			num, _ := strconv.Atoi(numch)
			row = append(row, num)
		}
		mat1.Values = append(mat1.Values, row)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		nums := strings.Split(strings.TrimSpace(scanner.Text()), "")
		row := []int{}
		for _, numch := range nums {
			num, _ := strconv.Atoi(numch)
			row = append(row, num)
		}
		mat2.Values = append(mat2.Values, row)
	}

	mat1.Print()
	mat2.Print()
}
