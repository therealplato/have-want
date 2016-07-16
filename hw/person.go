package hw

type Person interface {
	IHave(Thing)
	IWant(Thing)
	DoIWant(Thing) bool
	DoIHave(Thing) bool
}
