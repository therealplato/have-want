package hw

import (
	"fmt"
	"strings"
)

// Thing interface describes behavior of anything you can put on a have/want board
type Thing interface {
	Title() string
	Description() string
	IsTangible() bool
	IsDigital() bool
	Equal(Thing) bool
}

// MockThing is an in-memory development implementation of Thing
type MockThing struct {
	Title_   string
	Desc_    string
	Digital  bool
	Tangible bool
}

// Title returns a short title of the thing suitable for listing on a h/w board
func (t *MockThing) Title() string {
	return t.Title_
}

// Description contains details about the thing
func (t *MockThing) Description() string {
	return t.Desc_
}

// IsTangible returns true if the thing can be physically touched
func (t *MockThing) IsTangible() bool {
	return t.Tangible
}

// IsDigital returns true if the thing can be represented by bits
func (t *MockThing) IsDigital() bool {
	return t.Digital
}

// Equal returns true if thing t is the same as thing x
func (t *MockThing) Equal(x Thing) bool {
	t1 := strings.ToLower(strings.TrimSpace(t.Title()))
	x1 := strings.ToLower(strings.TrimSpace(x.Title()))
	return t1 == x1
}

func (t *MockThing) String() string {
	var tmp string
	if t.Digital {
		tmp = "Digital"
	} else if t.Tangible {
		tmp = "Tangible"
	} else {
		tmp = "Abstract"
	}
	return fmt.Sprintf("%s %s:\n%s", tmp, t.Title(), t.Description())
}
