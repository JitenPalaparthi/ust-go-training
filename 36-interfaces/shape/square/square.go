package square

import "errors"

type Square struct {
	Side float32
}

func New(s float32) (*Square, error) {
	if s == 0 {
		return nil, errors.New("invalid data")
	}
	return &Square{Side: s}, nil
}

func (s *Square) Area() float32 {
	return s.Side * s.Side
}
