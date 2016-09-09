package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	str := C.CString("Hello, world!\n")
	defer C.free(unsafe.Pointer(str))
	C.printf("hello")
}
