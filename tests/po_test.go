package tests

import (
	"github.com/gouef/translator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPO(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		p, err := translator.NewPO("./locales")

		assert.NoError(t, err)
		if p != nil {
			err = p.SetLanguage("cs_CZ")
			assert.NoError(t, err)

			assert.Equal(t, "cs_CZ", p.GetLanguage())

			err, tr := p.Translate("messages.hello")
			assert.NoError(t, err)
			assert.Equal(t, "Ahoj", tr)
		}
	})

	t.Run("Set Language error", func(t *testing.T) {
		p, err := translator.NewPO("./locales")

		assert.NoError(t, err)
		if p != nil {
			err = p.SetLanguage("cs_EN")
			assert.Error(t, err)
		}
	})

	t.Run("Non Language error", func(t *testing.T) {
		p, err := translator.NewPO("./locales2")

		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("Domain translate", func(t *testing.T) {
		p, err := translator.NewPO("./locales")

		assert.NoError(t, err)
		if p != nil {
			err = p.SetLanguage("cs_CZ")
			assert.NoError(t, err)

			assert.Equal(t, "cs_CZ", p.GetLanguage())

			err, tr := p.Translate("errors.hello")
			assert.NoError(t, err)
			assert.Equal(t, "Error ahoj", tr)
		}
	})
}
