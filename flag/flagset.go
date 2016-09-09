package main

import (
	"flag"
	"fmt"
)

var (
	flagset = flag.NewFlagSet("flagset", errorHandling)
	num     = flagset.Int("num", 2, "num num ")
	name    = flagset.String("name", "foo", "name")
)

func main() {
	flagSet.Parse(os.Args[1:])

	fmt.Println(num)
	fmt.Println(name)
}
