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

type Position struct {
	x          int
	y          int
	distToAvgX int
	distToAvgY int
}

type PositionStats struct {
	Positions []*Position
	MinMovs   []int
	AvgX      int
	AvgY      int
}

func (ps *PositionStats) Len() int {
	return len(ps.Positions)
}

func (ps *PositionStats) Less(i, j int) bool {
	return ps.Positions[i].distToAvgX+ps.Positions[i].distToAvgY < ps.Positions[j].distToAvgX+ps.Positions[j].distToAvgY
}

func (ps *PositionStats) Swap(i, j int) {
	ps.Positions[i], ps.Positions[j] = ps.Positions[j], ps.Positions[i]
}

func (ps *PositionStats) getStats() {
	if len(ps.Positions) == 0 {
		return
	}

	sumx := float64(0)
	sumy := float64(0)
	for _, p := range ps.Positions {
		sumx += float64(p.x)
		sumy += float64(p.y)
	}
	ps.AvgX = int(math.Round(sumx / float64(len(ps.Positions))))
	ps.AvgY = int(math.Round(sumy / float64(len(ps.Positions))))

	for _, p := range ps.Positions {
		p.distToAvgX = int(math.Abs(float64(p.x - ps.AvgX)))
		p.distToAvgY = int(math.Abs(float64(p.y - ps.AvgY)))
	}
}

func (ps *PositionStats) getMinimalMovements() {
	ps.MinMovs = append(ps.MinMovs, 0) // Trivial. Always will be 0.
	for matched := 2; matched <= len(ps.Positions); matched++ {
		movs := 0
		for j := 0; j < matched; j++ {
			movs += int(math.Abs(float64(ps.Positions[j].x - ps.AvgX)))
			movs += int(math.Abs(float64(ps.Positions[j].y - ps.AvgY)))
		}
		ps.MinMovs = append(ps.MinMovs, movs)
	}
}

func (ps *PositionStats) print() {
	fmt.Printf("AverageX: %d, AverageY: %d\n", ps.AvgX, ps.AvgY)
	for _, p := range ps.Positions {
		fmt.Printf("X: %d, Y: %d, Xd: %d, Yd: %d\n", p.x, p.y, p.distToAvgX, p.distToAvgY)
	}
}

func main() {
	positions := []*Position{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	n, _ := strconv.Atoi(l)
	for i := 0; i < n; i++ {
		scanner.Scan()
		l = scanner.Text()
		l = strings.TrimSpace(l)
		strs := strings.Split(l, " ")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		positions = append(positions, &Position{x: x, y: y})
	}

	ps := new(PositionStats)
	ps.Positions = positions
	ps.getStats()
	sort.Sort(ps)
	ps.getMinimalMovements()

	ps.print()

	result := fmt.Sprintf("%v", ps.MinMovs)
	result = strings.Replace(result, "[", "", 1)
	result = strings.Replace(result, "]", "", 1)
	fmt.Println(result)
}
