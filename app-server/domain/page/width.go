package page

import "errors"

type width struct {
	value int
}

func NewWidth(value int) (*width, error) {
	if value > 10 {
		return nil, errors.New("１０以上")
	}

	width := &width{}
	width.value = value
	return width, nil
}

func (width width) Add(other width) (*width, error) {
	newWidth, err := NewWidth(width.value + other.value)

	if err != nil {
		return nil, err
	}

	return newWidth, nil
}
