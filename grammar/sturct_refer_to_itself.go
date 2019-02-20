package grammar

import "fmt"

type refIA interface {
	f(str string)
}

type refIB interface {
	refIA
}

type refSA struct {
	refB refIB
}

func newSA(refB refIB) *refSA {
	sa := &refSA{
		refB: refB,
	}
	return sa
}

type refSB struct {
	str string
	*refSA
	*refSC
}

type refSC struct {
	refIA
}

type refSD struct {
}

func (sd *refSD) f(str string) {
	fmt.Println("refSD:", str)
}

// should print refSD: test
// struct can ref interface(invoke methods) if the arg is anonymous
func RefStructToItself() {
	sb := &refSB{
		str: "ckcclc",
	}

	sd := &refSD{}
	sc := &refSC{
		refIA: sd,
	}
	sb.refSC = sc

	sb.refSA = newSA(sb)
	sb.f("test")
}
