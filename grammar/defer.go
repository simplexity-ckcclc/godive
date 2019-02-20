package grammar

import "fmt"

func a() {
	fmt.Print("a ")
}

func b() {
	fmt.Print("b ")
}

func c() {
	fmt.Print("c ")
}

// should print "a b c"
// defer called when leaving method instead of block
func DeferCalledWhenLeavingMethod() {
	if true {
		a()
		defer c()
	}
	b()
}
