package grammar

import "fmt"

type myInterfaceFoo interface {
	Foo(str string)
}

type myInterfaceBar interface {
	Bar(str string)

	// can NOT have same method signature if myStruct has two interface
	// Foo(str string)
}

type myStruct struct {
	foo
	int int32
	myInterfaceFoo
	myInterfaceBar
}

type myFoo struct {
}

func (foo *myFoo) Foo(str string) {
	fmt.Println("Foo :", str)
}

type myBar struct {
}

func (bar *myBar) Bar(str string) {
	fmt.Println("Bar :", str)
}

// should print
//      Foo : ckcclc
//      Bar : ckcclc
// can call interface method directly,
// but, can NOT have same method signature if myStruct has two interface
func CallMethodOfInterfaceInStruct() {
	s := &myStruct{
		myInterfaceFoo: new(myFoo),
		myInterfaceBar: new(myBar),
	}
	s.Foo("ckcclc")
	s.Bar("ckcclc")
}
