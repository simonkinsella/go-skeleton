package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go version '%s' says Hello!", runtime.Version())
}
