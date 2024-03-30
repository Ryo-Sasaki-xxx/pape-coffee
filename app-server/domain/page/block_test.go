package page_test

import (
	"testing"

	"github.com/Ryo-Sasaki-xxx/pape-coffee/domain/page"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {

	t.Run("success NewTextBlock", func(t *testing.T) {
		content := "hogehoge"
		block, err := page.NewTextBlock(content)
		assert.Equal(t, "text", block.Type())
		assert.Equal(t, content, block.ExportContent())
		assert.Nil(t, err)
	})

	t.Run("success NewBlock", func(t *testing.T) {
		blockType := "text"
		block, err := page.NewBlock(blockType)
		assert.Equal(t, blockType, block.Type())
		assert.Nil(t, err)
	})

	t.Run("success EditContent", func(t *testing.T) {
		blockType := "text"
		block, err := page.NewBlock(blockType)
		assert.Equal(t, blockType, block.Type())
		assert.Nil(t, err)
	})
}
