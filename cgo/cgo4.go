package main

/*
#include "stdlib.h"

void * newPtrArray(int size){
	void *pa;
	pa = malloc(size * sizeof(void*));
	return pa;
}
void  setByOffset(void  **ppa,  int offset, void *p){
	ppa[offset] = p;
}
void *getByOffset(void **pa, int offset){
	return pa[offset];
}
void delPtrArray(void **ppa){
	free(ppa);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Point struct {
	x int
	y int
}

func main() {
	fmt.Println("start")
	var p *unsafe.Pointer
	p = (*unsafe.Pointer)(C.newPtrArray(16))
	fmt.Println(p)
	fmt.Println(&p)
	temp := new(Point)
	temp.x = 1
	temp.y = 1
	C.setByOffset(p, 1, unsafe.Pointer(temp))
	ret := unsafe.Pointer(C.getByOffset(p, 1))
	fmt.Println(ret)
	pp := (*Point)(ret)
	fmt.Println(pp)
	// p32 := &new(unsafe.Pointer)
	p32 := (*unsafe.Pointer)(new(unsafe.Pointer))
	fmt.Println(&p32)
	fmt.Printf("&p32:%t\n", &p32)
	pp32 := &p32
	p32 = (*unsafe.Pointer)(C.newPtrArray(32))
	fmt.Printf("p32:%#v\n", p32)
	C.setByOffset(p32, 1, unsafe.Pointer(temp))
	ret = unsafe.Pointer(C.getByOffset(*pp32, 1))
	fmt.Println(ret)
	p1 := (*Point)(ret)
	fmt.Println(p1)
	C.delPtrArray(p32)
	C.delPtrArray(p)

}
