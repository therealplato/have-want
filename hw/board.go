package hw

type Board struct {
	owner Person
	haves []Thing
	wants []Thing
}

// Have adds t to b's have list
func (b *Board) Have(t Thing) {
	for _, i := range b.haves {
		if i.Equal(t) {
			return
		}
	}
	b.haves = append(b.haves, t)
}

// IWant adds t to p's want list
func (b *Board) IWant(t Thing) {
	for _, i := range b.wants {
		if i.Equal(t) {
			return
		}
	}
	b.wants = append(b.wants, t)
}

// List returns p's haves and wants
func (b *Board) List() (haves, wants []Thing) {
	return b.haves, b.wants
}
