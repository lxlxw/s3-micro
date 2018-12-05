package log

import (
	"fmt"
	"testing"
)

type B struct {
	Name string
	Age  int
}

type A struct {
	BV   B
	BPtr *B
	I    interface{}
	Iptr interface{}
}

func TestPrintStruct(t *testing.T) {
	b := B{Name: "elvizlai", Age: 26}
	fmt.Println(PrintStruct(b))
	a := A{BV: b}
	fmt.Println(PrintStruct(a))
	a.BPtr = &b
	fmt.Println(PrintStruct(a))
	a.I = b
	a.Iptr = &b
	fmt.Println(PrintStruct(&a))
}
