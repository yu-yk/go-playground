package main

import (
	"fmt"
	"math"
)

func main() {
	x := 12.3456
	fmt.Println(math.Floor(x*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(x*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(x*100) / 100)  // 12.35 (round up)
}
