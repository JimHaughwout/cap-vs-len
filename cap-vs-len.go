package main

import (
	"fmt"
)


func print_stats(s []int) {
	fmt.Printf("cap %v, len %v, addr %p, slice %v\n", cap(s), len(s), s, s)
}

func main() {
	start_size := 3
	append_size := 20
	fmt.Printf("Base size is %d, incrementally appending %d elements\n", 
		start_size, append_size)

	s := make([] int, 0, start_size)
	print_stats(s)

	for i := 0; i < append_size; i++ {
		s = append(s, i)
		print_stats(s)
	}
}