package content

import "errors"

var ErrNegativeLikesCount = errors.New("count is not allowed to be negative")

type Likes struct {
	n int32
}

func NewLikes(n int32) (Likes, error) {
	if n < 0 {
		return Likes{}, ErrNegativeLikesCount
	}
	return Likes{n: n}, nil
}

func (l Likes) Int() int32 { return l.n }
