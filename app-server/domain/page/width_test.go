package page_test

import (
	"testing"

	"github.com/Ryo-Sasaki-xxx/pape_coffee/domain/page"

	"github.com/stretchr/testify/assert"
)

func TestWidth(t *testing.T) {

	t.Run("success NewComponent", func(t *testing.T) {
		width, err := page.NewWidth(1)
		assert.Equal(t, 1, width.ExportValue())
		assert.Nil(t, err)
	})

	t.Run("fail NewComponent", func(t *testing.T) {
		width, err := page.NewWidth(11)
		assert.Error(t, err)
		assert.Nil(t, width)
	})

	t.Run("success Add", func(t *testing.T) {
		one, _ := page.NewWidth(1)
		two, err := one.Add(*one)
		assert.Equal(t, 2, two.ExportValue())
		assert.Nil(t, err)
	})

	t.Run("fail Add", func(t *testing.T) {
		one, _ := page.NewWidth(1)
		ten, _ := page.NewWidth(10)
		eleven, err := one.Add(*ten)
		assert.Error(t, err)
		assert.Nil(t, eleven)
	})

}
