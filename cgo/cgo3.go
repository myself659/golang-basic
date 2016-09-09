package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char * format(char *s){
	char *buf;
	buf  = malloc(strlen(s));
	sprintf(buf, "%s", s);
	return buf;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	cs := C.format(C.CString("hello, cgo\n"))
	fmt.Println(C.GoString(cs))
	C.free(unsafe.Pointer(cs))
}
