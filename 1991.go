package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Left  = "left"
	Right = "right"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func remove(nodes []*Node, i int) []*Node {
	if len(nodes) <= 1 {
		return []*Node{}
	}

	return append(nodes[:i], nodes[i+1:]...)
}

func buildTree(nodes []*Node) *Node {
	for i := range nodes {

	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := strings.TrimSpace(scanner.Text())
	n, _ := strconv.Atoi(l)

	nodes := []*Node{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		l := strings.Split(scanner.Text(), " ")

		nodes = append(nodes, &Node{
			Value: l[0],
			Left: &Node{
				Value: l[1],
			},
			Right: &Node{
				Value: l[2],
			},
		})
	}

	for _, node := range nodes {
		fmt.Println(node)
	}
}
