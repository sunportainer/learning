package main

import (
	"fmt"
	"sort"
)

func print() {
	fmt.Println("dsafasdfsadfasd")
}

func main() {
	array := []int{3, 2, 4}
	sort.Ints(array)
	fmt.Println(array)
}
