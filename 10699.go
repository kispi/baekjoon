package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Split(time.Now().String(), " ")[0])
}
