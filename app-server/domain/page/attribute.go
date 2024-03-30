package page

import "errors"

type attribute interface {
}

func NewAttribute(attributeName string) (attribute, error) {
	if attributeName == "bold" {
		return NewBold(), nil
	}

	// if attributeName == "color" {
	// 	return NewColor(), nil
	// }

	return nil, errors.New("no attribute")
}

type bold struct {
}

func NewBold() bold {
	return bold{}
}
