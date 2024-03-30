package page

func (pageTitle pageTitle) ExportTitle() string {
	return pageTitle.title
}

func (block TextBlock) ExportContent() string {
	return block.content
}

func (content content) ExportValue() string {
	return content.value
}

func (component component) ExportBlock() []Block {
	return component.blocks
}

func (width width) ExportValue() int {
	return width.value
}
