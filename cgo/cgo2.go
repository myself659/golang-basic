package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"
import "fmt"

func main() {
	str := C.CString("Hello, world!\n")
	defer C.free(unsafe.Pointer(str))
	fmt.Println(C.GoString(str))
}
