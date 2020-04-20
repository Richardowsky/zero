package src

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var n = 4
var array = []int{1, 1, 2, 2}
var n2 = 6
var array2 = []int{1, 2, 3, 4, 5, 6}

func Test1(t *testing.T) {
	assert.Equal(t, CheckZero(n, array), "YES")
}

func Test2(t *testing.T) {
	assert.Equal(t, CheckZero(n2, array2), "NO")
}

func Benchmark1(b *testing.B) {
	benchmarkZero(n, array, b)
}

func Benchmark2(b *testing.B) {
	benchmarkZero(n2, array2, b)
}

func benchmarkZero(n int, array []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckZero(n, array)
	}
}
