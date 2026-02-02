package escape

import "testing"

//go:noinline
func foo4(a []string) {
	return
}
func TestEscape4(t *testing.T) {
	data := []string{"hello"}
	foo4(data)
}
