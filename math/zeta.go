package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var zeta, s = 0., 0.
		fmt.Print("Enter s: ")
		fmt.Scanln(&s)

		for i := 1.; i <= 10000000; i++ {
			zeta += 1 / math.Pow(i, s)
		}
		fmt.Println("zeta:        ", zeta)
		fmt.Println("sqrt(6*zeta):", math.Sqrt(6*zeta))
	}
}
