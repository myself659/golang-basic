package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
)

var data = [][]string{{"Line1", "Hello Readers of:"}, {"Line2", "golangcode.com"}}

func fix(fv float64, a float64) float64 {
	fr := math.Floor((fv / a)) * a
	return fr
}
func main() {
	fv := 1.23456
	fmt.Println(math.Floor(fv))
	fmt.Println(fv / 0.01)
	fmt.Println(fix(fv, 0.01))
	fmt.Println(fix(fv, 0.1)) //为什么不是1.2 而是1.20000002
	fmt.Printf("%f\n", fix(fv, 0.1))
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
