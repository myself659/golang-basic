package main

import "sort"
import "fmt"

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	techs := []string{"docker", "cloud compute", "cdn"}

	sort.Sort(ByLength(techs)) /* 这个细节需要读go源码 */

	fmt.Println(techs)
}
