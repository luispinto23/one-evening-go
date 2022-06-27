package main

// What to use?

//As a rule of thumb:
//
//    If your methods don't modify the struct, use plain struct receivers.
//    If any of your methods modifies the struct, use pointer receivers everywhere for consistency.
//
//You might also use pointer receivers if your struct is huge and copying it on each method call would be expensive.
//You don't have to concern about it for now.

type Position struct {
	X int
	Y int
}

func main() {
	p := Position{X: 10, Y: 20}
	p.Move(5, -5)
}

func (p *Position) Move(dx, dy int) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}
