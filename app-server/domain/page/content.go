package page

type content struct {
	attributes []string
	value      string
}

func NewContent(attributes []string, value string) (*content, error) {

	content := &content{}
	content.attributes = attributes
	content.value = value

	return content, nil
}
