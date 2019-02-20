package grammar

import "fmt"

type assertInf interface {
	f()
}

type assertFoo struct {
	name string
}

type assertBar struct {
	int int32
}

func (foo *assertFoo) f() {
	fmt.Print(foo.name)
}

func (bar *assertBar) f() {
	fmt.Print(bar.int)
}

func AssertType() {
	bar := &assertBar{
		int: 123,
	}

	assertF(bar)
}

// should print
//      assert type match: false
//      assert: <nil>
// if type assert fail, the result will be nil and false
func assertF(inf assertInf) {
	assert, ok := inf.(*assertFoo)
	fmt.Println("assert:", assert)
	fmt.Println("assert type match:", ok)
}
