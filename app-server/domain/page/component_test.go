package page_test

import (
	"testing"

	"github.com/Ryo-Sasaki-xxx/pape-coffee/domain/page"

	"github.com/stretchr/testify/assert"
)

func TestComponent(t *testing.T) {

	t.Run("success NewComponent", func(t *testing.T) {
		textBlock, _ := page.NewBlock("text")
		component, err := page.NewComponent(textBlock)
		assert.Equal(t, []page.Block{textBlock}, component.ExportBlock())
		assert.Nil(t, err)
	})
}
