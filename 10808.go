package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Alpha struct {
	Key       string
	Occurence int
}

func main() {
	alphamap := make(map[string]*Alpha)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := strings.ToLower(strings.TrimSpace(scanner.Text()))
	for _, r := range word {
		if _, ok := alphamap[string(r)]; !ok {
			alphamap[string(r)] = &Alpha{
				string(r),
				1,
			}
		} else {
			alphamap[string(r)].Occurence++
		}
	}

	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if v, ok := alphamap[string(r)]; ok {
			fmt.Printf("%d ", v.Occurence)
		} else {
			fmt.Printf("%d ", 0)
		}
	}
}
