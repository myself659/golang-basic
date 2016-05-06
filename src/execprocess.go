package main

import "fmt"

//import "io/ioutil"
import "os/exec"

func main() {

	dircmd := exec.Command("dir")
	dirout, err := dircmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("dir= ")
	fmt.Println(string(dirout))
}

/* 与 */
