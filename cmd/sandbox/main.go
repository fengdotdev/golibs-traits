package main

import (
	"fmt"
	"reflect"
)

type opaque struct {
	content string
}

type Foo struct {
	Bar    string
	baz    string
	opaque opaque
}

func NewFoo() Foo {
	return Foo{
		Bar:    "--bar",
		baz:    "--baz",
		opaque: opaque{content: "--opaque"},
	}
}

func (f Foo) Baz() string {
	return f.baz
}

type Bazaer interface {
	Baz() string
}

func main() {
	var f Bazaer = Foo{
		Bar: "--bar",
		baz: "--baz",
	}

	r := reflect.TypeOf(f)
	fmt.Println(r.Name())        // Foo
	fmt.Println(r.Kind())        // struct
	fmt.Println(r.NumField())    // 2
	fmt.Println(r.Field(2).Type) // string
	fmt.Println(r.Field(2).Name) // Bar
	fmt.Println(r.Field(2).Tag)  // <nil>
	fmt.Println(r.Field(2).Offset) // 0
	fmt.Println(r.Field(2).Index) // [0 0]
	fmt.Println(r.Field(2).Type.Kind()) // string
	fmt.Println(r.Field(2).Type.Name()) // string
	fmt.Println(r.Field(2).Type.String()) // string
	fmt.Println(r.Field(2).Type.Elem()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Kind()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Name()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().String()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Size()) // 0
	fmt.Println(r.Field(2).Type.Elem().Align()) // 0
	fmt.Println(r.Field(2).Type.Elem().Field(0)) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Name) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Type) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Tag) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Offset) // 0
	fmt.Println(r.Field(2).Type.Elem().Field(0).Index) // [0]
	fmt.Println(r.Field(2).Type.Elem().Field(0).Type.Kind()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Type.Name()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Type.String()) // <nil>
	fmt.Println(r.Field(2).Type.Elem().Field(0).Type.Elem()) // <nil>
	
	// content of the field
	//v := reflect.ValueOf(f).Field(2)
	//fmt.Println(v.Interface()) // --bar

}
