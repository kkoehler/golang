package math_test

import (
	"fmt"
	"strconv"
	"testing"

	"golang.source-fellows.com/samples/unittest/math"
)

var result int

func BenchmarkAdd(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = math.Add(i, i)
	}
	result = r
}

func BenchmarkHelloWorld(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", i)
	}

}

func BenchmarkHelloWorldLong(b *testing.B) {

	for i := 0; i < b.N; i++ {
		strconv.Itoa(i)
	}

}
