package tests

import (
	"github.com/gouef/translator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslator(t *testing.T) {
	t.Run("Simple just PO", func(t *testing.T) {
		p, err := translator.NewPO("./locales")

		assert.NoError(t, err)

		translator.Init()
		translator.Register(p)

		if p != nil {
			err = translator.SetLanguage("cs_CZ")
			assert.NoError(t, err)

			assert.Equal(t, "cs_CZ", translator.GetLanguage())

			tr := translator.Translate("errors.hello")
			assert.Equal(t, "Error ahoj", tr)
		}
	})

	t.Run("Simple non exists lang PO", func(t *testing.T) {
		p, err := translator.NewPO("./locales")

		assert.NoError(t, err)

		translator.Init()
		translator.Register(p)

		err = translator.SetLanguage("cs_EN")
		assert.Error(t, err)
	})

	t.Run("Simple just Yaml", func(t *testing.T) {
		p, err := translator.NewYaml("./locales")

		assert.NoError(t, err)

		translator.Init()
		translator.Register(p)

		if p != nil {
			err = translator.SetLanguage("cs_CZ")
			assert.NoError(t, err)

			assert.Equal(t, "cs_CZ", translator.GetLanguage())

			tr := translator.Translate("messages.errors.hello")
			assert.Equal(t, "Error ahoj", tr)
		}
	})

	t.Run("Simple combine Yaml & PO", func(t *testing.T) {
		po, err := translator.NewPO("./locales")
		assert.NoError(t, err)
		ya, err := translator.NewYaml("./locales")
		assert.NoError(t, err)

		translator.Init()
		translator.Register(ya)
		translator.Register(po)

		err = translator.SetLanguage("cs_CZ")
		assert.NoError(t, err)

		assert.Equal(t, "cs_CZ", translator.GetLanguage())

		tr := translator.Translate("messages.errors.hello")
		assert.Equal(t, "Error ahoj", tr)
	})

	t.Run("Simple combine PO & Yaml", func(t *testing.T) {
		po, err := translator.NewPO("./locales")
		assert.NoError(t, err)
		ya, err := translator.NewYaml("./locales")
		assert.NoError(t, err)

		translator.Init()
		translator.Register(po)
		translator.Register(ya)

		err = translator.SetLanguage("cs_CZ")
		assert.NoError(t, err)

		assert.Equal(t, "cs_CZ", translator.GetLanguage())

		tr := translator.Translate("messages.errors.hello")
		assert.Equal(t, "messages.errors.hello", tr)
	})

	t.Run("Simple combine Yaml & PO skip to next", func(t *testing.T) {
		po, err := translator.NewPO("./locales")
		assert.NoError(t, err)
		ya, err := translator.NewYaml("./locales")
		assert.NoError(t, err)

		translator.Init()
		translator.Register(ya)
		translator.Register(po)

		err = translator.SetLanguage("cs_CZ")
		assert.NoError(t, err)

		assert.Equal(t, "cs_CZ", translator.GetLanguage())

		tr := translator.Translate("errors.hello")
		assert.Equal(t, "Error ahoj", tr)
	})

	t.Run("Simple combine Yaml return key", func(t *testing.T) {
		ya, err := translator.NewYaml("./locales")
		assert.NoError(t, err)

		translator.Init()
		translator.Register(ya)

		err = translator.SetLanguage("en_GB")
		assert.NoError(t, err)

		assert.Equal(t, "en_GB", translator.GetLanguage())

		tr := translator.Translate("errors.hello")
		assert.Equal(t, "errors.hello", tr)
	})
}
