package make_new

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
}

var list = map[string]Student{}

func TestMake(t *testing.T) {
	list = make(map[string]Student)
	stu := Student{
		Name: "xxx",
	}
	list["student"] = stu
	tmp := list["student"]
	tmp2 := list["student"]
	fmt.Printf("stu 的地址: %p\n", &stu)
	fmt.Printf("tmp 的地址: %p\n", &tmp)
	fmt.Printf("tmp2 的地址: %p\n", &tmp2)

}
