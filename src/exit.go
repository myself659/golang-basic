package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("defer done")

	os.Exit(3)

}
