package main

import (
	"fmt"

	"github.com/rblkn/main-svc/adder"
)

func main() {
	fmt.Printf("%d + %d = %d\n", 1, 2, adder.Add(1, 2))
}
