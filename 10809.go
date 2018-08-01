package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Alpha struct {
	Key            string
	FirstOccurence int
}

func main() {
	alphamap := make(map[string]*Alpha)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := strings.ToLower(strings.TrimSpace(scanner.Text()))
	for i, r := range word {
		if _, ok := alphamap[string(r)]; !ok {
			alphamap[string(r)] = &Alpha{
				string(r),
				i,
			}
		}
	}

	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if v, ok := alphamap[string(r)]; ok {
			fmt.Printf("%d ", v.FirstOccurence)
		} else {
			fmt.Printf("%d ", -1)
		}
	}
}
