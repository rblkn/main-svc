package main

import (
	"fmt"

	"github.com/rblkn/lib/adder"
	"github.com/rblkn/lib/power"
)

func main() {
	fmt.Printf("%d + %d = %d\n", 1, 2, adder.Add(1, 2))
	fmt.Printf("%d^%d = %.2f\n", 4, 2, power.Pow(4., 2.))
}
