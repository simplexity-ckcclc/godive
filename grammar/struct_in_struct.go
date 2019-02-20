package grammar

import "fmt"

type foo struct {
	name string
}

type bar struct {
	foo
	int int32
}

// should print ckcclc
// ref args of inner struct by name as if the args is of outer struct
func RefArgsOfStructInStruct() {
	mybar := &bar{
		foo: foo{
			name: "ckcclc",
		},
	}

	fmt.Println(mybar.name)
}
