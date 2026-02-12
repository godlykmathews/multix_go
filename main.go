package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: multi <number>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%2d x %2d = %3d\n", i, n, i*n)
	}
}
