package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Alpha struct {
	Key   string
	Value int
}

func main() {
	alphas := []*Alpha{}
	alphamap := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1000005)
	scanner.Buffer(buf, 1000005)
	scanner.Scan()
	word := strings.ToUpper(strings.TrimSpace(scanner.Text()))
	for _, r := range word {
		alphamap[string(r)]++
	}

	for k, v := range alphamap {
		alphas = append(alphas, &Alpha{k, v})
	}

	sort.Slice(alphas, func(i, j int) bool {
		return alphas[i].Value >= alphas[j].Value
	})

	if len(alphas) > 1 {
		if alphas[0].Value == alphas[1].Value {
			fmt.Println("?")
		} else {
			fmt.Printf("%s\n", alphas[0].Key)
		}
		return
	}

	if len(alphas) > 0 {
		fmt.Printf("%s\n", alphas[0].Key)
	}
}
