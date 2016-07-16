package hw

var w = `I want a vim script that behaves as follows:
Given <filename>_test.go does not exist
And the cursor is in <filename>.go
When <Leader>, is pressed
Then a file is created with the name <filename>_test.go
And that file becomes the active buffer`

// Hookup returns things that a wants and b has,and vice versa
func Hookup(a, b Person) []Thing {
	common := []Thing{}
	h1, w1 := a.List()
	h2, w2 := b.List()
	for _, need := range w1 {
		for _, offer := range h2 {
			if offer.Equal(need) {
				common = append(common, offer)
			}
		}
	}
	for _, need := range w2 {
		for _, offer := range h1 {
			if offer.Equal(need) {
				common = append(common, offer)
			}
		}
	}
	return common
}
