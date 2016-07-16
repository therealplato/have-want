package hw

import "github.com/therealplato/have-want/id"

// Person represents a person who has or wants things
type Person interface {
	IHave(Thing)
	IWant(Thing)
	// DoIWant(Thing) bool
	// DoIHave(Thing) bool
	List() ([]Thing, []Thing)
}

// MockPerson is an in-memory implementation of Person
type MockPerson struct {
	Name  string
	ID    id.Email
	haves []Thing
	wants []Thing
}

// IHave adds t to p's have list
func (p *MockPerson) IHave(t Thing) {
	for _, i := range p.haves {
		if i.Equal(t) {
			return
		}
	}
	p.haves = append(p.haves, t)
}

// IWant adds t to p's want list
func (p *MockPerson) IWant(t Thing) {
	for _, i := range p.wants {
		if i.Equal(t) {
			return
		}
	}
	p.wants = append(p.wants, t)
}

// List returns p's haves and wants
func (p *MockPerson) List() (haves, wants []Thing) {
	return p.haves, p.wants
}
