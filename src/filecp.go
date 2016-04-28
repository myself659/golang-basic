package main

import "fmt"
import "os"
import "io"

func usage() {
	fmt.Printf("%s %s %s\n", os.Args[0], "filename", "newfile")
}

func main() {

	if len(os.Args) != 3 {
		usage()
		return
	}

	filename_in := os.Args[1]
	fi, err := os.Open(filename_in)
	if err != nil {
		panic(err)
	}
	/* 当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回 */
	defer fi.Close()

	filename_out := os.Args[2]
	fo, err := os.Create(filename_out)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	/* make */
	var buff = make([]byte, 1024)
	for {
		n, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		if _, err := fo.Write(buff[:n]); err != nil {
			panic(err)
		}

	}
}
