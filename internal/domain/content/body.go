package content

import (
	"errors"
	"strings"
)

var (
	ErrEmptyBody    = errors.New("body is empty")
	ErrBodyTooShort = errors.New("body to short at least: 50 chars")
	ErrBodyTooLong  = errors.New("body is to long: max 1500 chars")
)

const (
	bodyMinLen = 50
	bodyMaxLen = 1500
)

type Body struct{ s string }

func NewBody(raw string) (Body, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return Body{}, ErrEmptyBody
	}
	if len([]rune(trimmed)) < bodyMinLen {
		return Body{}, ErrBodyTooShort
	}
	if len([]rune(trimmed)) > bodyMaxLen {
		return Body{}, ErrBodyTooLong
	}
	return Body{s: trimmed}, nil
}

func (b Body) String() string {
	return b.s
}
