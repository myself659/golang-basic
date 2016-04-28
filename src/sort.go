package main

import "fmt"
import "sort"

func main() {

	strs := []string{"c", "de", "a", "bc"}

	sort.Strings(strs)
	fmt.Println("sort strs:", strs)

	ints := []int{30, 20, 13, 10, 7, 6, 3}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	bsort := sort.IntsAreSorted(ints)
	fmt.Println("soted:", bsort)

}
