package main

import (
	"fmt"

	"github.com/therealplato/have-want/hw"
)

func main() {
	a := &hw.MockPerson{
		ID: "foo@bar.com",
	}
	a.IWant(&hw.MockThing{
		Title_: "banana",
	})

	b := &hw.MockPerson{
		ID: "baz@bin.com",
	}
	b.IHave(&hw.MockThing{
		Title_: "banana",
	})

	common := hw.Hookup(a, b)
	fmt.Println(common)

}
