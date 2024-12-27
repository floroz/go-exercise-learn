package main

import "fmt"

type Foo struct {
	eventType string
	foo       string
}

type Bar struct {
	eventType string
	bar       string
}

func typeSwitch[T Bar | Foo](event T) {
	switch v := any(event).(type) {
	case Bar:
		fmt.Printf("bar: %s\n", v.bar)
	case Foo:
		fmt.Printf("foo: %s\n", v.foo)
	default:
		panic("unknown type")
	}
}

func typeAssertion[T Bar | Foo](event T) {
	if v, ok := any(event).(Bar); ok {
		fmt.Printf("bar: %s\n", v.bar)
	} else if v, ok := any(event).(Foo); ok {
		fmt.Printf("foo: %s\n", v.foo)
	} else {
		panic("unknown type")
	}
}

func main() {
	typeSwitch(Bar{eventType: "A", bar: "B"})
	typeSwitch(Foo{eventType: "Something else", foo: "B"})
	typeAssertion(Bar{eventType: "A", bar: "B"})
	typeAssertion(Foo{eventType: "Something else", foo: "B"})
}
