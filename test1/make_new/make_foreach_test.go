package make_new

import (
	"fmt"
	"testing"
)

/*
测试foreach会不会出现问题，1.22版本后就修复了，不会出现问题。
*/
type student struct {
	Name string
	Age  int
}

func TestMakeForeach(t *testing.T) {
	m := make(map[string]*student)

	stus := []student{
		{Name: "tang", Age: 11},
		{Name: "xin", Age: 22},
		{Name: "kang", Age: 33},
	}
	for _, v := range stus {
		m[v.Name] = &v
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}

}
