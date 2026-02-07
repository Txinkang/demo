package _interface

import (
	"fmt"
	"testing"
)

/*测试1，多态样例*/
type Phone interface {
	call()
}
type APhone struct{}

func (a APhone) call() {
	fmt.Println("APhone")
}

type BPhone struct{}

func (b BPhone) call() {
	fmt.Println("BPhone")
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = new(APhone)
	phone.call()

	phone = new(BPhone)
	phone.call()
}

/*测试2，赋值问题*/
type People interface {
	say(string) string
}
type BPeople struct{}

func (b *BPeople) say(thinks string) string {
	if thinks != "" {
		return "YES"
	} else {
		return "NO"
	}
}
func TestInterface2(t *testing.T) {
	var people People = &BPeople{}
	fmt.Println(people.say("1"))
}
