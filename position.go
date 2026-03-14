package goscii

type Position struct {
	X, Y float64
}

func NewPosition(x, y float64) *Position {
	return &Position{X: x, Y: y}
}

func (p *Position) Add(pos *Position) *Position {
	return &Position{
		X: p.X + pos.X,
		Y: p.Y + pos.Y,
	}
}

func (p *Position) Sub(pos *Position) *Position {
	return &Position{
		X: p.X - pos.X,
		Y: p.Y - pos.Y,
	}
}

func (p *Position) Set(x, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
