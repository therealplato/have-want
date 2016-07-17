have-want
=========

I have stuff I don't need.

I have wants I don't have.

This project aims to use digital Have/Want boards to connect people who have
something with people who want it.


Planned Usage
-------------

I := hw.NewPerson()
I.Want("ferarris")
I.Have("USD")

U := hw.NewPerson()
// find existing opportunities:
ops := hw.Hookup(I, U)

// get notified of future opportunities?
I.Seek()

for op := <- I.Opportunities() {
  log.Println(op)
}
