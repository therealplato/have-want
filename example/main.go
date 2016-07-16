package main

import (
	"github.com/therealplato/have-want/hw"
	"github.com/therealplato/have-want/id"
)

func main() {
	a := hw.MockPerson{
		ID: id.Email{title: "foo@bar.com"},
	}
	a.IWant(Thing{"banana"})

	b := hw.MockPerson{
		ID: id.Email{title: "baz@bin.com"},
	}
	b.IHave(Thing{"banana"})

	hw.Hookup(a, b)

}
