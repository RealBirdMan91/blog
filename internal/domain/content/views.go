package content

import "errors"

var ErrNegativeViewsCount = errors.New("count is not allowed to be negative")

type Views struct {
	n int32
}

func NewViews(n int32) (Views, error) {
	if n < 0 {
		return Views{}, ErrNegativeViewsCount
	}
	return Views{n: n}, nil
}

func (v Views) Int() int32 { return v.n }
