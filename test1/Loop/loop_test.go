package Loop

import "testing"

func CreateSource(len int) []int {
	nums := make([]int, 0, len)
	for i := 0; i < len; i++ {
		nums = append(nums, i)
	}
	return nums
}

func BenchmarkLoopStep1(b *testing.B) {
	//制作源数据，长度为 10000
	src := CreateSource(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 1)
	}
}
func BenchmarkLoopStep4(b *testing.B) {
	//制作源数据，长度为 10000
	src := CreateSource(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 4)
	}
}
func BenchmarkLoopStep8(b *testing.B) {
	//制作源数据，长度为 10000
	src := CreateSource(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 8)
	}
}
func BenchmarkLoopStep16(b *testing.B) {
	//制作源数据，长度为 10000
	src := CreateSource(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 16)
	}
}
func BenchmarkLoopStep32(b *testing.B) {
	//制作源数据，长度为 10000
	src := CreateSource(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 32)
	}
}
func BenchmarkLoopStep40(b *testing.B) {
	//制作源数据，长度为 10000
	src := CreateSource(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 40)
	}
}
