package grammar

import "fmt"

type defaultValStruct struct {
	int int32
}

// should print 0
// default value of int32 is 0
func DefaultVal() {
	s := &defaultValStruct{}
	fmt.Print(s.int)
}
