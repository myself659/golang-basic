package grcost

import (
	"testing"
)

func BenchPrintOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintOnly()
	}
}

func BenchPrintByGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintByGoRoutine()
	}
}
