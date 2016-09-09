package benchdemo

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkPrintOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintOnly()
	}
}

func BenchmarkPrintByGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintByGoRoutine()
	}
}

func BenchmarkAddDMap(b *testing.B) {
	fmt.Println("BenchAddDMap start")
	for i := 0; i < b.N; i++ {
		k := "aaa_" + strconv.Itoa(i)
		AddDMap(k, i)
	}
	fmt.Println("BenchAddDMap end")
}

func BenchmarkAddIMap(b *testing.B) {
	k := "aaa"
	ii := 100
	for i := 0; i < b.N; i++ {
		//k := "aaa_" + strconv.Itoa(i)
		AddIMap(k, ii)
	}
}

func BenchmarkConcat(b *testing.B) {
	var ss []string
	for n := 0; n < 100; n++ {
		ss = append(ss, "foo")
	}

	for i := 0; i < b.N; i++ {
		Concat(ss...)
	}
}
