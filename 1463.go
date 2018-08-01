package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const max = 3000001

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	n, _ := strconv.Atoi(l)

	intArray := [max]int{}
	fillArray(n, &intArray)
	fmt.Print(intArray[n])
}

func fillArray(n int, intArray *[max]int) {
	intArray[0] = 0
	intArray[1] = 0
	for i := 1; i <= n; i++ {
		if (intArray[i*3] > intArray[i]+1) || intArray[i*3] == 0 {
			intArray[i*3] = intArray[i] + 1
		}
		if (intArray[i*2] > intArray[i]+1) || intArray[i*2] == 0 {
			intArray[i*2] = intArray[i] + 1
		}
		if (intArray[i+1] > intArray[i]+1) || intArray[i+1] == 0 {
			intArray[i+1] = intArray[i] + 1
		}
	}
}
