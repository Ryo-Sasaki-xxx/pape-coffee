package page_test

import (
	"testing"

	"github.com/Ryo-Sasaki-xxx/pape_coffee/domain/page"
	"github.com/stretchr/testify/assert"
)

func TestAttribute(t *testing.T) {

	t.Run("success NewAttribute", func(t *testing.T) {
		bold, err := page.NewAttribute("bold")
		assert.Equal(t, page.NewBold(), bold)
		assert.Nil(t, err)
	})
}
