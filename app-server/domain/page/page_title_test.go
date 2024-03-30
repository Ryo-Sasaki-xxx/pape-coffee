package page_test

import (
	"testing"

	"github.com/Ryo-Sasaki-xxx/pape-coffee/domain/page"

	"github.com/stretchr/testify/assert"
)

func TestPageId(t *testing.T) {

	t.Run("success NewPageTitle()", func(t *testing.T) {
		title := "hogehogefuga"
		pageTitle, err := page.NewPageTitle(title)
		assert.Equal(t, title, pageTitle.ExportTitle())
		assert.Nil(t, err)
	})

	t.Run("maxLengthError NewPageTitle", func(t *testing.T) {
		title := "hogehogheohogehgoehohgoehohgoehohogeho"
		pageTitle, err := page.NewPageTitle(title)
		assert.Nil(t, pageTitle)
		assert.Error(t, err)
	})

	t.Run("true IsEmpty() ", func(t *testing.T) {
		title := ""
		pageTitle, err := page.NewPageTitle(title)
		assert.True(t, pageTitle.IsEmpty())
		assert.Nil(t, err)
	})

}
