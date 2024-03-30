package page

import (
	"errors"
	"unicode/utf8"
)

type pageTitle struct {
	title string
}

func NewPageTitle(title string) (*pageTitle, error) {
	if len := utf8.RuneCountInString(title); len > 30 {
		return nil, errors.New("min length title error")
	}

	pageTitle := &pageTitle{}
	pageTitle.title = title

	return pageTitle, nil
}

func (pageTitle pageTitle) IsEmpty() bool {
	return pageTitle.title == ""
}
