package page

import "errors"

type Block interface {
	Type() string
}

func NewBlock(blockType string) (Block, error) {
	if blockType == "text" {
		block, _ := NewTextBlock("")
		return block, nil
	}

	return nil, errors.New("no type")

}

type TextBlock struct {
	content string
}

func NewTextBlock(content string) (*TextBlock, error) {
	block := &TextBlock{}
	block.content = content
	return block, nil
}

func (TextBlock) Type() string {
	return "text"
}
