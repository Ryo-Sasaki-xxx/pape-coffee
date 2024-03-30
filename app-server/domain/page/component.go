package page

type component struct {
	blocks []Block
}

func NewComponent(block Block) (*component, error) {
	component := &component{}
	component.blocks = []Block{block}
	return component, nil
}
