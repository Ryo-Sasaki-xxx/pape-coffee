package page

type Page struct {
	id int
}

func NewPage(id int) (*Page, error) {
	page := &Page{}
	page.id = id

	return page, nil
}
