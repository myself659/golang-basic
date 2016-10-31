package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Point struct {
	x string
	y string
}

func main() {
	records := []Point{
		{x: "abc", y: "def"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
