package page_test

import (
	"testing"

	"github.com/Ryo-Sasaki-xxx/pape_coffee/domain/page"

	"github.com/stretchr/testify/assert"
)

func TestContent(t *testing.T) {

	t.Run("success NewContent", func(t *testing.T) {
		value := "hogehoge"
		attributes := []string{"plane_text"}
		content, err := page.NewContent(attributes, value)
		assert.Equal(t, value, content.ExportValue())
		assert.Nil(t, err)
	})
}
