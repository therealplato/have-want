package hw

import "github.com/therealplato/have-want/id"

// Person represents a person who has or wants things
type Person interface {
	IHave(Thing)
	IWant(Thing)
	// DoIWant(Thing) bool
	// DoIHave(Thing) bool
	Board() Board
}

// MockPerson is an in-memory implementation of Person
type MockPerson struct {
	Name  string
	ID    id.Email
	board Board
}

// Board retrieves this persons have want board
func (p *MockPerson) Board() Board {
	return p.board
}
