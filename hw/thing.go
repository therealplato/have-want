package hw

// Thing interface describes behavior of anything you can put on a have/want board
type Thing interface {
	Title() string
	Description() string
	IsTangible() bool
}

// MockThing is an in-memory development implementation of Thing
type MockThing struct {
	title   string
	desc    string
	digital bool
}

// Title returns a short title of the thing suitable for listing on a h/w board
func (t *MockThing) Title() string {
	return t.title
}

// Description contains details about the thing
func (t *MockThing) Description() string {
	return t.desc
}

// IsTangible returns true if the thing can be physically touched
func (t *MockThing) IsTangible() bool {
	return !t.digital
}
